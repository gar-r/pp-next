package store

import (
	"os"
	"time"

	"github.com/juju/fslock"
	"okki.hu/garric/ppnext/model"
)

// Fs implements Repository with filesystem based storage.
type Fs struct {
	rootPath string
}

// NewFs returns a new Fs initialized to a given directory (root path).
func NewFs(rootPath string) *Fs {
	last := len(rootPath) - 1
	if rootPath[last] != os.PathSeparator {
		rootPath += string(os.PathSeparator)
	}
	err := os.MkdirAll(rootPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return &Fs{
		rootPath: rootPath,
	}
}

// Load reads and decodes a model.Room data from the filesystem.
// In case the file does not exist yet, a new empty room with the given name is returned.
func (r *Fs) Load(name string) (*model.Room, error) {
	path := r.getRoomPath(name)
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return model.NewRoom(name), nil
		}
		return nil, err
	}
	return model.Decode(f)
}

// Save persists a model.Room data to the filesystem. A file with the name of the model.Room
// is created. If it exists, it will be overwritten.
func (r *Fs) Save(room *model.Room) error {
	path := r.getRoomPath(room.Name)
	lock := fslock.New(path)
	err := lock.Lock()
	if err != nil {
		return err
	}
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	buf, err := model.Encode(room)
	if err != nil {
		return err
	}
	_, err = f.Write(buf.Bytes())
	return err
}

// Delete completely removes the file associated with a model.Room data.
func (r *Fs) Delete(name string) error {
	path := r.getRoomPath(name)
	return os.Remove(path)
}

// Exists scans every room file in the root path, and checks
// if the given user is present in any of them.
// This is a slow operation, since it has to scan all rooms.
func (r *Fs) Exists(user string) (bool, error) {
	files, err := r.getRooms()
	if err != nil {
		return false, err
	}
	for _, f := range files {
		room, err := r.Load(f)
		if err != nil {
			return false, err
		}
		for u := range room.Votes {
			if u == user {
				return true, nil
			}
		}
	}
	return false, nil
}

// Cleanup removes obsolete rooms from the filesystem.
func (r *Fs) Cleanup(maxAge time.Duration) error {
	ts := time.Now()
	files, err := r.getRooms()
	if err != nil {
		return err
	}
	for _, f := range files {
		room, err := r.Load(f)
		if err != nil {
			return err
		}
		if ts.Sub(room.ResetTs) > maxAge {
			r.Delete(room.Name)
		}
	}
	return nil
}

func (r *Fs) getRooms() (names []string, err error) {
	p, err := os.Open(r.rootPath)
	if err != nil {
		return
	}
	return p.Readdirnames(0)
}

func (r *Fs) getRoomPath(name string) string {
	return r.rootPath + name
}
