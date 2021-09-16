package store

import (
	"errors"
	"testing"

	"okki.hu/garric/ppnext/model"
)

func Test_Cache_Load(t *testing.T) {

	repo := &testRepository{}
	cache := NewCache(repo)

	t.Run("cache miss, followed by cache hit", func(t *testing.T) {
		// first call to Load results in a cache miss:
		room := repo.StubRoom("test")

		actual, _ := cache.Load("test")

		if actual != room {
			t.Errorf("expected %v, got %v", room, actual)
		}

		// second call to Load is a cache hit:
		repo.Room = nil

		actual, _ = cache.Load("test")

		if actual != room {
			t.Errorf("expected %v, got %v", room, actual)
		}
	})

	t.Run("error propagation", func(t *testing.T) {
		testErr := errors.New("test")
		repo.Err = testErr

		_, err := cache.Load("A")

		if err != testErr {
			t.Errorf("expected %v, got %v", testErr, err)
		}

	})

}

func Test_Cache_Save(t *testing.T) {

	repo := &testRepository{}
	cache := NewCache(repo)

	t.Run("save invalidates the cache", func(t *testing.T) {
		room := repo.StubRoom("A")

		// first let the room be cached
		cache.Load("A")

		// save should invalidate the cache
		cache.Save(room)

		room = repo.StubRoom("B")

		// reload, should pick up B
		actual, _ := cache.Load("A")

		if actual != room {
			t.Errorf("expected %v, got %v", room, actual)
		}
	})

	t.Run("error propagation", func(t *testing.T) {
		testErr := errors.New("test")
		repo.Err = testErr

		err := cache.Save(&model.Room{})

		if err != testErr {
			t.Errorf("expected %v, got %v", testErr, err)
		}
	})

}

func Test_Cache_Delete(t *testing.T) {

	repo := &testRepository{}
	cache := NewCache(repo)

	t.Run("delete invalidates cache", func(t *testing.T) {
		repo.StubRoom("A")

		// first let the room be cached
		cache.Load("A")

		// save should invalidate the cache
		cache.Delete("A")

		room := repo.StubRoom("B")

		// reload, should pick up B
		actual, _ := cache.Load("A")

		if actual != room {
			t.Errorf("expected %v, got %v", room, actual)
		}
	})

	t.Run("error propagation", func(t *testing.T) {
		testErr := errors.New("test")
		repo.Err = testErr

		err := cache.Delete("A")

		if err != testErr {
			t.Errorf("expected %v, got %v", testErr, err)
		}
	})

}

func Test_Cache_Invalidate(t *testing.T) {

	repo := &testRepository{}
	cache := NewCache(repo)

	repo.StubRoom("A")

	// load into cache
	cache.Load("A")

	room := repo.StubRoom("B")

	// invalidate cache
	cache.Invalidate("A")

	// reload, should pick up B
	actual, _ := cache.Load("A")

	if actual != room {
		t.Errorf("expected %v, got %v", room, actual)
	}

}

type testRepository struct {
	Room *model.Room
	Err  error
}

func (r *testRepository) StubRoom(name string) *model.Room {
	room := model.NewRoom(name)
	r.Room = room
	return room
}

func (r *testRepository) Load(name string) (*model.Room, error) {
	return r.Room, r.Err
}

func (r *testRepository) Save(room *model.Room) error {
	return r.Err
}

func (r *testRepository) Delete(name string) error {
	return r.Err
}
