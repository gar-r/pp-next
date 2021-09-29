package model

import (
	"testing"
	"time"

	"okki.hu/garric/ppnext/consts"
)

func TestRoom_NewRoom(t *testing.T) {
	name := "test"
	r := NewRoom(name)
	if r.Name != name {
		t.Errorf("expected %s, got %s", name, r.Name)
	}
}

func TestRoom_RegisterVote(t *testing.T) {
	r := NewRoom("test")

	v := NewVote("user", 5)
	r.RegisterVote(v)

	if r.Votes["user"] != v {
		t.Errorf("expected %v, got %v", v, r.Votes["user"])
	}
}

func TestRoom_Average(t *testing.T) {
	r := NewRoom("test")
	r.RegisterVote(NewVote("a", 5))
	r.RegisterVote(NewVote("b", 3))
	r.RegisterVote(NewVote("c", 7))
	r.RegisterVote(NewVote("d", consts.Coffee))
	r.RegisterVote(NewVote("e", consts.Nothing))
	r.RegisterVote(NewVote("f", consts.Question))
	r.RegisterVote(NewVote("g", consts.Large))

	avg := r.Average()
	expected := "2.14"

	if avg != expected {
		t.Errorf("expected %v, got %v", 0, avg)
	}

}

func TestRoom_Summary(t *testing.T) {
	r := NewRoom("test")
	r.RegisterVote(NewVote("a", 5))
	r.RegisterVote(NewVote("b", 3))
	r.RegisterVote(NewVote("c", 3))
	r.RegisterVote(NewVote("d", 3))
	r.RegisterVote(NewVote("e", 7))
	r.RegisterVote(NewVote("f", 7))
	r.RegisterVote(NewVote("g", consts.Question))
	r.RegisterVote(NewVote("h", consts.Question))

	sum := r.Summary()

	expected := []SummaryItem{
		{Category: 3, Count: 3},
		{Category: 5, Count: 1},
		{Category: 7, Count: 2},
		{Category: consts.Question, Count: 2},
	}

	if len(sum) != len(expected) {
		t.Errorf("size mismatch: %v, %v", sum, expected)
	}

	for i, s := range sum {
		e := expected[i]
		if s.Category != e.Category ||
			s.Count != e.Count {
			t.Errorf("expected %v, got %v", e, s)
		}
	}
}

func TestVote_Timestamp(t *testing.T) {
	v := NewVote("test", 1)
	now := time.Now()

	if v.Ts.After(now) {
		t.Errorf("new vote timestamp cannot be in the future: %v", v.Ts)
	}
}
