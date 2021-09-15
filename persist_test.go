package main

import (
	"os"
	"testing"
)

func Test_FsRepository_RoomPath(t *testing.T) {

	name := "demo"
	expected := "test/demo"

	t.Run("no ending path separator", func(t *testing.T) {
		r := NewFsRepository("test")
		p := r.getRoomPath(name)
		if p != expected {
			t.Errorf("expected %s, got %s", expected, p)
		}
	})

	t.Run("with ending path separator", func(t *testing.T) {
		r := NewFsRepository("test" + string(os.PathSeparator))
		p := r.getRoomPath(name)
		if p != expected {
			t.Errorf("expected %s, got %s", expected, p)
		}
	})

}

func Test_FsRepository_SaveLoad(t *testing.T) {

	name := "test"
	repo := NewFsRepository(os.TempDir())
	r1 := NewRoom(name)

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
