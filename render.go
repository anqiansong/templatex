package templatex

type FuncRender interface {
	IsSupportFn(fn string) bool
	Render(fn string, arg ...interface{}) (interface{}, error)
}
