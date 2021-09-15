package main

import "testing"

func Test_Encode_Decode(t *testing.T) {
	r := NewRoom("test")

	// encode
	reader, err := encode(r)
	if err != nil {
		t.Error(err)
	}

	// decode
	s, err := decode(reader)
	if err != nil {
		t.Error(err)
	}

	if r.Name != s.Name {
		t.Errorf("expected %s, got %s", r.Name, s.Name)
	}

}
