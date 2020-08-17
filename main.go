package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gobench-io/gobench/executor"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/master"
	"github.com/gobench-io/gobench/web"
)

// gobench -p 3000 -master -cluster localhost:3001
// gobench -worker -cluster localhost:5001 -route localhost:3001

var usageStr = `
Usage: gobench [options]

    --mode <mode>       Server mode. Must be one of the master, agent, executor mode.
                        Default is master
    --cluster-port <port>   Cluster port to solicit and connect (default: 6890)
                            Master and agent are required to have this option
    -h, --help          Show this message
    -v, --version       Show version

Master Options:
    -a, --addr <host>   Bind to host address (default: 0.0.0.0)
    -p, --port <port>   Use port for web client (default: 8080).

Agent Options:
    --route <host:port> The master address to solicit routes.
                        Every worker must have this option sothat worker can connect to a master

Executor Options:
    --agent-sock <file>     Unix socket address the local agent is listening to
    --executor-soc <file>   Unix socket address this executor will listen to
    --driver-path <file>    Location of the driver that this executor will execute
    --app-id <id>           ID of the application that this executor will run

`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

// printAndDie print message to Stderr and exit error
func printAndDie(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// printVersionAndExit will print our version and exit.
func printVersionAndExit() {
	fmt.Printf("gobench: v%s\n", VERSION)
	os.Exit(0)
}

func main() {
	exe := "gobench"

	// create a flagset and set the usage
	fs := flag.NewFlagSet(exe, flag.ExitOnError)
	fs.Usage = usage

	opts, err := ConfigureOptions(fs, os.Args[0:],
		printVersionAndExit,
		fs.Usage)
	if err != nil {
		printAndDie(fmt.Sprintf("%s: %s", exe, err))
	}

	logger := logger.NewStdLogger()

	if opts.Mode == Master {
		m, err := master.NewMaster(&master.Options{
			Port:    opts.Port,
			Program: opts.Program,
			DbPath:  opts.DbPath,
		}, logger)
		if err != nil {
			printAndDie(fmt.Sprintf("%s: %s", exe, err))
		}

		err = m.Start()
		if err != nil {
			printAndDie(fmt.Sprintf("%s: %s", exe, err))
		}
		web.Serve(m, logger)

		return
	}

	if opts.Mode == Executor {
		e, err := executor.NewExecutor(&executor.Options{
			AgentSock:    opts.AgentSock,
			ExecutorSock: opts.ExecutorSock,
			DriverPath:   opts.DriverPath,
			AppID:        opts.AppID,
		}, logger)
		if err != nil {
			printAndDie(fmt.Sprintf("%s: %s", exe, err))
		}
		e.Serve()

		return
	}
}
