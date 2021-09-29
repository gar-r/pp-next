package viewmodel

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
