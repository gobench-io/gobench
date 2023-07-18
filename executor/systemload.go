package executor

import (
	"context"
	"time"

	"github.com/gobench-io/gobench/v2/executor/metrics"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
)

const slLA1 string = "LA1"
const cpuUser string = "CPU"
const ramUsing string = "RAM"

// systemload report the current host system load like cpu, ram, and network
// status

// systemloadSetup setup the metrics for systemload
func (e *Executor) systemloadSetup() (err error) {
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
				Title: "CPU",
				Unit:  "%",
				Metrics: []metrics.Metric{
					{
						Title: cpuUser,
						Type:  metrics.Gauge,
					},
				},
			},
			{
				Title: "RAM",
				Unit:  "%",
				Metrics: []metrics.Metric{
					{
						Title: ramUsing,
						Type:  metrics.Gauge,
					},
				},
			},
		},
	}

	// generate network metrics
	nss, err := network.Get()
	if err != nil {
		return
	}
	txMetrics := make([]metrics.Metric, 0, len(nss))
	rxMetrics := make([]metrics.Metric, 0, len(nss))
	for _, ns := range nss {
		// txMetrics = append(txMetrics, metrics.Metric{
		txMetrics = append(txMetrics, metrics.Metric{
			Title: ns.Name + " transmit",
			Type:  metrics.Gauge,
		})
		rxMetrics = append(rxMetrics, metrics.Metric{
			Title: ns.Name + " receive",
			Type:  metrics.Gauge,
		})
	}
	group.Graphs = append(group.Graphs, metrics.Graph{
		Title:   "Network transmit",
		Unit:    "KiB/s",
		Metrics: txMetrics,
	})
	group.Graphs = append(group.Graphs, metrics.Graph{
		Title:   "Network receive",
		Unit:    "KiB/s",
		Metrics: rxMetrics,
	})

	// finally
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
func (e *Executor) systemloadRun(ctx context.Context) (err error) {
	ch := make(chan interface{})
	freq := 1 * time.Second

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	nssPre := make(map[string]rtxBytes)
	nssPreTime := time.Now()

	var cpuPre *cpu.Stats
	cpuPreTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			e.logger.Infow("systemloadRun canceled")
			return nil

		case <-ch:
			// load average
			if la, err := loadavg.Get(); err == nil {
				la1 := int64(la.Loadavg1 * 100)
				Notify(slLA1, la1)
			}

			// network status
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

			if cpuNow, err := cpu.Get(); err == nil {
				now := time.Now()
				diffTime := now.Sub(cpuPreTime).Seconds()
				if cpuPre != nil {
					user := float64(cpuNow.User-cpuPre.User) / diffTime
					Notify(cpuUser, int64(user))
				}
				// update prv values
				cpuPre = cpuNow
				cpuPreTime = now
			}

			if mem, err := memory.Get(); err == nil {
				r := float64(mem.Used) / float64(mem.Total) * 100.0
				Notify(ramUsing, int64(r))
			}
		}
	}
}
