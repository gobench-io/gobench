package scenario

import (
	"context"
	"fmt"
	"plugin"
)

type VuFunc func(context.Context, int)

type Vu struct {
	Nu   int
	Rate float64
	Fu   VuFunc
}

type Vus []Vu

func LoadPlugin(filename string) (Vus, error) {
	p, err := plugin.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed opening module: %v", err)
	}

	xexpf, err := p.Lookup("Export")
	if err != nil {
		return nil, fmt.Errorf("cannot find Export: %v", err)
	}

	expf := xexpf.(func() Vus)

	return expf(), nil
}
