package scenario_test

import (
	"testing"

	"github.com/gobench-io/gobench/scenario"
	"github.com/stretchr/testify/assert"
)

func TestValidPlugin(t *testing.T) {
	so := "valid.so"

	vus, err := scenario.LoadPlugin(so)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(vus))
	assert.Equal(t, 20, vus[0].Nu)
	assert.EqualValues(t, 100, vus[0].Rate)
	assert.NotNil(t, vus[0].Fu)
}

func TestInvalidPlugin(t *testing.T) {
	so := "invalid.so"
	_, err := scenario.LoadPlugin(so)
	assert.NotNil(t, err)
}
