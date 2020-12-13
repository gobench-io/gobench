package master

import (
	"runtime"
	"time"
)

// Snapshot this
var numCores int
var maxProcs int

func init() {
	numCores = runtime.NumCPU()
	maxProcs = runtime.GOMAXPROCS(0)
}

// Varz outputs master information at /varz
type Varz struct {
	ID        string    `json:"server_id"`
	Version   string    `json:"version"`
	GitCommit string    `json:"git_commit,omitempty"`
	GoVersion string    `json:"go"`
	Start     time.Time `json:"start"`
	Now       time.Time `json:"now"`
	Uptime    string    `json:"uptime"`
	Mem       int64     `json:"mem"`
	Cores     int       `json:"cores"`
	MaxProcs  int       `json:"gomaxprocs"`
	CPU       float64   `json:"cpu"`
}

// Varz returns a Varz struct containing the server information.
func (m *Master) Varz() (*Varz, error) {
	varz := &Varz{
		ID:        m.id,
		Version:   m.version,
		GitCommit: m.gitCommit,
		GoVersion: m.goVersion,
		Start:     m.start,
		Now:       time.Now(),
		Cores:     numCores,
		MaxProcs:  maxProcs,
	}

	return varz, nil
}
