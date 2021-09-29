package model

import (
	"fmt"
	"sort"
	"time"

	"okki.hu/garric/ppnext/consts"
)

// Room represents a planning poker room
type Room struct {
	Name       string           `json:"name"`
	Revealed   bool             `json:"revealed"`
	RevealedBy string           `json:"revealedBy"`
	Votes      map[string]*Vote `json:"votes"`
	Ts         time.Time        `json:"ts"`
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

// Average returns a string representation of the average value of current votes.
// Special votes defined in the consts package are not counted.
func (r *Room) Average() string {
	sum := 0
	cnt := 0
	for _, v := range r.Votes {
		if v.IsCoffee() || v.IsLarge() || v.IsNothing() || v.IsQuestion() {
			continue
		}
		sum += v.Vote
		cnt += 1
	}
	avg := float64(sum) / float64(cnt)
	return fmt.Sprintf("%.2f", avg)
}

// Summary counts votes by occurence and returns a slice of SummaryItems
// representing the groups. The slice is sorted by category in ascending order.
func (r *Room) Summary() []*SummaryItem {
	result := make([]*SummaryItem, 0)
	m := r.summaryMap()
	cat := make([]int, 0, len(m))
	for k := range m {
		cat = append(cat, k)
	}
	sort.Ints(cat)
	for _, k := range cat {
		result = append(result, &SummaryItem{
			Category: k,
			Count:    m[k],
		})
	}
	return result
}

// SummaryItem represents a single row of the summary table
type SummaryItem struct {
	Category int
	Count    int
}

func (r *Room) summaryMap() map[int]int {
	m := make(map[int]int)
	for _, v := range r.Votes {
		vote := v.Vote
		i, ok := m[vote]
		if !ok {
			m[vote] = 1
		} else {
			m[vote] = i + 1
		}
	}
	return m
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

// IsNothing returns if the vote value equals the const Nothing
func (v *Vote) IsNothing() bool {
	return v.Vote == consts.Nothing
}

// IsCoffee returns if the vote value equals the const Coffee
func (v *Vote) IsCoffee() bool {
	return v.Vote == consts.Coffee
}

// IsLarge returns if the vote value equals the const Large
func (v *Vote) IsLarge() bool {
	return v.Vote == consts.Large
}

// IsQuestion returns if the vote value equals the const Question
func (v *Vote) IsQuestion() bool {
	return v.Vote == consts.Question
}
