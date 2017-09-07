package clear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	cases := []struct {
		os              string
		expectedCommand string
	}{
		{
			os:              "linux",
			expectedCommand: "clear",
		},
		{
			os:              "darwin",
			expectedCommand: "clear",
		},
		{
			os:              "android",
			expectedCommand: "clear",
		},
		{
			os:              "solaris",
			expectedCommand: "clear",
		},
		{
			os:              "openbsd",
			expectedCommand: "clear",
		},
		{
			os:              "freebsd",
			expectedCommand: "clear",
		},
		{
			os:              "windows",
			expectedCommand: "cls",
		},
	}
	for _, c := range cases {
		setGOOS(c.os)
		err := Clear()
		assert.NoError(t, err, c.os)
		assert.Contains(t, lastCommand.Path, c.expectedCommand, c.os)
	}
}

func TestFail(t *testing.T) {
	setGOOS("invalid os")
	err := Clear()
	assert.Error(t, err)
}
