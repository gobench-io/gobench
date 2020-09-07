package executor

import (
	"context"
	"fmt"

	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/pb"

	gometrics "github.com/rcrowley/go-metrics"
)

// Setup is used for the driver to report the metrics that it will generate
func Setup(groups []metrics.Group) error {
	executorInst := getExecutor()

	ctx := context.TODO()

	units := make(map[string]unit)

	executorInst.mu.Lock()
	defer executorInst.mu.Unlock()

	for _, group := range groups {
		// create a new group if not existed
		egroup, err := executorInst.rc.FindCreateGroup(ctx, &pb.FCGroupReq{
			AppID: int64(executorInst.appID),
			Name:  group.Name,
		})
		if err != nil {
			return fmt.Errorf("failed create group: %v", err)
		}

		for _, graph := range group.Graphs {
			// create new graph if not existed
			egraph, err := executorInst.rc.FindCreateGraph(ctx, &pb.FCGraphReq{
				AppID:   int64(executorInst.appID),
				Title:   graph.Title,
				Unit:    graph.Unit,
				GroupID: egroup.Id,
			})
			if err != nil {
				return fmt.Errorf("failed create graph: %v", err)
			}

			for _, m := range graph.Metrics {
				// create new metric if not existed
				emetric, err := executorInst.rc.FindCreateMetric(ctx, &pb.FCMetricReq{
					AppID:   int64(executorInst.appID),
					Title:   m.Title,
					Type:    string(m.Type),
					GraphID: egraph.Id,
				})
				if err != nil {
					return fmt.Errorf("failed create metric: %v", err)
				}

				// counter type
				if m.Type == metrics.Counter {
					c := gometrics.NewCounter()
					if err := gometrics.Register(m.Title, c); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}

					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						c:        c,
					}
				}

				if m.Type == metrics.Histogram {
					s := gometrics.NewExpDecaySample(1028, 0.015)
					h := gometrics.NewHistogram(s)
					if err := gometrics.Register(m.Title, h); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}
					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						h:        h,
					}
				}

				if m.Type == metrics.Gauge {
					g := gometrics.NewGauge()
					if err := gometrics.Register(m.Title, g); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}
					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						g:        g,
					}
				}
			}
		}
	}

	// aggregate units
	for k, v := range units {
		executorInst.units[k] = v
	}

	return nil
}

// Notify saves the id with value into metrics which later save to database
// Return error when the title is not found from the metric list.
// The not found error may occur because
// a. The title has never ever register before
// b. The session is cancel but the scenario does not handle the ctx.Done signal
func Notify(title string, value int64) error {
	executorInst := getExecutor()

	executorInst.mu.Lock()
	defer executorInst.mu.Unlock()

	u, ok := executorInst.units[title]
	if !ok {
		executorInst.logger.Infow("metric not found", "title", title)
		return ErrIDNotFound
	}

	if u.Type == metrics.Counter {
		u.c.Inc(value)
	}

	if u.Type == metrics.Histogram {
		u.h.Update(value)
	}

	if u.Type == metrics.Gauge {
		u.g.Update(value)
	}

	return nil
}
