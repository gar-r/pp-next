package model

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoom_NewRoom(t *testing.T) {

	t.Run("room name", func(t *testing.T) {
		name := "name"
		r := NewRoom(name)
		assert.Equal(t, name, r.Name)
	})

	t.Run("new room timestamp", func(t *testing.T) {
		t1 := time.Now()
		r := NewRoom("test")
		t2 := time.Now()
		assert.False(t, r.ResetTs.Before(t1))
		assert.False(t, r.ResetTs.After(t2))
	})

	t.Run("votes map is initialized", func(t *testing.T) {
		r := NewRoom("test")
		assert.NotNil(t, r.Votes)
	})
}

func TestRoom_RegisterVote(t *testing.T) {

	t.Run("register single vote", func(t *testing.T) {
		r := NewRoom("test")
		v := NewVote("user", 5)
		r.RegisterVote(v)
		assert.Equal(t, v, r.Votes[v.User])
	})

	t.Run("register multiple votes for the same user", func(t *testing.T) {
		r := NewRoom("test")
		v1 := NewVote("user", 1)
		v2 := NewVote("user", 2)
		r.RegisterVote(v1)
		r.RegisterVote(v2)
		assert.Equal(t, v2, r.Votes["user"])
	})

}

func TestRoom_Reset(t *testing.T) {

	t.Run("users are retained after reset", func(t *testing.T) {
		r := NewRoom("test")
		r.RegisterVote(NewVote("a", 1))
		r.RegisterVote(NewVote("b", 1))
		r.RegisterVote(NewVote("c", 1))

		r.Reset("user")

		assert.Contains(t, r.Votes, "a")
		assert.Contains(t, r.Votes, "b")
		assert.Contains(t, r.Votes, "c")
	})

	t.Run("votes are reset", func(t *testing.T) {
		r := NewRoom("test")
		r.RegisterVote(NewVote("a", 1))
		r.RegisterVote(NewVote("b", 2))

		r.Reset("user")

		assert.Equal(t, float64(Nothing), r.Votes["a"].Vote)
		assert.Equal(t, float64(Nothing), r.Votes["b"].Vote)
	})

	t.Run("user requesting reset is saved", func(t *testing.T) {
		r := NewRoom("test")
		u := "user"
		r.Reset(u)
		assert.Equal(t, u, r.ResetBy)
	})

	t.Run("reset timestamp", func(t *testing.T) {
		r := NewRoom("test")
		t1 := time.Now()
		r.Reset("user")
		t2 := time.Now()
		assert.False(t, r.ResetTs.Before(t1))
		assert.False(t, r.ResetTs.After(t2))
	})

	t.Run("revealed flag is reset", func(t *testing.T) {
		r := NewRoom("test")
		r.Revealed = true
		r.Reset("user")
		assert.False(t, r.Revealed)
	})

	t.Run("revealedBy is reset", func(t *testing.T) {
		r := NewRoom("test")
		r.RevealedBy = "a"
		r.Reset("user")
		assert.Equal(t, "", r.RevealedBy)
	})

}

func TestRoom_Average(t *testing.T) {

	t.Run("average returned as string", func(t *testing.T) {
		r := NewRoom("test")
		r.RegisterVote(NewVote("a", 5))
		r.RegisterVote(NewVote("b", 2))
		r.RegisterVote(NewVote("c", 7))
		assert.Equal(t, "4.67", r.Average())
	})

	t.Run("special values are skipped", func(t *testing.T) {
		specials := []float64{
			Nothing,
			Coffee,
			Large,
			Question,
		}
		r := NewRoom("test")
		for i, s := range specials {
			r.RegisterVote(NewVote(strconv.Itoa(i), s))
		}

		assert.Equal(t, "0.00", r.Average())
	})
}

func TestRoom_Summary(t *testing.T) {
	r := NewRoom("test")
	r.RegisterVote(NewVote("a", 5))
	r.RegisterVote(NewVote("b", 3))
	r.RegisterVote(NewVote("c", 3))
	r.RegisterVote(NewVote("d", 3))
	r.RegisterVote(NewVote("e", 7))
	r.RegisterVote(NewVote("f", 7))
	r.RegisterVote(NewVote("g", Question))
	r.RegisterVote(NewVote("h", Question))

	sum := r.Summary()

	expected := []SummaryItem{
		{Category: 3, Count: 3},
		{Category: 5, Count: 1},
		{Category: 7, Count: 2},
		{Category: Question, Count: 2},
	}

	assert.Equal(t, len(expected), len(sum))
	for i, e := range expected {
		s := sum[i]
		assert.Equal(t, e.Category, s.Category)
		assert.Equal(t, e.Count, s.Count)
	}
}

func TestVote_NewVote(t *testing.T) {

	t.Run("username and vote", func(t *testing.T) {
		v := NewVote("user", 10)
		assert.Equal(t, "user", v.User)
		assert.Equal(t, float64(10), v.Vote)
	})

	t.Run("vote timestamp", func(t *testing.T) {
		t1 := time.Now()
		v := NewVote("x", 1)
		t2 := time.Now()
		assert.False(t, v.Ts.Before(t1))
		assert.False(t, v.Ts.After(t2))
	})

	t.Run("special values", func(t *testing.T) {

		t.Run("nothing", func(t *testing.T) {
			v := NewVote("", Nothing)
			assert.True(t, v.IsNothing())
		})

		t.Run("coffee", func(t *testing.T) {
			v := NewVote("", Coffee)
			assert.True(t, v.IsCoffee())
		})

		t.Run("large", func(t *testing.T) {
			v := NewVote("", Large)
			assert.True(t, v.IsLarge())
		})

		t.Run("question", func(t *testing.T) {
			v := NewVote("", Question)
			assert.True(t, v.IsQuestion())
		})

	})

}
