package exec

import (
	"encoding/json"
	"os"
	"os/exec"
)

// Execute executes a command with optional arguments.
func Execute(name string, option ...Option) error {
	cmd := exec.Command(name)
	builtinOption := append([]Option(nil), WithEnv(os.Environ()))
	builtinOption = append(builtinOption, option...)
	for _, opt := range builtinOption {
		opt(cmd)
	}

	if err := cmd.Run(); err != nil {
		exitCode := cmd.ProcessState.ExitCode()
		if exitCode != 0 {
			return ExitCodeError{code: exitCode, err: err}
		} else {
			return err
		}
	}

	return nil
}

// Marshal marshals a value to a string.
func Marshal(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Unmarshal unmarshals a string to a value.
func Unmarshal(s string, v any) error {
	return json.Unmarshal([]byte(s), v)
}
