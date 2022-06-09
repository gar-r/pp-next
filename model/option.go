package model

const (
	Nothing  = 900 // did not vote (default)
	Coffee   = 901 // needs a break
	Large    = 902 // story too large
	Question = 903 // needs discussion
)

var VoteOptions = []*VoteOption{
	{Text: "½", Value: 0.5},
	{Text: "1", Value: 1},
	{Text: "2", Value: 2},
	{Text: "3", Value: 3},
	{Text: "5", Value: 5},
	{Text: "8", Value: 8},
	{Text: "13", Value: 13},
	{Text: "20", Value: 20},
	{Text: "40", Value: 40},
	{Icon: "all_inclusive", Value: Large},
	{Icon: "coffee", Value: Coffee},
	{Icon: "help", Value: Question},
	{Icon: "hourglass_full", Value: Nothing, Hidden: true},
}

var VoteLookup = map[float64]*VoteOption{
	0.5:      VoteOptions[0],
	1:        VoteOptions[1],
	2:        VoteOptions[2],
	3:        VoteOptions[3],
	5:        VoteOptions[4],
	8:        VoteOptions[5],
	13:       VoteOptions[6],
	20:       VoteOptions[7],
	40:       VoteOptions[8],
	Large:    VoteOptions[9],
	Coffee:   VoteOptions[10],
	Question: VoteOptions[11],
	Nothing:  VoteOptions[12],
}

type VoteOption struct {
	Text   string
	Icon   string
	Value  float64
	Hidden bool
}

func (v *VoteOption) HasIcon() bool {
	return v.Icon != ""
}

func (v *VoteOption) Visible() bool {
	return !v.Hidden
}

func (v *VoteOption) IsChecked(user string, room *Room) string {
	vote, ok := room.Votes[user]
	if !ok {
		return ""
	}
	if v.Value == vote.Vote {
		return "checked"
	}
	return ""
}
