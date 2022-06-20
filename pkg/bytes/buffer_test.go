package bytes

import "testing"

func TestBuffer(t *testing.T) {
	t.Run("NewBuffer_EmptyBytes", func(t *testing.T) {
		buffer := NewBuffer(nil)
		if buffer == nil {
			t.Error("NewBuffer return nil")
		}
	})

	t.Run("NewBuffer_WithBytes", func(t *testing.T) {
		buffer := NewBuffer([]byte("test"))
		if buffer == nil {
			t.Error("NewBuffer return nil")
		}
		if buffer.String() != "test" {
			t.Errorf("expected: %s, actual: %s", "test", buffer.String())
		}
	})

	t.Run("NewBuffer_EmptyString", func(t *testing.T) {
		buffer := NewBufferString("")
		if buffer == nil {
			t.Error("NewBuffer return nil")
		}
	})

	t.Run("NewBuffer_WithString", func(t *testing.T) {
		buffer := NewBufferString("test")
		if buffer == nil {
			t.Error("NewBuffer return nil")
		}

		if buffer.String() != "test" {
			t.Errorf("expected: %s, actual: %s", "test", buffer.String())
		}
	})

	t.Run("String", func(t *testing.T) {
		buffer := NewBuffer([]byte(" \t\f\rtest\n\n"))
		if buffer == nil {
			t.Error("NewBuffer return nil")
		}
		if buffer.String() != "test" {
			t.Errorf("expected: %s, actual: %s", "test", buffer.String())
		}
	})

	t.Run("Bytes", func(t *testing.T) {
		buffer := NewBuffer([]byte(" \t\f\rtest\n\n"))
		if buffer == nil {
			t.Error("NewBuffer return nil")
		}
		if string(buffer.Bytes()) != "test" {
			t.Errorf("expected: %s, actual: %s", "test", buffer.Bytes())
		}
	})
}
