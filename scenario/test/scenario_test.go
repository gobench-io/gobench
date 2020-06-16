package scenario_test

import (
	"testing"

	"github.com/gobench-io/gobench/scenario"
	"github.com/stretchr/testify/assert"
)

func TestValidPlugin(t *testing.T) {
	so := "valid.so"

	assert.FileExistsf(t, so, "file valid.so must be compiled first")

	vus, err := scenario.LoadPlugin(so)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(vus))
	assert.Equal(t, 20, vus[0].Nu)
	assert.EqualValues(t, 100, vus[0].Rate)
	assert.NotNil(t, vus[0].Fu)
}

func TestInvalidPlugin(t *testing.T) {
	so := "invalid.so"

	assert.FileExistsf(t, so, "file invalid.so must be compiled first")

	_, err := scenario.LoadPlugin(so)
	assert.NotNil(t, err)
}
