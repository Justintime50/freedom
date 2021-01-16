package macos

import (
	"os"
	"os/exec"
	"testing"
)

const (
	testStdoutValue = ""
)

func TestFreeFinderWindowsSuccess(t *testing.T) {
	stdout, err := FreeFinderWindows(fakeExecCommandSuccess)
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

func TestFreeFinderWindowsFail(t *testing.T) {
	_, err := FreeFinderWindows(fakeExecCommandFailure)
	if err == nil {
		t.Errorf("Expected error due to shell command exiting with non-zero exit code")
	}
}

// fakeExecCommandSuccess is a function that initialises a new exec.Cmd, one which will
// simply call TestFreeFinderWindowsProcessSuccess rather than the command it is provided. It will
// also pass through the command and its arguments as an argument to TestFreeFinderWindowsProcessSuccess
func fakeExecCommandSuccess(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestFreeFinderWindowsProcessSuccess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func fakeExecCommandFailure(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestFreeFinderWindowsProcessFail", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

// TestFreeFinderWindowsProcessSuccess is a method that is called as a substitute for a shell command,
// the GO_TEST_PROCESS flag ensures that if it is called as part of the test suite, it is
// skipped.
func TestFreeFinderWindowsProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Stdout.WriteString(testStdoutValue)
	os.Exit(0)
}

func TestFreeFinderWindowsProcessFail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}
