package config

import "okki.hu/garric/ppnext/model"

const Support = "email@example.com"
const Addr = ":38080"
const RoomsPath = "./rooms"

const (
	Nothing  = 100 // did not vote (default)
	Coffee   = 101 // needs a break
	Large    = 102 // story too large
	Question = 103 // needs discussion
)

var VoteOptions = []model.VoteOption{
	{Text: "1", Value: 1},
	{Text: "2", Value: 2},
	{Text: "3", Value: 3},
	{Text: "5", Value: 5},
	{Text: "7", Value: 7},
	{Text: "13", Value: 13},
	{Icon: "upgrade", Value: Large},
	{Icon: "coffee", Value: Coffee},
	{Icon: "help", Value: Question},
}
