package parser

import "testing"

func TestParse(t *testing.T) {
	t.Run("funcIterator", func(t *testing.T) {
		funcIterator(nil)
	})

	t.Run("state_success", func(t *testing.T) {
		testData := []struct {
			data string
			want string
		}{
			{data: `{{"x"}}`},
			{data: `{{.}}`},
			{data: `{{.x}}`},
			{data: `{{.x .y}}`},
			{data: `{{fn .x}}`, want: "fn"},
			{data: `{{if .x}} {{end}}`},
			{data: `{{if not .x}} {{end}}`, want: "not"},
			{data: `{{not .x}}`, want: "not"},
			{data: `{{eq .x}}`, want: "eq"},
			{data: `{{"x" | hi }}`, want: "hi"},
			{data: `{{(say "hallo")}}`, want: "say"},
			{data: `{{range .x}}{{end}}`},
			{data: `{{with "hello"}} {{hello}} {{end}}`, want: "hello"},
			{data: `{{define "T1"}}ONE{{end}}
				{{define "T2"}}TWO{{end}}
				{{template "T1"}}
				{{template "T2"}}
				`,
			},
		}
		for _, d := range testData {
			actual, err := Parse(d.data)
			if err != nil {
				t.Errorf("Parse(%q) returned error %v", d.data, err)
			}

			if len(d.want) > 0 {
				if len(actual) != 1 {
					t.Errorf("Parse(%q) expected 1 node; got %d", d.data, len(actual))
				}

				if d.want != actual[0] {
					t.Errorf("Parse(%q) expected %s; got %s", d.data, d.want, actual[0])
				}
			} else {
				if len(actual) > 0 {
					t.Errorf("Parse(%q) expected no node; got %d", d.data, len(actual))
				}
			}
		}
	})
	t.Run("state_failure", func(t *testing.T) {
		testData := []string{"{{}}", "{{if .x}}"}
		for _, d := range testData {
			_, err := Parse(d)
			if err == nil {
				t.Errorf("Parse(%q) expected error; got none", d)
			}
		}
	})
}
