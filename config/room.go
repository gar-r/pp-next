package config

import (
	"okki.hu/garric/ppnext/consts"
	"okki.hu/garric/ppnext/viewmodel"
)

var VoteOptions = []*viewmodel.VoteOption{
	{Text: "1", Value: 1},
	{Text: "2", Value: 2},
	{Text: "3", Value: 3},
	{Text: "5", Value: 5},
	{Text: "7", Value: 7},
	{Text: "13", Value: 13},
	{Icon: "upgrade", Value: consts.Large},
	{Icon: "coffee", Value: consts.Coffee},
	{Icon: "help", Value: consts.Question},
	{Icon: "hourglass_full", Value: consts.Nothing, Hidden: true},
}

var VoteLookup = map[int]*viewmodel.VoteOption{
	1:               VoteOptions[0],
	2:               VoteOptions[1],
	3:               VoteOptions[2],
	5:               VoteOptions[3],
	7:               VoteOptions[4],
	13:              VoteOptions[5],
	consts.Large:    VoteOptions[6],
	consts.Coffee:   VoteOptions[7],
	consts.Question: VoteOptions[8],
	consts.Nothing:  VoteOptions[9],
}
