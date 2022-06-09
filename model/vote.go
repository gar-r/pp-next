package model

import (
	"time"
)

// Vote represents a single vote coming from a single user
type Vote struct {
	User string    `json:"user" bson:"user"`
	Vote float64   `json:"vote" bson:"vote"`
	Ts   time.Time `json:"ts"   bson:"ts"`
}

// NewVote creates a new Vote with the given user and vote.
// The vote is initialized with the current timestamp.
func NewVote(user string, vote float64) *Vote {
	return &Vote{
		User: user,
		Vote: vote,
		Ts:   time.Now(),
	}
}

// IsNothing returns if the vote value equals the const Nothing
func (v *Vote) IsNothing() bool {
	return v.Vote == Nothing
}

// IsCoffee returns if the vote value equals the const Coffee
func (v *Vote) IsCoffee() bool {
	return v.Vote == Coffee
}

// IsLarge returns if the vote value equals the const Large
func (v *Vote) IsLarge() bool {
	return v.Vote == Large
}

// IsQuestion returns if the vote value equals the const Question
func (v *Vote) IsQuestion() bool {
	return v.Vote == Question
}
