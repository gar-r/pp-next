package model

import (
	"time"

	"okki.hu/garric/ppnext/consts"
)

// Room represents a planning poker room
type Room struct {
	Name     string           `json:"name"`
	Revealed bool             `json:"revealed"`
	Votes    map[string]*Vote `json:"votes"`
	Ts       time.Time        `json:"ts"`
}

// NewRoom creates a new Room with a pre-defined name
func NewRoom(name string) *Room {
	return &Room{
		Name:  name,
		Votes: make(map[string]*Vote),
		Ts:    time.Now(),
	}
}

// RegisterVote makes the Room register a user Vote.
func (r *Room) RegisterVote(v *Vote) {
	r.Votes[v.User] = v
}

// Reset the room, clearing all the votes, but preserving
// the name of joined users.
func (r *Room) Reset() {
	for name := range r.Votes {
		r.RegisterVote(NewVote(name, consts.Nothing))
	}
}

// Vote represents a single vote coming from a single user
type Vote struct {
	User string    `json:"user"`
	Vote int       `json:"vote"`
	Ts   time.Time `json:"ts"`
}

// NewVote creates a new Vote with the given user and vote.
// The vote is initialized with the current timestamp.
func NewVote(user string, vote int) *Vote {
	return &Vote{
		User: user,
		Vote: vote,
		Ts:   time.Now(),
	}
}
