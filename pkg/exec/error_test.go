package exec

import (
	"errors"
	"strings"
	"testing"
)

func TestExitCodeError_Code(t *testing.T) {
	t.Run("Code", func(t *testing.T) {
		err := ExitCodeError{code: 1}
		if err.Code() != 1 {
			t.Errorf("expected 1, got %d", err.Code())
		}
	})

	t.Run("Error", func(t *testing.T) {
		err := ExitCodeError{code: 1, err: errors.New("test")}
		if !strings.Contains(err.Error(), "test") {
			t.Errorf("expected error to contain test, got %s", err.Error())
		}
	})

	t.Run("Error_errNil", func(t *testing.T) {
		err := ExitCodeError{code: 1}
		if err.Error() != "exitCode: 1" {
			t.Errorf("expected error to be exitCode: 1, got %s", err.Error())
		}
	})
}
