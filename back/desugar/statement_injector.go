package desugar

import "github.com/falcinspire/scriptblock/ast"

type StatementRewriter struct {
	Body  []ast.Statement
	index int
}

func NewStatementRewriter(body []ast.Statement, index int) *StatementRewriter {
	return &StatementRewriter{body, index}
}

func RewriteStatement(newStatement ast.Statement, rewriter *StatementRewriter) {
	rewriter.Body[rewriter.index] = newStatement
}
