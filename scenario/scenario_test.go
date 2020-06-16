package scenario

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPlugin(t *testing.T) {
	// so := "scripts/valid/valid.so"
	so := "valid.so"

	vus, err := LoadPlugin(so)
	assert.Nil(t, err)
	assert.Equal(t, Vus{
		{
			Nu:   20,
			Rate: 100,
		},
	}, vus)
}

func TestInvalidPlugin(t *testing.T) {
	so := "scripts/invalid/invalid.so"

	vus, err := LoadPlugin(so)
	log.Println(err)
	log.Println(vus)
	assert.NotNil(t, err)
}
