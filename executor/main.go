// executor is started from agent as a separated process
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
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
	--agent-sock		The socket of the agent who call this process
	--driver-path		Location of the driver plugin
	--app-id			Application ID
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

	// create new executor
	logger := logger.NewStdLogger()
	e, err := NewExecutor(opts, logger)
	if err != nil {
		printAndDie(fmt.Sprintf("%s: %s", exe, err))
	}

	// register rpc
	rpc.Register(e)
	rpc.HandleHTTP()

	// bind rpc to executor sock
	sockname := opts.ExecutorSock
	os.Remove(sockname)
	l, err := net.Listen("unix", sockname)
	if err != nil {
		printAndDie(fmt.Sprintf("%s: %s", exe, err))
	}
	http.Serve(l, nil)
}
