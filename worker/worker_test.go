package worker

import (
	"context"
	"testing"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/metrics"
	gometrics "github.com/rcrowley/go-metrics"
	"github.com/stretchr/testify/assert"
)

func loadValidPlugin(w *Worker) error {
	so := "./script/valid.so"
	return w.Load(so)
}

type nilLog struct{}

func (l *nilLog) Counter(ctx context.Context, id, title string, time, c int64) error {
	return nil
}

func (l *nilLog) Histogram(ctx context.Context, id, title string, time int64, h gometrics.Histogram) error {
	return nil
}

func (l *nilLog) Gauge(ctx context.Context, id, title string, time int64, g int64) error {
	return nil
}

func (l *nilLog) FindCreateGroup(ctx context.Context, mg metrics.Group) (
	*ent.Group, error,
) {
	return nil, nil
}
func (l *nilLog) FindCreateGraph(ctx context.Context, mgraph metrics.Graph, groupID int) (
	*ent.Graph, error,
) {
	return nil, nil
}

func (l *nilLog) FindCreateMetric(ctx context.Context, mmetric metrics.Metric, graphID int) (
	*ent.Metric, error,
) {
	return nil, nil
}

func newNilLog() metricLogger {
	return &nilLog{}
}

func TestNew(t *testing.T) {
	n1, err := NewWorker(newNilLog())
	assert.Nil(t, err)
	n2, err := NewWorker(newNilLog())
	assert.Nil(t, err)

	assert.Equal(t, n1, n2)

	assert.Equal(t, n1.status, Idle)
	assert.Nil(t, n1.cancel)
	assert.Equal(t, n1.units, make(map[string]unit))

	assert.False(t, n1.Running())
}

func TestLoadPlugin(t *testing.T) {
	n, _ := NewWorker(newNilLog())
	so := "./script/valid.so"
	assert.Nil(t, n.Load(so))
	assert.NotNil(t, n.vus)
	assert.False(t, n.Running())
}

// func TestRunPlugin(t *testing.T) {
// 	n, _ := NewWorker(newNilLog())
// 	assert.Nil(t, loadValidPlugin(n))
// 	assert.Nil(t, n.Run())
// 	assert.False(t, n.Running())
// }
