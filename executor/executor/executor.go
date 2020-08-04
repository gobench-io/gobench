package executor

import (
	"github.com/gobench-io/gobench/executor/driver"
	"github.com/gobench-io/gobench/executor/option"
	"github.com/gobench-io/gobench/logger"
)

type Executor struct {
	logger logger.Logger
	appID  int
	driver *driver.Driver
}

func NewExecutor(opts *option.Options, logger logger.Logger) (e *Executor, err error) {
	e = &Executor{
		logger: logger,
		appID:  opts.AppID,
	}

	e.driver, err = driver.NewDriver(e, logger, opts.DriverPath, opts.AppID)
	if err != nil {
		return
	}

	return
}
