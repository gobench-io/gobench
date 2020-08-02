package option

import (
	"flag"

	"github.com/gobench-io/gobench/logger"
)

type Options struct {
	Logger       logger.Logger
	AgentSock    string
	ExecutorSock string
	DriverPath   string // the plugin user wrote
}

func ConfigureOptions(fs *flag.FlagSet, args []string,
	printVersion, printHelp func()) (opts *Options, err error) {
	var (
		showVersion  bool
		showHelp     bool
		agentSock    string
		executorSock string
		driverPath   string
	)
	fs.BoolVar(&showVersion, "v", false, "Print version information")
	fs.BoolVar(&showVersion, "version", false, "Print version information")
	fs.BoolVar(&showHelp, "h", false, "Show this message")
	fs.BoolVar(&showHelp, "help", false, "Show this message")

	fs.StringVar(&agentSock, "agent-sock", "", "Socket of the agent")
	fs.StringVar(&executorSock, "executor-sock", "", "Socket for this executor")
	fs.StringVar(&driverPath, "driver-path", "", "Location of the driver plugin")

	if err = fs.Parse(args); err != nil {
		return
	}
	// show version
	// show help

	opts = &Options{}

	opts.AgentSock = agentSock
	opts.ExecutorSock = executorSock
	opts.DriverPath = driverPath

	return
}
