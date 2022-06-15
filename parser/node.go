package parser

import "text/template/parse"

func getFuncFromNode(node parse.Node, fn *map[string]struct{}) {
	switch n := node.(type) {
	case *parse.ActionNode:
		getFuncFromActionNode(n, fn)
	case *parse.BranchNode:
		getFuncFromBranchNode(n, fn)
	case *parse.CommandNode:
		getFuncFromCommandNode(n, fn)
	case *parse.IdentifierNode:
		getFuncFromIdentifierNode(n, fn)
	case *parse.IfNode:
		getFuncFromIfNode(n, fn)
	case *parse.ListNode:
		getFuncFromListNode(n, fn)
	case *parse.PipeNode:
		getFuncFromPipeNode(n, fn)
	case *parse.RangeNode:
		getFuncFromRangeNode(n, fn)
	case *parse.TemplateNode:
		getFuncFromTemplateNode(n, fn)
	case *parse.WithNode:
		getFuncFromWithNode(n, fn)
	default:
		return
	}
}

func getFuncFromActionNode(node *parse.ActionNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	getFuncFromPipeNode(node.Pipe, fn)
}

func getFuncFromBranchNode(node *parse.BranchNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	getFuncFromPipeNode(node.Pipe, fn)
	getFuncFromListNode(node.List, fn)
	getFuncFromListNode(node.ElseList, fn)
}

func getFuncFromCommandNode(node *parse.CommandNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	for _, arg := range node.Args {
		getFuncFromNode(arg, fn)
	}
}

func getFuncFromIdentifierNode(node *parse.IdentifierNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	m := *fn
	m[node.Ident] = struct{}{}
	*fn = m
}

func getFuncFromIfNode(node *parse.IfNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	getFuncFromBranchNode(&node.BranchNode, fn)
}

func getFuncFromListNode(node *parse.ListNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	for _, n := range node.Nodes {
		getFuncFromNode(n, fn)
	}
}

func getFuncFromPipeNode(node *parse.PipeNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	for _, cmd := range node.Cmds {
		getFuncFromCommandNode(cmd, fn)
	}
}

func getFuncFromRangeNode(node *parse.RangeNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	getFuncFromBranchNode(&node.BranchNode, fn)
}

func getFuncFromTemplateNode(node *parse.TemplateNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	getFuncFromPipeNode(node.Pipe, fn)
}

func getFuncFromWithNode(node *parse.WithNode, fn *map[string]struct{}) {
	if node == nil {
		return
	}

	getFuncFromBranchNode(&node.BranchNode, fn)
}
