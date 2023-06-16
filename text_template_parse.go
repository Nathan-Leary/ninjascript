package ninjascript
import (
text_template_parse "text/template/parse"
)
func init() {if _, ok := Api["text/template/parse"]; !ok {
   Api["text/template/parse"] = map[string]interface{}{}
}
Api["text/template/parse"]["IsEmptyTree"] = text_template_parse.IsEmptyTree
Api["text/template/parse"]["Parse"] = text_template_parse.Parse
Api["text/template/parse"]["ActionNode"] = text_template_parse.ActionNode{}
Api["text/template/parse"]["BoolNode"] = text_template_parse.BoolNode{}
Api["text/template/parse"]["BranchNode"] = text_template_parse.BranchNode{}
Api["text/template/parse"]["BreakNode"] = text_template_parse.BreakNode{}
Api["text/template/parse"]["ChainNode"] = text_template_parse.ChainNode{}
Api["text/template/parse"]["CommandNode"] = text_template_parse.CommandNode{}
Api["text/template/parse"]["CommentNode"] = text_template_parse.CommentNode{}
Api["text/template/parse"]["ContinueNode"] = text_template_parse.ContinueNode{}
Api["text/template/parse"]["DotNode"] = text_template_parse.DotNode{}
Api["text/template/parse"]["FieldNode"] = text_template_parse.FieldNode{}
Api["text/template/parse"]["IdentifierNode"] = text_template_parse.IdentifierNode{}
Api["text/template/parse"]["NewIdentifier"] = text_template_parse.NewIdentifier
Api["text/template/parse"]["IfNode"] = text_template_parse.IfNode{}
Api["text/template/parse"]["ListNode"] = text_template_parse.ListNode{}
Api["text/template/parse"]["ParseComments"] = text_template_parse.ParseComments
Api["text/template/parse"]["SkipFuncCheck"] = text_template_parse.SkipFuncCheck
Api["text/template/parse"]["NilNode"] = text_template_parse.NilNode{}
Api["text/template/parse"]["NodeText"] = text_template_parse.NodeText
Api["text/template/parse"]["NodeAction"] = text_template_parse.NodeAction
Api["text/template/parse"]["NodeBool"] = text_template_parse.NodeBool
Api["text/template/parse"]["NodeChain"] = text_template_parse.NodeChain
Api["text/template/parse"]["NodeCommand"] = text_template_parse.NodeCommand
Api["text/template/parse"]["NodeDot"] = text_template_parse.NodeDot
Api["text/template/parse"]["NodeField"] = text_template_parse.NodeField
Api["text/template/parse"]["NodeIdentifier"] = text_template_parse.NodeIdentifier
Api["text/template/parse"]["NodeIf"] = text_template_parse.NodeIf
Api["text/template/parse"]["NodeList"] = text_template_parse.NodeList
Api["text/template/parse"]["NodeNil"] = text_template_parse.NodeNil
Api["text/template/parse"]["NodeNumber"] = text_template_parse.NodeNumber
Api["text/template/parse"]["NodePipe"] = text_template_parse.NodePipe
Api["text/template/parse"]["NodeRange"] = text_template_parse.NodeRange
Api["text/template/parse"]["NodeString"] = text_template_parse.NodeString
Api["text/template/parse"]["NodeTemplate"] = text_template_parse.NodeTemplate
Api["text/template/parse"]["NodeVariable"] = text_template_parse.NodeVariable
Api["text/template/parse"]["NodeWith"] = text_template_parse.NodeWith
Api["text/template/parse"]["NodeComment"] = text_template_parse.NodeComment
Api["text/template/parse"]["NodeBreak"] = text_template_parse.NodeBreak
Api["text/template/parse"]["NodeContinue"] = text_template_parse.NodeContinue
Api["text/template/parse"]["NumberNode"] = text_template_parse.NumberNode{}
Api["text/template/parse"]["PipeNode"] = text_template_parse.PipeNode{}
Api["text/template/parse"]["RangeNode"] = text_template_parse.RangeNode{}
Api["text/template/parse"]["StringNode"] = text_template_parse.StringNode{}
Api["text/template/parse"]["TemplateNode"] = text_template_parse.TemplateNode{}
Api["text/template/parse"]["TextNode"] = text_template_parse.TextNode{}
Api["text/template/parse"]["Tree"] = text_template_parse.Tree{}
Api["text/template/parse"]["New"] = text_template_parse.New
Api["text/template/parse"]["VariableNode"] = text_template_parse.VariableNode{}
Api["text/template/parse"]["WithNode"] = text_template_parse.WithNode{}
Api["text/template/parse"]["IsEmptyTree"] = text_template_parse.IsEmptyTree
Api["text/template/parse"]["Parse"] = text_template_parse.Parse
Api["text/template/parse"]["NewIdentifier"] = text_template_parse.NewIdentifier
Api["text/template/parse"]["New"] = text_template_parse.New

}