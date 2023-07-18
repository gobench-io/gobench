package executor

import (
	"html/template"
	"io"
)

var tmpl = template.Must(template.
	New("executor").
	Parse(`
package main

import (
	"flag"
	"os"

	"github.com/gobench-io/gobench/v2/executor"
	"github.com/gobench-io/gobench/v2/logger"
)

func configureOptions(fs *flag.FlagSet, args []string) (
	agentSock, executorSock string, err error,
) {
	fs.StringVar(&agentSock, "agent-sock", "", "Socket of the agent")
	fs.StringVar(&executorSock, "executor-sock", "", "Socket for this executor")
	err = fs.Parse(args)
	return
}

func main() {
	exe := "executor"
	logger := logger.NewStdLogger()

	fs := flag.NewFlagSet(exe, flag.ExitOnError)

	agentSock, executorSock, err := configureOptions(fs, os.Args[1:])
	if err != nil {
		logger.Fatalw("Fail getting args", "err", err)

	}

	e, err := executor.NewExecutor(&executor.Options{
		AgentSock:    agentSock,
		ExecutorSock: executorSock,
		AppID:        {{ .AppID }},
		Vus:          export(),
	}, logger)

	if err != nil {
		panic(err)
	}

	logger.Infow("serving",
		"agent sock", agentSock,
		"executor sock", executorSock,
	)

	if err = e.Serve(); err != nil {
		logger.Fatalf("Fail serving", "err", err)
	}
}
`))

// Generate creates an executor go file that is used to compiled to a binary
func Generate(wr io.Writer, appID int) (err error) {
	type Args struct {
		AppID int
	}

	err = tmpl.Execute(wr, Args{
		appID,
	})

	return
}
