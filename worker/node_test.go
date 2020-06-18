package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func loadValidPlugin(n *Node) error {
	so := "./script/valid.so"
	return n.Load(so)
}

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

	assert.False(t, n1.Running())
}

func TestLoadPlugin(t *testing.T) {
	n, _ := New()
	so := "./script/valid.so"
	assert.Nil(t, n.Load(so))
	assert.NotNil(t, n.vus)
	assert.False(t, n.Running())
}

// func TestRunPlugin(t *testing.T) {
// 	n, _ := New()
// 	assert.Nil(t, loadValidPlugin(n))
// 	assert.Nil(t, n.Run())
// 	assert.False(t, n.Running())
// }
