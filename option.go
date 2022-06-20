package templatex

type Option func(*TemplateX)

func WithBinary(binary string) Option {
	return WithRender(NewBinaryRender(binary))
}

func WithRender(render FuncRender) Option {
	return func(t *TemplateX) {
		t.render = render
	}
}
