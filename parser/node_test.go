package parser

import (
	"testing"
	"text/template/parse"
)

func Test_getFuncFromNode(t *testing.T) {
	t.Run("NIL", func(t *testing.T) {
		var branchNode *parse.BranchNode
		var listNode *parse.ListNode
		var commandNode *parse.CommandNode
		getFuncFromNode(branchNode, nil)
		getFuncFromNode(commandNode, nil)
		getFuncFromNode(listNode, nil)
		getFuncFromActionNode(nil, nil)
		getFuncFromListNode(nil, nil)
		getFuncFromListNode(nil, nil)
		getFuncFromBranchNode(nil, nil)
		getFuncFromCommandNode(nil, nil)
		getFuncFromIdentifierNode(nil, nil)
		getFuncFromIfNode(nil, nil)
		getFuncFromRangeNode(nil, nil)
		getFuncFromTemplateNode(nil, nil)
		getFuncFromWithNode(nil, nil)
	})
}
