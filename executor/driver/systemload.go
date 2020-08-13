package driver

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/metrics"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/network"
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
				Unit:  "*100",
				Metrics: []metrics.Metric{
					{
						Title: slLA1,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title:   "Network transmit (bytes)",
				Unit:    "Bytes",
				Metrics: []metrics.Metric{},
			},
		},
	}

	nss, err := network.Get()
	if err != nil {
		return
	}
	for _, ns := range nss {
		group.Graphs[1].Metrics = append(group.Graphs[1].Metrics, metrics.Metric{
			Title: ns.Name,
			Type:  metrics.Gauge,
		})
	}

	groups := []metrics.Group{
		group,
	}
	err = Setup(groups)
	return
}

// systemloadRun start collect the metrics
// todo: handle cancel or finish signal
func (d *Driver) systemloadRun(ctx context.Context) (err error) {
	ch := make(chan interface{})
	freq := 1 * time.Second

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	for range ch {
		if la, err := loadavg.Get(); err == nil {
			la1 := int64(la.Loadavg1 * 100)
			log.Printf("la1: %d\n", la1)
			Notify(slLA1, la1)
		}
	}

	return
}
