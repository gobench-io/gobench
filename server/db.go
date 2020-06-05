package gobench

import (
	"context"
	"fmt"
	"log"

	"github.com/gobench-io/gobench/ent"
	_ "github.com/mattn/go-sqlite3"

	entApplication "github.com/gobench-io/gobench/ent/application"
	entGroup "github.com/gobench-io/gobench/ent/group"
	entMetric "github.com/gobench-io/gobench/ent/metric"

	"github.com/gobench-io/gobench/metrics"
)

// func dSetupTables(db *sql.DB, metrics map[string]Unit) error {
// 	table := "value"

// 	// counter: id, time, count
// 	// historgram: id, time, count, min, max, mean, stddev, median, 75, 95, 99, 99.9
// 	// gauge: id, time, value
// 	sqlStmt := "create table if not exists '" + table + "' (id integer not null primary key autoincrement, metricId text, time time, count integer, min integer, max integer, mean real, stddev real, median real, '75' real, '95' real, '99' real, '99.9' real, value integer);"

// 	if _, err := db.Exec(sqlStmt); err != nil {
// 		return err
// 	}

// 	return nil
// }

func (c *Collect) setupDb(filename string) error {
	log.Printf("gobench result will be save in %s\n", filename)

	client, err := ent.Open(
		"sqlite3",
		filename+"?mode=rwc&cache=shared&&_busy_timeout=9999999&_fk=1")

	if err != nil {
		return fmt.Errorf("failed openning sqlite3 connection: %v", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	c.dbFilename = filename
	c.DB = client

	return nil
}

func (c *Collect) createApplication(ctx context.Context) error {
	_, err := c.DB.Application.
		Create().
		SetName(c.name).
		SetStatus(string(c.status)).
		Save(ctx)
	return err
}

func (c *Collect) updateApplication(ctx context.Context) error {
	_, err := c.DB.Application.
		Update().
		Where(
			entApplication.Name(c.name),
		).
		SetStatus(string(c.status)).
		SetFinishedAt(c.finishedAt).
		Save(ctx)

	return err
}

func bFindOrCreateGroup(ctx context.Context, db *ent.Client, group metrics.Group) (*ent.Group, bool, error) {
	created := false

	groupInst, err := db.Group.
		Query().
		Where(entGroup.NameEQ(group.Name)).
		Only(ctx)
	if err == nil {
		return groupInst, created, err
	}

	groupInst, err = db.Group.
		Create().
		SetName(group.Name).
		Save(ctx)
	created = true

	return groupInst, created, err
}

// func bFindOrCreateGraph(ctx context.Context, db *ent.Client, graph metrics.Graph, group *ent.Group) (*ent.Graph, error) {
// 	graphInst, err := db.Graph.
// 		Query().
// 		Where(entGraph.TitleEQ(graph.Title)).
// 		Only(ctx)
// 	if err == nil {
// 		return graphInst, err
// 	}

// 	graphInst, err = db.Graph.
// 		Create().
// 		SetTitle(graph.Title).
// 		SetUnit(graph.Unit).
// 		SetGroup(group).
// 		Save(ctx)

// 	return graphInst, err
// }

func bCreateGraph(ctx context.Context, db *ent.Client, graph metrics.Graph, group *ent.Group) (*ent.Graph, error) {
	return db.Graph.
		Create().
		SetTitle(graph.Title).
		SetUnit(graph.Unit).
		SetGroup(group).
		Save(ctx)
}

func bFindOrCreateMetric(ctx context.Context, db *ent.Client, metric metrics.Metric, graph *ent.Graph) (*ent.Metric, error) {
	metricInst, err := db.Metric.
		Query().
		Where(entMetric.TitleEQ(metric.Title)).
		Only(ctx)
	if err == nil {
		return metricInst, err
	}

	metricInst, err = db.Metric.
		Create().
		SetTitle(metric.Title).
		SetType(string(metric.Type)).
		SetGraph(graph).
		Save(ctx)

	return metricInst, err
}
