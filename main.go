package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gobench-io/gobench/v2/logger"
	"github.com/gobench-io/gobench/v2/master"
	"github.com/gobench-io/gobench/v2/web"
)

// gobench -p 3000 -master -cluster localhost:3001
// gobench -worker -cluster localhost:5001 -route localhost:3001

var usageStr = `
Usage: gobench [options]

    --mode <mode>       Server mode. Must be one of the master, agent mode.
                        Default is master
    --cluster-port <port>   Cluster port to solicit and connect (default: 6890)
                            Master and agent are required to have this option
    -h, --help          Show this message
    -v, --version       Show version

Master Options:
    -a, --addr <host>   Bind to host address (default: 0.0.0.0)
    -p, --port <port>   Use port for web client (default: 8080).
    --dir <dir path>    Working directory (default: ${HOME}). The result database and logs will be stored on this folder.
    -db <file>          Location for the server database (default: ${HOME}/gobench.sqlite3)
    --admin-password    Password required to login web dashboard

Agent Options:
    --route <host:port> The master address to solicit routes.
                        Every worker must have this option sothat worker can connect to a master
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
			HomeDir: opts.Dir,
		}, logger)
		if err != nil {
			printAndDie(fmt.Sprintf("%s: %s", exe, err))
		}

		if err = m.Start(); err != nil {
			printAndDie(fmt.Sprintf("%s: %s", exe, err))
		}

		// cleanup the previous running session
		if err = m.CleanupRunningApps(); err != nil {
			printAndDie(fmt.Sprintf("%s: %s", exe, err))
		}

		web.Serve(m, opts.AdminPassword, logger)

		return
	}
}
