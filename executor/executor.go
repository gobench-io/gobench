// executor is started from agent as a separated process
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gobench-io/gobench/executor/option"
	"github.com/gobench-io/gobench/logger"
)

type Options struct {
	Logger       logger.Logger
	AgentSock    string
	ExecutorSock string
	DriverPath   string // the plugin user wrote
}

func usage() {
	const usageStr = `
Usage: executor [options]
	--executor-sock		The socket for this executor
	...
	-h, --help			Show this message
	-v, --version		Show version
`
	fmt.Println(usageStr)
	os.Exit(0)
}

func printVersionAndExit() {
	os.Exit(0)
}

func printAndDie(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// run the program with args
// --agent-sock
// --executor-sock
// --driver-path
func main() {
	exe := "executor"

	fs := flag.NewFlagSet(exe, flag.ExitOnError)
	fs.Usage = usage

	opts, err := option.ConfigureOptions(fs, os.Args[1:], printVersionAndExit, fs.Usage)
	if err != nil {
		printAndDie(fmt.Sprintf("%s: %s", exe, err))
	}
	log.Println(opts)
}
