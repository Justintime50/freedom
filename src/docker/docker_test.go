package docker

import (
	"os"
	"os/exec"
	"testing"
)

const (
	mockStdout = "Total reclaimed space: 0B\n"
)

func TestPruneSuccess(t *testing.T) {
	stdout, err := Prune(fakeExecCommandSuccess)
	if err != nil {
		t.Error(err)
		return
	}

	// TODO: Assert the process was called with the right commands/args

	// Check to make sure the stdout is returned properly
	// Note: value matching is not checked since the command is not run
	stdoutStr := stdout.String()
	if stdoutStr != mockStdout {
		t.Errorf("stdout mismatch:\n%s\n vs \n%s", stdoutStr, mockStdout)
	}
}

func TestPruneFail(t *testing.T) {
	_, err := Prune(fakeExecCommandFailure)
	if err == nil {
		t.Errorf("Expected error due to shell command exiting with non-zero exit code")
	}
}

// fakeExecCommandSuccess is a function that initialises a new exec.Cmd, one which will
// simply call TestPruneProcessSuccess rather than the command it is provided. It will
// also pass through the command and its arguments as an argument to TestPruneProcessSuccess
func fakeExecCommandSuccess(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestPruneProcessSuccess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func fakeExecCommandFailure(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestPruneProcessFail", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestPruneProcessSuccess is a method that is called as a substitute for a shell command,
// the GO_TEST_PROCESS flag ensures that if it is called as part of the test suite, it is
// skipped.
func TestPruneProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Stdout.WriteString(mockStdout)
	os.Exit(0)
}

func TestPruneProcessFail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}
