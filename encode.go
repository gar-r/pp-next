package main

import (
	"bytes"
	"encoding/gob"
	"io"
)

func encode(r *Room) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)
	err := e.Encode(r)
	return buf, err
}

func decode(reader io.Reader) (*Room, error) {
	e := gob.NewDecoder(reader)
	var r Room
	err := e.Decode(&r)
	return &r, err
}
