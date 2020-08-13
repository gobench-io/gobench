package driver

import (
	"time"

	"github.com/gobench-io/gobench/metrics"
	"github.com/mackerelio/go-osstat/loadavg"
)

// load average
const slLA1 string = "LA1"

// systemload report the current host system load like cpu, ram, and network
// status

// systemloadSetup setup the metrics for systemload
func (d *Driver) systemloadSetup() (err error) {
	group := metrics.Group{
		Name: "System Load",
		Graphs: []metrics.Graph{
			{
				Title: "Load average",
				Unit:  "%",
				Metrics: []metrics.Metric{
					{
						Title: slLA1,
						Type:  metrics.Gauge,
					},
				},
			},
		},
	}
	groups := []metrics.Group{
		group,
	}
	err = Setup(groups)
	return
}

// systemloadRun start collect the metrics
func (d *Driver) systemloadRun() (err error) {
	ch := make(chan interface{})
	freq := 1 * time.Second

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	for range ch {
		if la, err := loadavg.Get(); err != nil {
			Notify(slLA1, int64(la.Loadavg1*100))
		}
	}

	return
}
