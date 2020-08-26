package scenario

import (
	"context"
)

type VuFunc func(context.Context, int)

type Vu struct {
	Nu   int
	Rate float64
	Fu   VuFunc
}

type Vus []Vu
