package docker

import (
	"os"
	"testing"

	"github.com/justintime50/mockcmd/mockcmd"
)

func TestPruneSuccess(t *testing.T) {
	stdout, err := Prune(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestPruneFail(t *testing.T) {
	_, err := Prune(mockcmd.MockExecFailure)
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
