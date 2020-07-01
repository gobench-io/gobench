package server

import (
	"time"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
	"golang.org/x/net/context"
)

// schedule get a pending application from the db if there is no active job
func (s *Server) schedule() {
	for {
		if !s.isSchedule {
			break
		}

		time.Sleep(1 * time.Second)

		if s.curr != nil {
			continue
		}

		app, err := s.nextApplication()

		if err != nil {
			continue
		}

		// save to state
		s.curr = app

		// compile the scenario
		s.compile(s.curr.Scenario)
	}
}

// provision compiles a scenario to golang plugin, distribute the application to
// worker. Return success when the workers confirm that the plugin is ready
func (s *Server) provision() (*ent.Application, error) {
	// compile
	return nil, nil
}

func (s *Server) nextApplication() (*ent.Application, error) {
	ctx, _ := context.WithCancel(context.TODO())

	app, err := s.master.db.
		Application.
		Query().
		Where(
			application.Status(string(appPending)),
		).
		Order(
			ent.Asc(application.FieldCreatedAt),
		).First(ctx)

	return app, err
}

// compile using go to compile a scenario in plugin build mode
// the result is path to so file
func (s *Server) compile(scen string) (string, error) {
	return "", nil
}
