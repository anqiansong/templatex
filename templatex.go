package templatex

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"github.com/anqiansong/templatex/pkg/parser"
)

type TemplateX struct {
	*template.Template
	render FuncRender
}

func New(name string, options ...Option) *TemplateX {
	instance := &TemplateX{Template: template.New(name)}
	for _, option := range options {
		option(instance)
	}
	return instance
}

func (t *TemplateX) Parse(text string) (*TemplateX, error) {
	if _, err := t.funcs(text, ""); err != nil {
		return nil, err
	}

	tmp, err := t.Template.Parse(text)
	if err != nil {
		return nil, err
	}

	t.Template = tmp
	return t, nil
}

func (t *TemplateX) ParseFiles(filenames ...string) (*TemplateX, error) {
	if _, err := t.fileFuncs(filenames...); err != nil {
		return nil, err
	}

	tmp, err := t.Template.ParseFiles(filenames...)
	if err != nil {
		return nil, err
	}

	t.Template = tmp
	return t, nil
}

func (t *TemplateX) ParseFS(fsys fs.FS, patterns ...string) (*TemplateX, error) {
	var filenames []string
	for _, pattern := range patterns {
		list, err := fs.Glob(fsys, pattern)
		if err != nil {
			return nil, err
		}

		if len(list) == 0 {
			return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
		}

		filenames = append(filenames, list...)
	}

	if _, err := t.fileFuncs(filenames...); err != nil {
		return nil, err
	}

	tmp, err := t.Template.ParseFS(fsys, patterns...)
	if err != nil {
		return nil, err
	}

	t.Template = tmp
	return t, nil
}

func (t *TemplateX) ParseGlob(pattern string) (*TemplateX, error) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if _, err = t.fileFuncs(filenames...); err != nil {
		return nil, err
	}

	tmp, err := t.Template.ParseGlob(pattern)
	if err != nil {
		return nil, err
	}

	t.Template = tmp
	return t, nil
}

func (t *TemplateX) Funcs(funcMap template.FuncMap) *TemplateX {
	tmp := t.Template.Funcs(funcMap)
	t.Template = tmp
	return t
}

func (t *TemplateX) funcs(text string, filename string) (*TemplateX, error) {
	if t.render != nil {
		fns, err := parser.Parse(text)
		if err != nil {
			return nil, err
		}
		funcMap := make(template.FuncMap)
		for _, fn := range fns {
			if !t.render.IsSupportFn(fn) {
				continue
			}

			funcMap[fn] = func(in ...interface{}) (interface{}, error) {
				output, err := t.render.Render(fn, in...)
				if err != nil {
					if len(filename) > 0 {
						return "", fmt.Errorf("file:%s, err:%s", filename, err.Error())
					}
					return "", err
				}
				return output, nil
			}
		}

		t.Funcs(funcMap)
	}
	return t, nil
}

func (t *TemplateX) fileFuncs(filenames ...string) (*TemplateX, error) {
	for _, filename := range filenames {
		text, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		if _, err = t.funcs(string(text), filename); err != nil {
			return nil, err
		}
	}

	return t, nil
}
