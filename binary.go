package templatex

import (
	"github.com/anqiansong/templatex/pkg/bytes"
	"github.com/anqiansong/templatex/pkg/exec"
)

type BinaryRender struct {
	binaryFile string
}

func NewBinaryRender(binaryFile string) *BinaryRender {
	return &BinaryRender{binaryFile: binaryFile}
}

func (b *BinaryRender) getAllArgs() []string {
	var (
		out  bytes.Buffer
		list []string
	)
	if err := exec.Execute(b.binaryFile, exec.WithArgs("templatex"), exec.WithStdout(&out)); err != nil {
		return nil
	}

	if err := exec.Unmarshal(out.String(), &list); err != nil {
		return nil
	}

	return list
}

func (b *BinaryRender) IsSupportFn(fn string) bool {
	supportFlags := b.getAllArgs()
	m := make(map[string]struct{})
	for _, item := range supportFlags {
		m[item] = struct{}{}
	}

	_, ok := m[fn]
	return ok
}

func (b *BinaryRender) Render(fn string, arg ...any) (any, error) {
	argStr, err := exec.Marshal(arg[:])
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	if err := exec.Execute(
		b.binaryFile,
		exec.WithArgs(fn, argStr),
		exec.WithStdout(&out)); err != nil {
		return "", err
	}

	return out.String(), nil
}
