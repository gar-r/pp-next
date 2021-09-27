package model

import "time"

// Room represents a planning poker room
type Room struct {
	Name     string
	Revealed bool
	Votes    map[string]*Vote
}

// NewRoom creates a new Room with a pre-defined name
func NewRoom(name string) *Room {
	return &Room{
		Name:  name,
		Votes: make(map[string]*Vote),
	}
}

// RegisterVote makes the Room register a user Vote
func (r *Room) RegisterVote(v *Vote) {
	r.Votes[v.User] = v
}

// Vote represents a single vote coming from a single user
type Vote struct {
	User string    `json:"user"`
	Vote int       `json:"vote"`
	Ts   time.Time `json:"ts"`
}

// NewVote creates a new Vote with the given user and vote.
// The vote is initialized to the current timestamp.
func NewVote(user string, vote int) *Vote {
	return &Vote{
		User: user,
		Vote: vote,
		Ts:   time.Now(),
	}
}

type LoginForm struct {
	Room string `form:"room"`
	Name string `form:"name"`
}

type LoginQueryParams struct {
	LoginForm
	Valid string `form:"valid"`
}

type VoteOption struct {
	Text  string
	Icon  string
	Value int
}

func (v *VoteOption) HasIcon() bool {
	return v.Icon != ""
}
