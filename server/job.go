package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
