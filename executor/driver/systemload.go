package driver

import (
	"context"
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
				Title:   "Network transmit",
				Unit:    "KiB/s",
				Metrics: []metrics.Metric{},
			},
			{
				Title:   "Network receive",
				Unit:    "KiB/s",
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
			Title: ns.Name + " transmit",
			Type:  metrics.Gauge,
		})
		group.Graphs[2].Metrics = append(group.Graphs[2].Metrics, metrics.Metric{
			Title: ns.Name + " receive",
			Type:  metrics.Gauge,
		})
	}

	groups := []metrics.Group{
		group,
	}
	err = Setup(groups)
	return
}

type rtxBytes struct {
	rxBytes, txBytes uint64
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

	nssPre := make(map[string]rtxBytes)
	nssPreTime := time.Now()

	for range ch {
		if la, err := loadavg.Get(); err == nil {
			la1 := int64(la.Loadavg1 * 100)
			Notify(slLA1, la1)
		}

		if nss, err := network.Get(); err == nil {
			now := time.Now()
			for _, ns := range nss {
				prv, ok := nssPre[ns.Name]
				if ok {
					diffTime := now.Sub(nssPreTime).Seconds()

					tx := float64(ns.TxBytes-prv.txBytes) / diffTime // Bps
					rx := float64(ns.RxBytes-prv.rxBytes) / diffTime // Bps
					Notify(ns.Name+" transmit", int64(tx/1000))      // KBps
					Notify(ns.Name+" receive", int64(rx/1000))       // KBps
				}
				// update prv values
				nssPre[ns.Name] = rtxBytes{
					rxBytes: ns.RxBytes,
					txBytes: ns.TxBytes,
				}
			}
			nssPreTime = now
		}
	}

	return
}
