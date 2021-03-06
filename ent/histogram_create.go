// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gobench-io/gobench/ent/histogram"
	"github.com/gobench-io/gobench/ent/metric"
)

// HistogramCreate is the builder for creating a Histogram entity.
type HistogramCreate struct {
	config
	mutation *HistogramMutation
	hooks    []Hook
}

// SetTime sets the time field.
func (hc *HistogramCreate) SetTime(i int64) *HistogramCreate {
	hc.mutation.SetTime(i)
	return hc
}

// SetCount sets the count field.
func (hc *HistogramCreate) SetCount(i int64) *HistogramCreate {
	hc.mutation.SetCount(i)
	return hc
}

// SetMin sets the min field.
func (hc *HistogramCreate) SetMin(i int64) *HistogramCreate {
	hc.mutation.SetMin(i)
	return hc
}

// SetMax sets the max field.
func (hc *HistogramCreate) SetMax(i int64) *HistogramCreate {
	hc.mutation.SetMax(i)
	return hc
}

// SetMean sets the mean field.
func (hc *HistogramCreate) SetMean(f float64) *HistogramCreate {
	hc.mutation.SetMean(f)
	return hc
}

// SetStddev sets the stddev field.
func (hc *HistogramCreate) SetStddev(f float64) *HistogramCreate {
	hc.mutation.SetStddev(f)
	return hc
}

// SetMedian sets the median field.
func (hc *HistogramCreate) SetMedian(f float64) *HistogramCreate {
	hc.mutation.SetMedian(f)
	return hc
}

// SetP75 sets the p75 field.
func (hc *HistogramCreate) SetP75(f float64) *HistogramCreate {
	hc.mutation.SetP75(f)
	return hc
}

// SetP95 sets the p95 field.
func (hc *HistogramCreate) SetP95(f float64) *HistogramCreate {
	hc.mutation.SetP95(f)
	return hc
}

// SetP99 sets the p99 field.
func (hc *HistogramCreate) SetP99(f float64) *HistogramCreate {
	hc.mutation.SetP99(f)
	return hc
}

// SetP999 sets the p999 field.
func (hc *HistogramCreate) SetP999(f float64) *HistogramCreate {
	hc.mutation.SetP999(f)
	return hc
}

// SetWID sets the wID field.
func (hc *HistogramCreate) SetWID(s string) *HistogramCreate {
	hc.mutation.SetWID(s)
	return hc
}

// SetMetricID sets the metric edge to Metric by id.
func (hc *HistogramCreate) SetMetricID(id int) *HistogramCreate {
	hc.mutation.SetMetricID(id)
	return hc
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (hc *HistogramCreate) SetNillableMetricID(id *int) *HistogramCreate {
	if id != nil {
		hc = hc.SetMetricID(*id)
	}
	return hc
}

// SetMetric sets the metric edge to Metric.
func (hc *HistogramCreate) SetMetric(m *Metric) *HistogramCreate {
	return hc.SetMetricID(m.ID)
}

// Mutation returns the HistogramMutation object of the builder.
func (hc *HistogramCreate) Mutation() *HistogramMutation {
	return hc.mutation
}

// Save creates the Histogram in the database.
func (hc *HistogramCreate) Save(ctx context.Context) (*Histogram, error) {
	var (
		err  error
		node *Histogram
	)
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistogramMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			node, err = hc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HistogramCreate) SaveX(ctx context.Context) *Histogram {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (hc *HistogramCreate) check() error {
	if _, ok := hc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New("ent: missing required field \"time\"")}
	}
	if _, ok := hc.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New("ent: missing required field \"count\"")}
	}
	if _, ok := hc.mutation.Min(); !ok {
		return &ValidationError{Name: "min", err: errors.New("ent: missing required field \"min\"")}
	}
	if _, ok := hc.mutation.Max(); !ok {
		return &ValidationError{Name: "max", err: errors.New("ent: missing required field \"max\"")}
	}
	if _, ok := hc.mutation.Mean(); !ok {
		return &ValidationError{Name: "mean", err: errors.New("ent: missing required field \"mean\"")}
	}
	if _, ok := hc.mutation.Stddev(); !ok {
		return &ValidationError{Name: "stddev", err: errors.New("ent: missing required field \"stddev\"")}
	}
	if _, ok := hc.mutation.Median(); !ok {
		return &ValidationError{Name: "median", err: errors.New("ent: missing required field \"median\"")}
	}
	if _, ok := hc.mutation.P75(); !ok {
		return &ValidationError{Name: "p75", err: errors.New("ent: missing required field \"p75\"")}
	}
	if _, ok := hc.mutation.P95(); !ok {
		return &ValidationError{Name: "p95", err: errors.New("ent: missing required field \"p95\"")}
	}
	if _, ok := hc.mutation.P99(); !ok {
		return &ValidationError{Name: "p99", err: errors.New("ent: missing required field \"p99\"")}
	}
	if _, ok := hc.mutation.P999(); !ok {
		return &ValidationError{Name: "p999", err: errors.New("ent: missing required field \"p999\"")}
	}
	if _, ok := hc.mutation.WID(); !ok {
		return &ValidationError{Name: "wID", err: errors.New("ent: missing required field \"wID\"")}
	}
	return nil
}

func (hc *HistogramCreate) sqlSave(ctx context.Context) (*Histogram, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hc *HistogramCreate) createSpec() (*Histogram, *sqlgraph.CreateSpec) {
	var (
		_node = &Histogram{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: histogram.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: histogram.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldTime,
		})
		_node.Time = value
	}
	if value, ok := hc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := hc.mutation.Min(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMin,
		})
		_node.Min = value
	}
	if value, ok := hc.mutation.Max(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: histogram.FieldMax,
		})
		_node.Max = value
	}
	if value, ok := hc.mutation.Mean(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMean,
		})
		_node.Mean = value
	}
	if value, ok := hc.mutation.Stddev(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldStddev,
		})
		_node.Stddev = value
	}
	if value, ok := hc.mutation.Median(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldMedian,
		})
		_node.Median = value
	}
	if value, ok := hc.mutation.P75(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP75,
		})
		_node.P75 = value
	}
	if value, ok := hc.mutation.P95(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP95,
		})
		_node.P95 = value
	}
	if value, ok := hc.mutation.P99(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP99,
		})
		_node.P99 = value
	}
	if value, ok := hc.mutation.P999(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: histogram.FieldP999,
		})
		_node.P999 = value
	}
	if value, ok := hc.mutation.WID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: histogram.FieldWID,
		})
		_node.WID = value
	}
	if nodes := hc.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   histogram.MetricTable,
			Columns: []string{histogram.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HistogramCreateBulk is the builder for creating a bulk of Histogram entities.
type HistogramCreateBulk struct {
	config
	builders []*HistogramCreate
}

// Save creates the Histogram entities in the database.
func (hcb *HistogramCreateBulk) Save(ctx context.Context) ([]*Histogram, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Histogram, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HistogramMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (hcb *HistogramCreateBulk) SaveX(ctx context.Context) []*Histogram {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
