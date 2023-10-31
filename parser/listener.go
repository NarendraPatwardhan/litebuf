// Code generated from LitebufParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // LitebufParser

import "github.com/antlr4-go/antlr/v4"

// LitebufListener is a complete listener for a parse tree produced by LitebufParser.
type LitebufListener interface {
	antlr.ParseTreeListener

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterDef is called when entering the def production.
	EnterDef(c *DefContext)

	// EnterStruct is called when entering the struct production.
	EnterStruct(c *StructContext)

	// EnterUnion is called when entering the union production.
	EnterUnion(c *UnionContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterEnum is called when entering the enum production.
	EnterEnum(c *EnumContext)

	// EnterVariant is called when entering the variant production.
	EnterVariant(c *VariantContext)

	// EnterOrder is called when entering the order production.
	EnterOrder(c *OrderContext)

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterFn is called when entering the fn production.
	EnterFn(c *FnContext)

	// EnterChannel is called when entering the channel production.
	EnterChannel(c *ChannelContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterPrimitive is called when entering the primitive production.
	EnterPrimitive(c *PrimitiveContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitDef is called when exiting the def production.
	ExitDef(c *DefContext)

	// ExitStruct is called when exiting the struct production.
	ExitStruct(c *StructContext)

	// ExitUnion is called when exiting the union production.
	ExitUnion(c *UnionContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitEnum is called when exiting the enum production.
	ExitEnum(c *EnumContext)

	// ExitVariant is called when exiting the variant production.
	ExitVariant(c *VariantContext)

	// ExitOrder is called when exiting the order production.
	ExitOrder(c *OrderContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitFn is called when exiting the fn production.
	ExitFn(c *FnContext)

	// ExitChannel is called when exiting the channel production.
	ExitChannel(c *ChannelContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitPrimitive is called when exiting the primitive production.
	ExitPrimitive(c *PrimitiveContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}

// BaseLitebufListener is a complete listener for a parse tree produced by LitebufParser.
type BaseLitebufListener struct{}

var _ LitebufListener = &BaseLitebufListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLitebufListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLitebufListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLitebufListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLitebufListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSource is called when production source is entered.
func (s *BaseLitebufListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production source is exited.
func (s *BaseLitebufListener) ExitSource(ctx *SourceContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseLitebufListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseLitebufListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterInclude is called when production include is entered.
func (s *BaseLitebufListener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *BaseLitebufListener) ExitInclude(ctx *IncludeContext) {}

// EnterDef is called when production def is entered.
func (s *BaseLitebufListener) EnterDef(ctx *DefContext) {}

// ExitDef is called when production def is exited.
func (s *BaseLitebufListener) ExitDef(ctx *DefContext) {}

// EnterStruct is called when production struct is entered.
func (s *BaseLitebufListener) EnterStruct(ctx *StructContext) {}

// ExitStruct is called when production struct is exited.
func (s *BaseLitebufListener) ExitStruct(ctx *StructContext) {}

// EnterUnion is called when production union is entered.
func (s *BaseLitebufListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production union is exited.
func (s *BaseLitebufListener) ExitUnion(ctx *UnionContext) {}

// EnterField is called when production field is entered.
func (s *BaseLitebufListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLitebufListener) ExitField(ctx *FieldContext) {}

// EnterEnum is called when production enum is entered.
func (s *BaseLitebufListener) EnterEnum(ctx *EnumContext) {}

// ExitEnum is called when production enum is exited.
func (s *BaseLitebufListener) ExitEnum(ctx *EnumContext) {}

// EnterVariant is called when production variant is entered.
func (s *BaseLitebufListener) EnterVariant(ctx *VariantContext) {}

// ExitVariant is called when production variant is exited.
func (s *BaseLitebufListener) ExitVariant(ctx *VariantContext) {}

// EnterOrder is called when production order is entered.
func (s *BaseLitebufListener) EnterOrder(ctx *OrderContext) {}

// ExitOrder is called when production order is exited.
func (s *BaseLitebufListener) ExitOrder(ctx *OrderContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseLitebufListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseLitebufListener) ExitRoot(ctx *RootContext) {}

// EnterService is called when production service is entered.
func (s *BaseLitebufListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseLitebufListener) ExitService(ctx *ServiceContext) {}

// EnterFn is called when production fn is entered.
func (s *BaseLitebufListener) EnterFn(ctx *FnContext) {}

// ExitFn is called when production fn is exited.
func (s *BaseLitebufListener) ExitFn(ctx *FnContext) {}

// EnterChannel is called when production channel is entered.
func (s *BaseLitebufListener) EnterChannel(ctx *ChannelContext) {}

// ExitChannel is called when production channel is exited.
func (s *BaseLitebufListener) ExitChannel(ctx *ChannelContext) {}

// EnterType is called when production type is entered.
func (s *BaseLitebufListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseLitebufListener) ExitType(ctx *TypeContext) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BaseLitebufListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BaseLitebufListener) ExitPrimitive(ctx *PrimitiveContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseLitebufListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseLitebufListener) ExitIdentifier(ctx *IdentifierContext) {}
