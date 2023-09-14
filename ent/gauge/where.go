// Code generated by ent, DO NOT EDIT.

package gauge

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gobench-io/gobench/v2/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Gauge {
	return predicate.Gauge(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Gauge {
	return predicate.Gauge(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Gauge {
	return predicate.Gauge(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Gauge {
	return predicate.Gauge(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Gauge {
	return predicate.Gauge(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Gauge {
	return predicate.Gauge(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Gauge {
	return predicate.Gauge(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Gauge {
	return predicate.Gauge(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Gauge {
	return predicate.Gauge(sql.FieldLTE(FieldID, id))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldEQ(FieldTime, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldEQ(FieldValue, v))
}

// WID applies equality check predicate on the "wID" field. It's identical to WIDEQ.
func WID(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldEQ(FieldWID, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldLTE(FieldTime, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int64) predicate.Gauge {
	return predicate.Gauge(sql.FieldLTE(FieldValue, v))
}

// WIDEQ applies the EQ predicate on the "wID" field.
func WIDEQ(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldEQ(FieldWID, v))
}

// WIDNEQ applies the NEQ predicate on the "wID" field.
func WIDNEQ(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldNEQ(FieldWID, v))
}

// WIDIn applies the In predicate on the "wID" field.
func WIDIn(vs ...string) predicate.Gauge {
	return predicate.Gauge(sql.FieldIn(FieldWID, vs...))
}

// WIDNotIn applies the NotIn predicate on the "wID" field.
func WIDNotIn(vs ...string) predicate.Gauge {
	return predicate.Gauge(sql.FieldNotIn(FieldWID, vs...))
}

// WIDGT applies the GT predicate on the "wID" field.
func WIDGT(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldGT(FieldWID, v))
}

// WIDGTE applies the GTE predicate on the "wID" field.
func WIDGTE(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldGTE(FieldWID, v))
}

// WIDLT applies the LT predicate on the "wID" field.
func WIDLT(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldLT(FieldWID, v))
}

// WIDLTE applies the LTE predicate on the "wID" field.
func WIDLTE(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldLTE(FieldWID, v))
}

// WIDContains applies the Contains predicate on the "wID" field.
func WIDContains(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldContains(FieldWID, v))
}

// WIDHasPrefix applies the HasPrefix predicate on the "wID" field.
func WIDHasPrefix(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldHasPrefix(FieldWID, v))
}

// WIDHasSuffix applies the HasSuffix predicate on the "wID" field.
func WIDHasSuffix(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldHasSuffix(FieldWID, v))
}

// WIDEqualFold applies the EqualFold predicate on the "wID" field.
func WIDEqualFold(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldEqualFold(FieldWID, v))
}

// WIDContainsFold applies the ContainsFold predicate on the "wID" field.
func WIDContainsFold(v string) predicate.Gauge {
	return predicate.Gauge(sql.FieldContainsFold(FieldWID, v))
}

// HasMetric applies the HasEdge predicate on the "metric" edge.
func HasMetric() predicate.Gauge {
	return predicate.Gauge(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MetricTable, MetricColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetricWith applies the HasEdge predicate on the "metric" edge with a given conditions (other predicates).
func HasMetricWith(preds ...predicate.Metric) predicate.Gauge {
	return predicate.Gauge(func(s *sql.Selector) {
		step := newMetricStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Gauge) predicate.Gauge {
	return predicate.Gauge(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Gauge) predicate.Gauge {
	return predicate.Gauge(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Gauge) predicate.Gauge {
	return predicate.Gauge(func(s *sql.Selector) {
		p(s.Not())
	})
}
