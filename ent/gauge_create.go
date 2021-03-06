// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gobench-io/gobench/ent/gauge"
	"github.com/gobench-io/gobench/ent/metric"
)

// GaugeCreate is the builder for creating a Gauge entity.
type GaugeCreate struct {
	config
	mutation *GaugeMutation
	hooks    []Hook
}

// SetTime sets the time field.
func (gc *GaugeCreate) SetTime(i int64) *GaugeCreate {
	gc.mutation.SetTime(i)
	return gc
}

// SetValue sets the value field.
func (gc *GaugeCreate) SetValue(i int64) *GaugeCreate {
	gc.mutation.SetValue(i)
	return gc
}

// SetWID sets the wID field.
func (gc *GaugeCreate) SetWID(s string) *GaugeCreate {
	gc.mutation.SetWID(s)
	return gc
}

// SetMetricID sets the metric edge to Metric by id.
func (gc *GaugeCreate) SetMetricID(id int) *GaugeCreate {
	gc.mutation.SetMetricID(id)
	return gc
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (gc *GaugeCreate) SetNillableMetricID(id *int) *GaugeCreate {
	if id != nil {
		gc = gc.SetMetricID(*id)
	}
	return gc
}

// SetMetric sets the metric edge to Metric.
func (gc *GaugeCreate) SetMetric(m *Metric) *GaugeCreate {
	return gc.SetMetricID(m.ID)
}

// Mutation returns the GaugeMutation object of the builder.
func (gc *GaugeCreate) Mutation() *GaugeMutation {
	return gc.mutation
}

// Save creates the Gauge in the database.
func (gc *GaugeCreate) Save(ctx context.Context) (*Gauge, error) {
	var (
		err  error
		node *Gauge
	)
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GaugeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GaugeCreate) SaveX(ctx context.Context) *Gauge {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (gc *GaugeCreate) check() error {
	if _, ok := gc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New("ent: missing required field \"time\"")}
	}
	if _, ok := gc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New("ent: missing required field \"value\"")}
	}
	if _, ok := gc.mutation.WID(); !ok {
		return &ValidationError{Name: "wID", err: errors.New("ent: missing required field \"wID\"")}
	}
	return nil
}

func (gc *GaugeCreate) sqlSave(ctx context.Context) (*Gauge, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gc *GaugeCreate) createSpec() (*Gauge, *sqlgraph.CreateSpec) {
	var (
		_node = &Gauge{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gauge.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gauge.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: gauge.FieldTime,
		})
		_node.Time = value
	}
	if value, ok := gc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: gauge.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := gc.mutation.WID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gauge.FieldWID,
		})
		_node.WID = value
	}
	if nodes := gc.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gauge.MetricTable,
			Columns: []string{gauge.MetricColumn},
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

// GaugeCreateBulk is the builder for creating a bulk of Gauge entities.
type GaugeCreateBulk struct {
	config
	builders []*GaugeCreate
}

// Save creates the Gauge entities in the database.
func (gcb *GaugeCreateBulk) Save(ctx context.Context) ([]*Gauge, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Gauge, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GaugeMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (gcb *GaugeCreateBulk) SaveX(ctx context.Context) []*Gauge {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
