package sample

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunction(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{32, "255.255.255.255"},
		{24, "255.255.255.0"},
		{16, "255.255.0.0"},
		{8, "255.0.0.0"},
		{0, ""},
		{-1, ""},
		{33, ""},
	}
	for _, tt := range tests {
		assert := assert.New(t)
		res, _ := CidrToMask(tt.in)
		assert.Equal(tt.out, res, "The two words should be the same.")
	}
}
