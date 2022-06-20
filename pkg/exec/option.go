package exec

import (
	"io"
	"os/exec"
	"testing"
)

// Option is a function that can be used to modify the behavior of Execute.
type Option func(cmd *exec.Cmd)

// WithEnv sets the environment variables for the command.
func WithEnv(env []string) Option {
	return func(cmd *exec.Cmd) {
		cmd.Env = append(cmd.Env, env...)
	}
}

// WithArgs sets the arguments for the command.
func WithArgs(args ...string) Option {
	return func(cmd *exec.Cmd) {
		cmd.Args = append(cmd.Args, args...)
	}
}

// WithWorkDir sets the working directory for the command.
func WithWorkDir(workDir string) Option {
	return func(cmd *exec.Cmd) {
		cmd.Dir = workDir
	}
}

// WithStdin sets the stdin for the command.
func WithStdin(stdin io.Reader) Option {
	return func(cmd *exec.Cmd) {
		cmd.Stdin = stdin
	}
}

// WithStdout sets the stdout for the command.
func WithStdout(stdout io.Writer) Option {
	return func(cmd *exec.Cmd) {
		cmd.Stdout = stdout
	}
}

// WithStderr sets the stderr for the command.
func WithStderr(stderr io.Writer) Option {
	return func(cmd *exec.Cmd) {
		cmd.Stderr = stderr
	}
}

func withTest(t *testing.T, testFn func(cmd *exec.Cmd, t *testing.T)) Option {
	return func(cmd *exec.Cmd) {
		testFn(cmd, t)
	}
}
