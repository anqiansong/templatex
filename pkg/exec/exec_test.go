//go:build !windows

package exec

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/anqiansong/templatex/pkg/bytes"
)

const dummyCommand = "_dummy_"

func TestExecute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var output bytes.Buffer
		err := Execute("echo",
			WithArgs("hello"),
			WithStdout(&output),
		)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if output.String() != "hello" {
			t.Errorf("expected output 'hello', got %q", output.String())
		}
	})

	t.Run("ExitCodeError", func(t *testing.T) {
		err := Execute(dummyCommand)
		if err == nil {
			t.Errorf("expected error,actual nil")
		}
		if !errors.As(err, &ExitCodeError{}) {
			t.Errorf("expected ExitCodeError,actual %v", err)
		}
	})

	t.Run("WithEnv", func(t *testing.T) {
		err := Execute("echo",
			WithEnv([]string{"A=1"}),
			WithArgs("hello"),
			withTest(t, func(cmd *exec.Cmd, t *testing.T) {
				env := cmd.Env
				for _, e := range env {
					if e == "A=1" {
						return
					}
				}
				t.Error("expected A=1, but not exists")
			}),
		)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("WithArgs", func(t *testing.T) {
		var buf bytes.Buffer
		err := Execute("echo",
			WithArgs("hello"),
			WithStdout(&buf),
		)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if buf.String() != "hello" {
			t.Errorf("expected output 'hello', got %q", buf.String())
		}
	})

	t.Run("WithWorkDir", func(t *testing.T) {
		tmp := t.TempDir()
		err := Execute("echo",
			WithWorkDir(tmp),
			WithArgs("hello"),
			withTest(t, func(cmd *exec.Cmd, t *testing.T) {
				if cmd.Dir != tmp {
					t.Errorf("expected Dir %q, got %q", tmp, cmd.Dir)
				}
			}),
		)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("WithStdin", func(t *testing.T) {
		err := Execute("echo",
			WithStdin(bytes.NewBufferString("hello")),
			WithArgs("hello"),
			withTest(t, func(cmd *exec.Cmd, t *testing.T) {
				if cmd.Stdin == nil {
					t.Error("expected Stdin, but nil")
				}
			}),
		)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("WithStderr", func(t *testing.T) {
		var buf bytes.Buffer
		err := Execute("echo",
			WithStderr(&buf),
			withTest(t, func(cmd *exec.Cmd, t *testing.T) {
				if cmd.Stderr == nil {
					t.Errorf("expected Stderr, but nil")
				}
			}),
		)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Marshal_Unmarshal", func(t *testing.T) {
		var s = "test"
		data, err := Marshal(s)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		var receive string
		err = Unmarshal(data, &receive)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

	})
}
