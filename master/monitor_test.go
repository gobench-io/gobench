package master

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMyUptime(t *testing.T) {
	// Make sure we print this stuff right.

	s := 22 * time.Second
	m := 4 * time.Minute
	h := 4 * time.Hour
	d := 32 * 24 * time.Hour
	y := 22 * 365 * 24 * time.Hour

	var flagtests = []struct {
		in  time.Duration
		out string
	}{
		{s, "22s"},
		{m + s, "4m22s"},
		{h + m + s, "4h4m22s"},
		{d + h + m + s, "32d4h4m22s"},
		{y + d + h + m + s, "22y32d4h4m22s"},
	}

	for _, tt := range flagtests {
		assert.Equal(t, tt.out, myUptime(tt.in))
	}
}
