package model

import (
	"testing"
	"time"
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

func TestVote_Timestamp(t *testing.T) {
	v := NewVote("test", 1)
	now := time.Now()

	if v.Ts.After(now) {
		t.Errorf("new vote timestamp cannot be in the future: %v", v.Ts)
	}
}
