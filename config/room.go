package config

import (
	"okki.hu/garric/ppnext/consts"
	"okki.hu/garric/ppnext/viewmodel"
)

var VoteOptions = []viewmodel.VoteOption{
	{Text: "1", Value: 1},
	{Text: "2", Value: 2},
	{Text: "3", Value: 3},
	{Text: "5", Value: 5},
	{Text: "7", Value: 7},
	{Text: "13", Value: 13},
	{Icon: "upgrade", Value: consts.Large},
	{Icon: "coffee", Value: consts.Coffee},
	{Icon: "help", Value: consts.Question},
}
