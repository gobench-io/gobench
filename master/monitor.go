package master

import "time"

// Varz outputs master information at /varz
type Varz struct {
	ID        string    `json:"server_id"`
	Version   string    `json:"version"`
	GitCommit int       `json:"git_commit,omitempty`
	GoVersion string    `json:"go"`
	Start     time.Time `json:"start"`
	Now       time.Time `json:"now"`
	Uptime    string    `json:"uptime"`
	Mem       int64     `json:"mem"`
	Cores     int       `json:"cores"`
	MaxProcs  int       `json:"gomaxprocs`
	CPU       float64   `json:"cpu"`
}

// Varz returns a Varz struct containing the server information.
func (m *Master) Varz() (*Varz, error) {
}
