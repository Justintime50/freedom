package ports

import (
	"os"
	"os/exec"
	"testing"
)

const (
	testStdoutValue = ""
)

func TestKillSuccess(t *testing.T) {
	stdout, err := Kill(fakeExecCommandSuccess, 5000)
	if err != nil {
		t.Error(err)
		return
	}

	// TODO: Assert the process was called with the right commands/args

	stdoutStr := stdout.String()
	if stdoutStr != testStdoutValue {
		t.Errorf("stdout mismatch:\n%s\n vs \n%s", stdoutStr, testStdoutValue)
	}
}

func TestKillFail(t *testing.T) {
	_, err := Kill(fakeExecCommandFailure, 5000)
	if err == nil {
		t.Errorf("Expected error due to shell command exiting with non-zero exit code")
	}
}

// fakeExecCommandSuccess is a function that initialises a new exec.Cmd, one which will
// simply call TestKillProcessSuccess rather than the command it is provided. It will
// also pass through the command and its arguments as an argument to TestKillProcessSuccess
func fakeExecCommandSuccess(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestKillProcessSuccess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func fakeExecCommandFailure(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestKillProcessFail", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestKillProcessSuccess is a method that is called as a substitute for a shell command,
// the GO_TEST_PROCESS flag ensures that if it is called as part of the test suite, it is
// skipped.
func TestKillProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Stdout.WriteString(testStdoutValue)
	os.Exit(0)
}

func TestKillProcessFail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}
