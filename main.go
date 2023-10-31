package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"machinelearning.one/litebuf/parser"
)

type TreeShapeListener struct {
	*parser.BaseLitebufListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (l *TreeShapeListener) EnterNamespace(ctx *parser.NamespaceContext) {
	// Get the text from the namespace rule's context
	idCtx := ctx.Identifier().IDENTIFIER()
	fmt.Println(idCtx.GetText())
}

func main() {
	input := antlr.NewInputStream("namespace foobar")

	lexer := parser.NewLitebufLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewLitebufParser(stream)
	tree := p.Source()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
