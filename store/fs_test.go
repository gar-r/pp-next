package store

import (
	"os"
	"testing"

	"okki.hu/garric/ppnext/model"
)

func Test_FsRepository_RoomPath(t *testing.T) {

	name := "demo"
	expected := "test/demo"

	t.Run("no ending path separator", func(t *testing.T) {
		r := NewFs("test")
		p := r.getRoomPath(name)
		if p != expected {
			t.Errorf("expected %s, got %s", expected, p)
		}
	})

	t.Run("with ending path separator", func(t *testing.T) {
		r := NewFs("test" + string(os.PathSeparator))
		p := r.getRoomPath(name)
		if p != expected {
			t.Errorf("expected %s, got %s", expected, p)
		}
	})

}

func Test_FsRepository_SaveLoad(t *testing.T) {

	name := "test"

	repo := NewFs(createTempPath(t))
	r1 := model.NewRoom(name)

	err := repo.Save(r1)
	if err != nil {
		t.Error(err)
	}

	r2, err := repo.Load(name)
	if err != nil {
		t.Error(err)
	}

	if r1.Name != r2.Name {
		t.Errorf("expected %s, got %s", r1.Name, r2.Name)
	}

	err = repo.Delete(name)
	if err != nil {
		t.Error(err)
	}

	path := repo.getRoomPath(name)
	_, err = os.Stat(path)
	if err == nil {
		t.Errorf("path '%s' still exists after delete", path)
	}

}

func Test_FsRepository_LoadNonExisting(t *testing.T) {
	repo := NewFs(createTempPath(t))
	r, err := repo.Load("test")
	defer repo.Delete("test")
	if err != nil {
		t.Error(err)
	}
	if r.Name != "test" {
		t.Errorf("expected %s, got %s", "test", r.Name)
	}
}

func Test_FsRepository_Exists(t *testing.T) {

	repo := NewFs(createTempPath(t))
	user := "user"

	// empty repository, user should not exist
	ex, err := repo.Exists(user)
	if err != nil {
		t.Error(err)
	}
	if ex {
		t.Errorf("user '%s' should not exist", user)
	}

	// add user and verify it exists
	r, err := repo.Load("room")
	if err != nil {
		t.Error(err)
	}
	r.RegisterVote(&model.Vote{
		User: user,
		Vote: 3,
	})
	repo.Save(r)
	ex, err = repo.Exists(user)
	if err != nil {
		t.Error(err)
	}
	if !ex {
		t.Errorf("expected user '%s' to exist", user)
	}
}

func createTempPath(t *testing.T) string {
	t.Helper()
	path, err := os.MkdirTemp(os.TempDir(), "fstest")
	if err != nil {
		t.Error(err)
	}
	return path
}
