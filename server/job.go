package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"context"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
)

// to is the function to set new state for an application
// save new state to the db
func (jb *job) to(ctx context.Context, state jobState) error {
	return jb.app.Update().
		SetStatus(string(state)).
		Exec(ctx)
}

// schedule get a pending application from the db if there is no active job
func (s *Server) schedule() {
	for {
		if !s.isSchedule {
			break
		}

		time.Sleep(1 * time.Second)

		ctx, _ := context.WithCancel(context.Background())

		app, err := s.nextApplication(ctx)

		if err != nil {
			continue
		}

		s.job = &job{
			app: app,
		}

		// compile the scenario
		if s.job.plugin, err = s.compile(ctx, s.job.app.Scenario); err != nil {
			continue
		}
	}
}

// provision compiles a scenario to golang plugin, distribute the application to
// worker. Return success when the workers confirm that the plugin is ready
func (s *Server) provision() (*ent.Application, error) {
	// compile
	return nil, nil
}

func (s *Server) nextApplication(ctx context.Context) (*ent.Application, error) {
	app, err := s.master.db.
		Application.
		Query().
		Where(
			application.Status(string(jobPending)),
		).
		Order(
			ent.Asc(application.FieldCreatedAt),
		).First(ctx)

	return app, err
}

// compile using go to compile a scenario in plugin build mode
// the result is path to so file
func (s *Server) compile(ctx context.Context, scen string) (string, error) {
	var path string
	var err error

	// save the scenario to a tmp file
	tmpScenF, err := ioutil.TempFile("", "gobench-scenario-*.go")
	if err != nil {
		return path, fmt.Errorf("failed creating temp scenario file: %v", err)
	}
	tmpScenName := tmpScenF.Name()

	defer os.Remove(tmpScenName) // cleanup

	_, err = tmpScenF.Write([]byte(scen))
	if err != nil {
		return path, fmt.Errorf("failed write to scenario file: %v", err)
	}

	if err = tmpScenF.Close(); err != nil {
		return path, fmt.Errorf("failed close the scenario file: %v", err)
	}

	path = fmt.Sprintf("%s.out", tmpScenName)

	// compile the scenario to a tmp file
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o",
		path, tmpScenName)

	// if out, err := cmd.CombinedOutput(); err != nil {
	if err := cmd.Run(); err != nil {
		return path, fmt.Errorf("failed compiling the scenario: %v", err)
	}

	return path, err
}
