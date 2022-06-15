package parser

import "text/template/parse"

func Parse(text string) ([]string, error) {
	tree := parse.New("test")
	tree.Mode = parse.SkipFuncCheck
	tree, err := tree.Parse(text, "", "", map[string]*parse.Tree{})
	if err != nil {
		return nil, err
	}

	var fns []string
	fn := funcIterator(tree)
	for f := range fn {
		fns = append(fns, f)
	}

	return fns, nil
}

func funcIterator(tree *parse.Tree) map[string]struct{} {
	if tree == nil {
		return nil
	}
	m := make(map[string]struct{})
	getFuncFromListNode(tree.Root, &m)
	return m
}
