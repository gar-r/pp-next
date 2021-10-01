package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"okki.hu/garric/ppnext/model"
)

func Test_Cache_Load(t *testing.T) {

	repo := &testRepository{}

	t.Run("cache miss, followed by cache hit", func(t *testing.T) {
		cache := NewCache(repo)

		// first call to Load results in a cache miss
		room := repo.StubRoom("test")
		actual, _ := cache.Load("test")

		assert.Equal(t, room, actual)

		// second call to Load should be a cache hit
		repo.Room = nil // change room in repository to something else
		actual, _ = cache.Load("test")

		assert.Equal(t, room, actual)
	})

	t.Run("error propagation", func(t *testing.T) {
		cache := NewCache(repo)
		repo.Err = errors.New("test")

		_, err := cache.Load("A")
		assert.Error(t, err)
	})

}

func Test_Cache_Save(t *testing.T) {

	repo := &testRepository{}

	t.Run("save invalidates the cache", func(t *testing.T) {
		cache := NewCache(repo)

		room := repo.StubRoom("A")

		// first let the room be cached
		cache.Load("A")

		// save should invalidate the cache
		cache.Save(room)

		room = repo.StubRoom("B")

		// reload, should pick up B
		actual, _ := cache.Load("A")

		assert.Equal(t, room, actual)
	})

	t.Run("error propagation", func(t *testing.T) {
		cache := NewCache(repo)
		repo.Err = errors.New("test")

		err := cache.Save(&model.Room{})
		assert.Error(t, err)
	})

}

func Test_Cache_Delete(t *testing.T) {

	repo := &testRepository{}

	t.Run("delete invalidates cache", func(t *testing.T) {
		cache := NewCache(repo)

		repo.StubRoom("A")

		// first let the room be cached
		cache.Load("A")

		// save should invalidate the cache
		cache.Delete("A")

		room := repo.StubRoom("B")

		// reload, should pick up B
		actual, _ := cache.Load("A")

		assert.Equal(t, room, actual)
	})

	t.Run("error propagation", func(t *testing.T) {
		cache := NewCache(repo)
		repo.Err = errors.New("test")

		err := cache.Delete("A")
		assert.Error(t, err)
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
	assert.Equal(t, room, actual)
}

func Test_Cache_Exists(t *testing.T) {

	repo := &testRepository{}

	t.Run("user exists in cache", func(t *testing.T) {
		cache := NewCache(repo)
		room := repo.StubRoom("test")
		room.RegisterVote(&model.Vote{
			User: "user",
			Vote: 5,
		})

		// load room into cache
		cache.Load("test")
		exists, _ := cache.Exists("user")
		assert.True(t, exists)
	})

	t.Run("user exists in repo only", func(t *testing.T) {
		cache := NewCache(repo)
		room := repo.StubRoom("test")
		room.RegisterVote(&model.Vote{
			User: "user",
			Vote: 5,
		})

		exists, _ := cache.Exists("user")
		assert.True(t, exists)
	})

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

func (r *testRepository) Exists(user string) (bool, error) {
	return true, nil
}
