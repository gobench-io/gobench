package option

import "testing"

func TestConfigureOptions(t *testing.T) {
	type testPrint struct {
		args          []string
		version, help func()
	}

	version := func() {}
	help := func() {}

	testFuncs := []testPrint{
		{[]string{
			"--agent-sock", "agent/sock",
			"--exectutor-sock", "executor/soc",
			"--driver-path", "driver/path",
		}, version, help}
	}
}
