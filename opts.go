package main

import (
	"errors"
	"flag"
)

type mode string

// Modes of a server
const (
	Executor mode = "executor"
)

// Err messages
var (
	ErrInvalidFlags = errors.New("invalid flags")
)

type Options struct {
	Mode mode

	// executor mode
	AgentSock    string
	ExecutorSock string
	DriverPath   string // the plugin user wrote
	AppID        int
}

func ConfigureOptions(fs *flag.FlagSet, args []string, printVersion, printHelp func()) (
	opts *Options, err error,
) {
	var (
		showVersion bool
		showHelp    bool

		modeS string

		// executor mode
		agentSock    string
		executorSock string
		driverPath   string
		appID        int
	)
	fs.BoolVar(&showVersion, "v", false, "Print version information")
	fs.BoolVar(&showVersion, "version", false, "Print version information")
	fs.BoolVar(&showHelp, "h", false, "Show this message")
	fs.BoolVar(&showHelp, "help", false, "Show this message")

	fs.StringVar(&modeS, "mode", "master", "Operation mode of the program, either master, agent, or executor")

	fs.StringVar(&agentSock, "agent-sock", "", "Socket of the agent")
	fs.StringVar(&executorSock, "executor-sock", "", "Socket for this executor")
	fs.StringVar(&driverPath, "driver-path", "", "Location of the driver plugin")
	fs.IntVar(&appID, "app-id", -1, "Application ID")

	if err = fs.Parse(args); err != nil {
		return
	}

	if showHelp {
		printHelp()
		return
	}

	if showVersion {
		printVersion()
		return
	}

	mode := mode(modeS)

	if mode == Executor {
		if agentSock == "" || executorSock == "" || driverPath == "" || appID < 0 {
			err = ErrInvalidFlags
			return
		}

		opts = &Options{}

		opts.AgentSock = agentSock
		opts.ExecutorSock = executorSock
		opts.DriverPath = driverPath
		opts.AppID = appID

		return
	}

	err = ErrInvalidFlags

	return
}
