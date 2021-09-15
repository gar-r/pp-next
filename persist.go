package main

import (
	"os"
	"sync"
)

type Repository interface {
	Load(name string) (*Room, error)
	Save(room *Room) error
	Delete(name string) error
}

type FsRepository struct {
	rootPath string
	mux      sync.Mutex
}

func NewFsRepository(rootPath string) *FsRepository {
	last := len(rootPath) - 1
	if rootPath[last] != os.PathSeparator {
		rootPath += string(os.PathSeparator)
	}
	return &FsRepository{
		rootPath: rootPath,
	}
}

func (r *FsRepository) Load(name string) (*Room, error) {
	path := r.getRoomPath(name)
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return NewRoom(name), nil
		}
		return nil, err
	}
	return decode(f)
}

func (r *FsRepository) Save(room *Room) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	path := r.getRoomPath(room.Name)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	buf, err := encode(room)
	if err != nil {
		return err
	}
	_, err = f.Write(buf.Bytes())
	return err
}

func (r *FsRepository) Delete(name string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	path := r.getRoomPath(name)
	return os.Remove(path)
}

func (r *FsRepository) getRoomPath(name string) string {
	return r.rootPath + name
}
