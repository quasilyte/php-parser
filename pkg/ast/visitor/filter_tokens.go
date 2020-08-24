package visitor

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type FilterTokens struct {
	Null
}

func (v *FilterTokens) EnterNode(n ast.Vertex) bool {
	n.GetNode().Tokens = nil
	n.Accept(v)
	return true
}

func (v *FilterTokens) StmtUse(n *ast.StmtUse) {
	n.UseTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtGroupUse(n *ast.StmtGroupUse) {
	n.UseTkn = nil
	n.LeadingNsSeparatorTkn = nil
	n.NsSeparatorTkn = nil
	n.OpenCurlyBracketTkn = nil
	n.CloseCurlyBracketTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtUseDeclaration(n *ast.StmtUseDeclaration) {
	n.NsSeparatorTkn = nil
	n.AsTkn = nil
	n.CommaTkn = nil
}

func (v *FilterTokens) NameNamePart(n *ast.NameNamePart) {
	n.NsSeparatorTkn = nil
	n.StringTkn = nil
}

func (v *FilterTokens) NameName(n *ast.NameName) {
	n.ListSeparatorTkn = nil
}

func (v *FilterTokens) NameFullyQualified(n *ast.NameFullyQualified) {
	n.NsSeparatorTkn = nil
	n.ListSeparatorTkn = nil
}

func (v *FilterTokens) NameRelative(n *ast.NameRelative) {
	n.NsTkn = nil
	n.NsSeparatorTkn = nil
	n.ListSeparatorTkn = nil
}

func (v *FilterTokens) StmtNamespace(n *ast.StmtNamespace) {
	n.NsTkn = nil
	n.OpenCurlyBracket = nil
	n.CloseCurlyBracket = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtHaltCompiler(n *ast.StmtHaltCompiler) {
	n.HaltCompilerTkn = nil
	n.OpenParenthesisTkn = nil
	n.CloseParenthesisTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtConstList(n *ast.StmtConstList) {
	n.ConstTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtClassConstList(n *ast.StmtClassConstList) {
	n.ConstTkn = nil
	n.SemiColonTkn = nil
}

func (v *FilterTokens) StmtConstant(n *ast.StmtConstant) {
	n.EqualTkn = nil
	n.CommaTkn = nil
}