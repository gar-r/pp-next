package model

import "time"

const (
	Nothing  = iota // did not vote (default)
	Coffee   = iota // needs a break
	Large    = iota // story too large
	Question = iota // needs discussion
)

// Room represents a planning poker room
type Room struct {
	Name  string
	Votes map[string]int
}

// NewRoom creates a new Room with a pre-defined name
func NewRoom(name string) *Room {
	return &Room{
		Name:  name,
		Votes: make(map[string]int),
	}
}

// RegisterVote makes the Room register a user Vote
func (r *Room) RegisterVote(v *Vote) {
	r.Votes[v.User] = v.Vote
}

// Vote represents a single vote coming from a single user
type Vote struct {
	User string    `json:"user"`
	Vote int       `json:"vote"`
	Ts   time.Time `json:"ts"`
}

type LoginForm struct {
	Room string `form:"room"`
	Name string `form:"name"`
}

type LoginQueryParams struct {
	LoginForm
	Valid string `form:"valid"`
}
