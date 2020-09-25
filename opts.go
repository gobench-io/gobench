package main

import (
	"errors"
	"flag"
	"os"
	"os/user"
	"path/filepath"
)

type mode string

// Modes of a server
const (
	Master mode = "master"
	Agent  mode = "agent"
)

// Err messages
var (
	ErrInvalidFlags = errors.New("invalid flags")
)

type Options struct {
	Mode    mode
	Program string

	// master, agent mode
	ClusterPort   int
	AdminPassword string

	// master mode
	Port   int
	DbPath string
	Dir    string

	// agent mode
	Route string
}

// func (o Options) String() string {
// 	return fmt.Sprintf("mode: %s\n", string(o.Mode))
// }

func ConfigureOptions(fs *flag.FlagSet, args []string, printVersion, printHelp func()) (
	*Options, error,
) {
	var err error

	var (
		showVersion bool
		showHelp    bool

		modeS string

		// master mode
		port          int
		dbPath        string
		adminPassword string
		dir           string

		// agent mode
		route       string
		clusterPort int
	)
	// gen default working dir
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	defDir := filepath.Join(u.HomeDir, ".gobench")

	fs.BoolVar(&showVersion, "v", false, "Print version information")
	fs.BoolVar(&showVersion, "version", false, "Print version information")
	fs.BoolVar(&showHelp, "h", false, "Show this message")
	fs.BoolVar(&showHelp, "help", false, "Show this message")

	fs.StringVar(&modeS, "mode", "master", "Operation mode of the program, either master, agent, or executor")

	// master
	fs.IntVar(&port, "p", DEFAULT_PORT, "Port of the master server.")
	fs.IntVar(&port, "port", DEFAULT_PORT, "Port of the master server.")
	fs.StringVar(&dbPath, "db", "", "Name of the database.")
	fs.StringVar(&adminPassword, "admin-password", "", "Admin password to login to web dashboard")
	fs.StringVar(&dir, "dir", defDir, "Working directory (default: ${HOME}). The result database and logs will be stored on this folder.")

	// agent
	fs.IntVar(&clusterPort, "clusterPort", DEFAULT_CLUSTER_PORT, "Cluster port to solicit and connect.")

	// master + agent
	fs.StringVar(&route, "route", "", "Master address to solicit routes.")

	program := args[0]
	if err = fs.Parse(args[1:]); err != nil {
		return nil, err
	}

	if showHelp {
		printHelp()
		return nil, nil
	}

	if showVersion {
		printVersion()
		return nil, nil
	}

	opts := &Options{
		Mode:    mode(modeS),
		Program: program,
	}

	if opts.Mode == Master {
		if dbPath == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return opts, err
			}
			dbPath = filepath.Join(home, DEFAULT_DB_NAME)
		}
		opts.Port = port
		opts.ClusterPort = clusterPort
		opts.DbPath = dbPath
		opts.AdminPassword = adminPassword
		opts.Dir = dir
		return opts, nil
	}

	if opts.Mode == Agent {
		if route == "" {
			return nil, errors.New("agent must have route to a master")
		}
		opts.Route = route
		opts.ClusterPort = clusterPort
		return opts, nil
	}

	err = errors.New("mode must be either master, agent, or executor")

	return nil, err
}
