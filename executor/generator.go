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
	"github.com/gobench-io/gobench/executor"
	"github.com/gobench-io/gobench/logger"
)

func main() {
	logger := logger.NewStdLogger()

	e, err := executor.NewExecutor(&executor.Options{
		AgentSock:    "{{ .AgentSock }}",
		ExecutorSock: "{{ .ExecutorSock }}",
		AppID:        {{ .AppID }},
		Vus:          export(),
	}, logger)

	if err != nil {
		panic(err)
	}

	e.Serve()
}
`))

// Generate creates an executor go file that is used to compiled to a binary
func Generate(wr io.Writer, agentSock, executorSock string, appID int) (err error) {
	type Args struct {
		AgentSock    string
		ExecutorSock string
		AppID        int
	}

	err = tmpl.Execute(wr, Args{
		agentSock, executorSock, appID,
	})

	return
}
