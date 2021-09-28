package viewmodel

type VoteOption struct {
	Text  string
	Icon  string
	Value int
}

func (v *VoteOption) HasIcon() bool {
	return v.Icon != ""
}
