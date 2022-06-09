package model

import (
	"sync"
	"time"
)

// Room represents a planning poker room
type Room struct {
	Name       string           `bson:"name"`
	Votes      map[string]*Vote `bson:"votes"`
	Revealed   bool             `bson:"revealed"`
	RevealedBy string           `bson:"revealedBy"`
	ResetBy    string           `bson:"resetBy"`
	ResetTs    time.Time        `bson:"resetTs"`
	mux        sync.Mutex
}

// NewRoom creates a new Room with a pre-defined name
func NewRoom(name string) *Room {
	return &Room{
		Name:    name,
		Votes:   make(map[string]*Vote),
		ResetTs: time.Now(),
	}
}

// RegisterVote makes the Room register a user Vote.
func (r *Room) RegisterVote(v *Vote) {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.Votes[v.User] = v
}

// Reset the room, clearing all the votes, but preserving
// the name of joined users. The user who requested the reset
// will be stored in the ResetBy field.
func (r *Room) Reset(user string) {
	r.Revealed = false
	r.RevealedBy = ""
	r.ResetBy = user
	r.ResetTs = time.Now()
	for name := range r.Votes {
		r.RegisterVote(NewVote(name, Nothing))
	}
}
