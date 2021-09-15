package main

const (
	Coffee   = iota
	Large    = iota
	Question = iota
)

type Vote struct {
	User string `json:"user"`
	Vote int    `json:"vote"`
}

type Room struct {
	Name  string
	Votes map[string]int
}

func NewRoom(name string) *Room {
	return &Room{
		Name:  name,
		Votes: make(map[string]int),
	}
}

func (r *Room) RegisterVote(v *Vote) {
	r.Votes[v.User] = v.Vote
}
