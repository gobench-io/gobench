package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNodeWithConfig(t *testing.T) {
// }

func TestNew(t *testing.T) {
	n1, err := New()
	assert.Nil(t, err)
	n2, err := New()
	assert.Nil(t, err)

	assert.Equal(t, n1, n2)

	assert.Equal(t, n1.status, idle)
	assert.Nil(t, n1.cancel)
	assert.Equal(t, n1.units, make(map[string]unit))
}

func TestLoadPlugin(t *testing.T) {

}
