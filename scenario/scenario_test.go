package scenario

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPlugin(t *testing.T) {
	// so := "scenario/scripts/valid/valid.so"
	so := "scripts/valid/valid.so"

	vus, err := LoadPlugin(so)
	assert.Nil(t, err)
	log.Println(vus)
	log.Println(err)
	assert.Error(t, err)
	assert.Equal(t, Vus{
		{
			Nu:   20,
			Rate: 100,
		},
	}, vus)
}

func TestInvalidPlugin(t *testing.T) {
	so := "scenario/scripts/invalid/invalid.so"

	vus, err := LoadPlugin(so)
	log.Println(vus)
	assert.Error(t, err)
}
