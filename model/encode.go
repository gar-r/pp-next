package model

import (
	"bytes"
	"encoding/gob"
	"io"
)

// Encode tries to encode the supplied Room into binary.
func Encode(r *Room) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)
	err := e.Encode(r)
	return buf, err
}

// Decode uses the supplied reader, and tries to decode a Room.
func Decode(reader io.Reader) (*Room, error) {
	e := gob.NewDecoder(reader)
	var r Room
	err := e.Decode(&r)
	return &r, err
}
