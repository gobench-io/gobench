package master

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gobench-io/gobench/v2/pse"
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
	Mem       uint64    `json:"mem"`
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
	var pcpu float64
	var mem uint64
	if err := pse.ProcUsage(&pcpu, &mem); err != nil {
		return varz, err
	}

	varz.CPU = pcpu
	varz.Mem = mem

	varz.Uptime = myUptime(varz.Now.Sub(varz.Start))

	return varz, nil
}

// myUptime returns duration in year-day-hour-min-second format
// inspired by nats-server
func myUptime(d time.Duration) string {
	tsecs := d / time.Second
	tmins := tsecs / 60
	thrs := tmins / 60
	tdays := thrs / 24
	tyrs := tdays / 365

	if tyrs > 0 {
		return fmt.Sprintf("%dy%dd%dh%dm%ds", tyrs, tdays%365, thrs%24, tmins%60, tsecs%60)
	}
	if tdays > 0 {
		return fmt.Sprintf("%dd%dh%dm%ds", tdays, thrs%24, tmins%60, tsecs%60)
	}
	if thrs > 0 {
		return fmt.Sprintf("%dh%dm%ds", thrs, tmins%60, tsecs%60)
	}
	if tmins > 0 {
		return fmt.Sprintf("%dm%ds", tmins, tsecs%60)
	}
	return fmt.Sprintf("%ds", tsecs)
}
