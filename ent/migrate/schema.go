// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ApplicationsColumns holds the columns for the "applications" table.
	ApplicationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "finished_at", Type: field.TypeTime, Nullable: true},
		{Name: "scenario", Type: field.TypeString, Size: 2147483647},
	}
	// ApplicationsTable holds the schema information for the "applications" table.
	ApplicationsTable = &schema.Table{
		Name:        "applications",
		Columns:     ApplicationsColumns,
		PrimaryKey:  []*schema.Column{ApplicationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CountersColumns holds the columns for the "counters" table.
	CountersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "time", Type: field.TypeInt64},
		{Name: "count", Type: field.TypeInt64},
		{Name: "metric_counters", Type: field.TypeInt, Nullable: true},
	}
	// CountersTable holds the schema information for the "counters" table.
	CountersTable = &schema.Table{
		Name:       "counters",
		Columns:    CountersColumns,
		PrimaryKey: []*schema.Column{CountersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "counters_metrics_counters",
				Columns: []*schema.Column{CountersColumns[3]},

				RefColumns: []*schema.Column{MetricsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GaugesColumns holds the columns for the "gauges" table.
	GaugesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "time", Type: field.TypeInt64},
		{Name: "value", Type: field.TypeInt64},
		{Name: "metric_gauges", Type: field.TypeInt, Nullable: true},
	}
	// GaugesTable holds the schema information for the "gauges" table.
	GaugesTable = &schema.Table{
		Name:       "gauges",
		Columns:    GaugesColumns,
		PrimaryKey: []*schema.Column{GaugesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "gauges_metrics_gauges",
				Columns: []*schema.Column{GaugesColumns[3]},

				RefColumns: []*schema.Column{MetricsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GraphsColumns holds the columns for the "graphs" table.
	GraphsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "unit", Type: field.TypeString},
		{Name: "group_graphs", Type: field.TypeInt, Nullable: true},
	}
	// GraphsTable holds the schema information for the "graphs" table.
	GraphsTable = &schema.Table{
		Name:       "graphs",
		Columns:    GraphsColumns,
		PrimaryKey: []*schema.Column{GraphsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "graphs_groups_graphs",
				Columns: []*schema.Column{GraphsColumns[3]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "application_groups", Type: field.TypeInt, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "groups_applications_groups",
				Columns: []*schema.Column{GroupsColumns[2]},

				RefColumns: []*schema.Column{ApplicationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HistogramsColumns holds the columns for the "histograms" table.
	HistogramsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "time", Type: field.TypeInt64},
		{Name: "count", Type: field.TypeInt64},
		{Name: "min", Type: field.TypeInt64},
		{Name: "max", Type: field.TypeInt64},
		{Name: "mean", Type: field.TypeFloat64},
		{Name: "stddev", Type: field.TypeFloat64},
		{Name: "median", Type: field.TypeFloat64},
		{Name: "p75", Type: field.TypeFloat64},
		{Name: "p95", Type: field.TypeFloat64},
		{Name: "p99", Type: field.TypeFloat64},
		{Name: "p999", Type: field.TypeFloat64},
		{Name: "metric_histograms", Type: field.TypeInt, Nullable: true},
	}
	// HistogramsTable holds the schema information for the "histograms" table.
	HistogramsTable = &schema.Table{
		Name:       "histograms",
		Columns:    HistogramsColumns,
		PrimaryKey: []*schema.Column{HistogramsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "histograms_metrics_histograms",
				Columns: []*schema.Column{HistogramsColumns[12]},

				RefColumns: []*schema.Column{MetricsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MetricsColumns holds the columns for the "metrics" table.
	MetricsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeString},
		{Name: "graph_metrics", Type: field.TypeInt, Nullable: true},
	}
	// MetricsTable holds the schema information for the "metrics" table.
	MetricsTable = &schema.Table{
		Name:       "metrics",
		Columns:    MetricsColumns,
		PrimaryKey: []*schema.Column{MetricsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "metrics_graphs_metrics",
				Columns: []*schema.Column{MetricsColumns[3]},

				RefColumns: []*schema.Column{GraphsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApplicationsTable,
		CountersTable,
		GaugesTable,
		GraphsTable,
		GroupsTable,
		HistogramsTable,
		MetricsTable,
	}
)

func init() {
	CountersTable.ForeignKeys[0].RefTable = MetricsTable
	GaugesTable.ForeignKeys[0].RefTable = MetricsTable
	GraphsTable.ForeignKeys[0].RefTable = GroupsTable
	GroupsTable.ForeignKeys[0].RefTable = ApplicationsTable
	HistogramsTable.ForeignKeys[0].RefTable = MetricsTable
	MetricsTable.ForeignKeys[0].RefTable = GraphsTable
}
