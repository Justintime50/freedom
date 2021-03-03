package ports

import (
	"os"
	"testing"

	"github.com/justintime50/mockcmd/mockcmd"
)

func TestKillSuccess(t *testing.T) {
	stdout, err := Kill(mockcmd.MockExecSuccess, 5000)
	mockcmd.Success(t, stdout, err)
}

func TestKillFail(t *testing.T) {
	_, err := Kill(mockcmd.MockExecFailure, 5000)
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
