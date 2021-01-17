package macos

import (
	"github.com/justintime50/mockcmd/mockcmd"
	"os"
	"testing"
)

func TestFreeFinderWindowsSuccess(t *testing.T) {
	stdout, err := FreeFinderWindows(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestFreeFinderWindowsFail(t *testing.T) {
	_, err := FreeFinderWindows(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

func TestMockProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Stdout.WriteString("mocked Stdout")
	os.Exit(0)
}

func TestMockProcessFailure(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}
