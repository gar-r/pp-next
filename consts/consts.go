package consts

import "time"

const Support = "email@example.com"
const Addr = ":38080"
const Domain = "localhost"
const RoomsPath = "./rooms"

const CleanupFrequency = 10 * time.Minute
const MaximumRoomAge = 12 * time.Hour

const (
	Nothing  = 100 // did not vote (default)
	Coffee   = 101 // needs a break
	Large    = 102 // story too large
	Question = 103 // needs discussion
)
