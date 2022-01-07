package viewmodel

import "okki.hu/garric/ppnext/model"

type VoteOption struct {
	Text   string
	Icon   string
	Value  int
	Hidden bool
}

func (v *VoteOption) HasIcon() bool {
	return v.Icon != ""
}

func (v *VoteOption) Visible() bool {
	return !v.Hidden
}

func (v *VoteOption) IsChecked(user string, room *model.Room) string {
	vote, ok := room.Votes[user]
	if !ok {
		return ""
	}
	if v.Value == vote.Vote {
		return "checked"
	}
	return ""
}
