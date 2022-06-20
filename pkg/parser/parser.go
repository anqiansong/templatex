package parser

import "text/template/parse"

// treeName is a dummy name for the root tree.
const treeName = "_tree"

// Parse returns the function identifiers for the given template definition.
func Parse(text string) ([]string, error) {
	tree := parse.New(treeName)
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
