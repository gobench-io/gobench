// executor is started from agent as a separated process
package main

import "github.com/gobench-io/gobench/logger"

type Options struct {
	Logger       logger.Logger
	AgentSock    string
	ExecutorSock string
	DriverPath   string // the plugin user wrote
}

// run the program with args
// --agent-sock
// --executor-sock
// --driver-path
func main() {
}
