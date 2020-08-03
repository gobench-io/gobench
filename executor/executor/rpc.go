package executor

import "context"

func (e *Executor) Start(args *bool, reply *bool) (err error) {
	ctx := context.TODO()

	err = e.driver.Run(ctx)

	return
}
