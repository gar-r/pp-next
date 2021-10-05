package store

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"okki.hu/garric/ppnext/model"
)

func Test_FsRepository_RoomPath(t *testing.T) {

	path := createTempPath(t)
	name := "demo"
	expected := path + "/demo"

	t.Run("no ending path separator", func(t *testing.T) {
		r := NewFs(path)
		p := r.getRoomPath(name)
		assert.Equal(t, expected, p)
	})

	t.Run("with ending path separator", func(t *testing.T) {
		r := NewFs(path + string(os.PathSeparator))
		p := r.getRoomPath(name)
		assert.Equal(t, expected, p)
	})

}

func Test_FsRepository_SaveLoad(t *testing.T) {

	name := "test"

	repo := NewFs(createTempPath(t))
	r1 := model.NewRoom(name)

	err := repo.Save(r1)
	assert.NoError(t, err)

	r2, err := repo.Load(name)
	assert.NoError(t, err)

	assert.Equal(t, r1.Name, r2.Name)

	err = repo.Delete(name)
	assert.NoError(t, err)

	path := repo.getRoomPath(name)
	assert.NoDirExists(t, path)
}

func Test_FsRepository_LoadNonExisting(t *testing.T) {
	repo := NewFs(createTempPath(t))
	r, err := repo.Load("test")
	defer repo.Delete("test")
	assert.NoError(t, err)
	assert.Equal(t, "test", r.Name)
}

func Test_FsRepository_Exists(t *testing.T) {

	repo := NewFs(createTempPath(t))
	user := "user"

	// empty repository, user should not exist
	ex, err := repo.Exists(user)
	assert.NoError(t, err)
	assert.Falsef(t, ex, "user '%s' should not exist")

	// add user and verify it exists
	r, err := repo.Load("room")
	assert.NoError(t, err)
	r.RegisterVote(&model.Vote{
		User: user,
		Vote: 3,
	})
	repo.Save(r)
	defer repo.Delete(r.Name)
	ex, err = repo.Exists(user)
	assert.NoError(t, err)
	assert.Truef(t, ex, "expected user '%s' to exist")
}

func Test_FsRepository_Cleanup(t *testing.T) {
	repo := NewFs(createTempPath(t))
	r1 := model.NewRoom("obsolete")
	r1.ResetTs = time.Now().Add(-1 * time.Hour)
	repo.Save(r1)
	r2 := model.NewRoom("fresh")
	r2.ResetTs = time.Now()
	repo.Save(r2)

	err := repo.Cleanup(55 * time.Minute)

	assert.NoError(t, err)

	x, _ := repo.Load("obsolete")
	assert.NotEqual(t, x.ResetTs.UnixMilli(), r1.ResetTs.UnixMilli())

	y, _ := repo.Load("fresh")
	assert.Equal(t, y.ResetTs.UnixMilli(), r2.ResetTs.UnixMilli())
}

func createTempPath(t *testing.T) string {
	t.Helper()
	path, err := os.MkdirTemp(os.TempDir(), "fstest")
	if err != nil {
		t.Error(err)
	}
	return path
}
