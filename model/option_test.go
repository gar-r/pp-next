package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVoteOption_HasIcon(t *testing.T) {

	opt := &VoteOption{}

	t.Run("has no icon", func(t *testing.T) {
		opt.Icon = ""
		assert.False(t, opt.HasIcon())
	})

	t.Run("has icon", func(t *testing.T) {
		opt.Icon = "something"
		assert.True(t, opt.HasIcon())
	})

}

func TestVoteOption_Visible(t *testing.T) {

	opt := &VoteOption{}

	t.Run("visible", func(t *testing.T) {
		opt.Hidden = false
		assert.True(t, opt.Visible())
	})

	t.Run("invisible", func(t *testing.T) {
		opt.Hidden = true
		assert.False(t, opt.Visible())
	})

}

func TestVoteOption_IsChecked(t *testing.T) {

	opt := &VoteOption{}
	r := NewRoom("test")

	t.Run("checked", func(t *testing.T) {
		r.RegisterVote(NewVote("user", 2))
		opt.Value = 2
		assert.Equal(t, "checked", opt.IsChecked("user", r))
	})

	t.Run("unchecked", func(t *testing.T) {
		r.RegisterVote(NewVote("user", 3))
		opt.Value = 2
		assert.Equal(t, "", opt.IsChecked("user", r))
	})

	t.Run("no vote", func(t *testing.T) {
		opt.Value = 2
		assert.Equal(t, "", opt.IsChecked("user", r))
	})

}
