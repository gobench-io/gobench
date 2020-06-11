package scenario

import (
	"sync"
)

type VuFunc func(int, *sync.WaitGroup)

type Vu struct {
	Number   int
	Function VuFunc
}

type Vus []Vu
