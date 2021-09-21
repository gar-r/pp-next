package config

import "okki.hu/garric/ppnext/model"

const Support = "email@example.com"
const Addr = ":38080"
const RoomsPath = "./rooms"

var VoteOptions = []model.VoteOption{
	{Content: "1", Value: 1},
	{Content: "2", Value: 2},
	{Content: "3", Value: 3},
	{Content: "5", Value: 5},
	{Content: "7", Value: 7},
	{Content: "13", Value: 13},
	{Content: "Large", Value: model.Large},
	{Content: "Coffee", Value: model.Coffee},
	{Content: "Question", Value: model.Question},
}
