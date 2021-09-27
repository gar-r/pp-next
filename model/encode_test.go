package model

import "testing"

func Test_Encode_Decode(t *testing.T) {
	r := NewRoom("name")

	// encode
	reader, err := Encode(r)
	if err != nil {
		t.Error(err)
	}

	// decode
	s, err := Decode(reader)
	if err != nil {
		t.Error(err)
	}

	if r.Name != s.Name {
		t.Errorf("expected %v, got %v", r, s)
	}

}
