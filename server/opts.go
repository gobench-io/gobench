package server

import (
	"flag"
	"fmt"
)

// Options block for gobench server
type Options struct {
	isWorker    bool
	Master      bool
	Addr        string
	Port        int
	ClusterPort int    // master address to solicit and connect
	Route       string // master address for worker to connect to
}

func setBaselineOptions(opts *Options) {
	if opts.Addr == "" {
		opts.Addr = DEFAULT_HOST
	}

	if !opts.isWorker {
		opts.Master = true
		if opts.Port == 0 {
			opts.Port = DEFAULT_PORT
		}
		if opts.ClusterPort == 0 {
			opts.ClusterPort = DEFAULT_CLUSTER_PORT
		}
	} else {
		opts.Master = false
		if opts.Route == "" {
			opts.Route = fmt.Sprintf("0.0.0.0:%d", DEFAULT_CLUSTER_PORT)
		}
	}
}

// ConfigureOptions accepts a flag set and augments it with gobench specific
// flags. On success, an options structure is returned configured based on the
// selected flags
func ConfigureOptions(fs *flag.FlagSet, args []string, printVersion, printHelp func()) (*Options, error) {
	return nil, nil
}
