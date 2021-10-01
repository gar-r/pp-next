package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encode_Decode(t *testing.T) {
	r := NewRoom("name")

	// encode
	reader, err := Encode(r)
	assert.NoError(t, err)

	// decode
	s, err := Decode(reader)
	assert.NoError(t, err)

	assert.Equal(t, r.Name, s.Name)
	assert.Equal(t, r.ResetTs.UnixMilli(), s.ResetTs.UnixMilli())
}
