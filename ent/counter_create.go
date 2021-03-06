// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gobench-io/gobench/ent/counter"
	"github.com/gobench-io/gobench/ent/metric"
)

// CounterCreate is the builder for creating a Counter entity.
type CounterCreate struct {
	config
	mutation *CounterMutation
	hooks    []Hook
}

// SetTime sets the time field.
func (cc *CounterCreate) SetTime(i int64) *CounterCreate {
	cc.mutation.SetTime(i)
	return cc
}

// SetCount sets the count field.
func (cc *CounterCreate) SetCount(i int64) *CounterCreate {
	cc.mutation.SetCount(i)
	return cc
}

// SetWID sets the wID field.
func (cc *CounterCreate) SetWID(s string) *CounterCreate {
	cc.mutation.SetWID(s)
	return cc
}

// SetMetricID sets the metric edge to Metric by id.
func (cc *CounterCreate) SetMetricID(id int) *CounterCreate {
	cc.mutation.SetMetricID(id)
	return cc
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (cc *CounterCreate) SetNillableMetricID(id *int) *CounterCreate {
	if id != nil {
		cc = cc.SetMetricID(*id)
	}
	return cc
}

// SetMetric sets the metric edge to Metric.
func (cc *CounterCreate) SetMetric(m *Metric) *CounterCreate {
	return cc.SetMetricID(m.ID)
}

// Mutation returns the CounterMutation object of the builder.
func (cc *CounterCreate) Mutation() *CounterMutation {
	return cc.mutation
}

// Save creates the Counter in the database.
func (cc *CounterCreate) Save(ctx context.Context) (*Counter, error) {
	var (
		err  error
		node *Counter
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CounterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CounterCreate) SaveX(ctx context.Context) *Counter {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *CounterCreate) check() error {
	if _, ok := cc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New("ent: missing required field \"time\"")}
	}
	if _, ok := cc.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New("ent: missing required field \"count\"")}
	}
	if _, ok := cc.mutation.WID(); !ok {
		return &ValidationError{Name: "wID", err: errors.New("ent: missing required field \"wID\"")}
	}
	return nil
}

func (cc *CounterCreate) sqlSave(ctx context.Context) (*Counter, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CounterCreate) createSpec() (*Counter, *sqlgraph.CreateSpec) {
	var (
		_node = &Counter{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: counter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: counter.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldTime,
		})
		_node.Time = value
	}
	if value, ok := cc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: counter.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := cc.mutation.WID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: counter.FieldWID,
		})
		_node.WID = value
	}
	if nodes := cc.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   counter.MetricTable,
			Columns: []string{counter.MetricColumn},
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

// CounterCreateBulk is the builder for creating a bulk of Counter entities.
type CounterCreateBulk struct {
	config
	builders []*CounterCreate
}

// Save creates the Counter entities in the database.
func (ccb *CounterCreateBulk) Save(ctx context.Context) ([]*Counter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Counter, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CounterMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ccb *CounterCreateBulk) SaveX(ctx context.Context) []*Counter {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
