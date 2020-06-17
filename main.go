package main

import (
	"log"

	"github.com/gobench-io/gobench/server"
	"github.com/gobench-io/gobench/web"
)

// gobench -p 3000 -master -cluster localhost:3001
// gobench -worker -cluster localhost:5001 -route localhost:3001

var usageStr = `
Usage: gobench [options]

Server Options:
	-a, --addr <host>		Bind to host address (default: 0.0.0.0)
	-p, --port <port>		Use port for web client (default: 6891).
							Worker does not need this option.
	-m, --master			Run server as master mode
	-w, --worker 			Run server as worker mode

Cluster Options
		--cluster <port>	Cluster port to solicit and connect (default: 6890)
							Master is required to have this option
		--route <host:port> The master address to solicit routes.
							Every worker must have this option sothat worker can
							connect to a master

Common Options:
	-h, --help				Show this message
	-v, --version			Show version
`

func main() {
	server, _ := server.New()

	if err := server.Start(); err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	web.Serve(server, 3001)
}
