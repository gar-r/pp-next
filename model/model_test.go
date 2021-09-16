package model

import "testing"

func TestRoom_NewRoom(t *testing.T) {
	name := "test"
	r := NewRoom(name)
	if r.Name != name {
		t.Errorf("expected %s, got %s", name, r.Name)
	}
}

func TestRoom_RegisterVote(t *testing.T) {
	r := NewRoom("test")

	v := &Vote{
		User: "user",
		Vote: 5,
	}

	r.RegisterVote(v)

	if r.Votes["user"] != 5 {
		t.Errorf("expected %d, got %d", 5, r.Votes["user"])
	}
}
