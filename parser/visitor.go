// Code generated from LitebufParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // LitebufParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LitebufParser.
type LitebufVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LitebufParser#source.
	VisitSource(ctx *SourceContext) interface{}

	// Visit a parse tree produced by LitebufParser#namespace.
	VisitNamespace(ctx *NamespaceContext) interface{}

	// Visit a parse tree produced by LitebufParser#include.
	VisitInclude(ctx *IncludeContext) interface{}

	// Visit a parse tree produced by LitebufParser#def.
	VisitDef(ctx *DefContext) interface{}

	// Visit a parse tree produced by LitebufParser#struct.
	VisitStruct(ctx *StructContext) interface{}

	// Visit a parse tree produced by LitebufParser#union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by LitebufParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by LitebufParser#enum.
	VisitEnum(ctx *EnumContext) interface{}

	// Visit a parse tree produced by LitebufParser#variant.
	VisitVariant(ctx *VariantContext) interface{}

	// Visit a parse tree produced by LitebufParser#order.
	VisitOrder(ctx *OrderContext) interface{}

	// Visit a parse tree produced by LitebufParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by LitebufParser#service.
	VisitService(ctx *ServiceContext) interface{}

	// Visit a parse tree produced by LitebufParser#fn.
	VisitFn(ctx *FnContext) interface{}

	// Visit a parse tree produced by LitebufParser#channel.
	VisitChannel(ctx *ChannelContext) interface{}

	// Visit a parse tree produced by LitebufParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by LitebufParser#primitive.
	VisitPrimitive(ctx *PrimitiveContext) interface{}

	// Visit a parse tree produced by LitebufParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}

type BaseLitebufVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLitebufVisitor) VisitSource(ctx *SourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitNamespace(ctx *NamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitInclude(ctx *IncludeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitDef(ctx *DefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitUnion(ctx *UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitEnum(ctx *EnumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitVariant(ctx *VariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitOrder(ctx *OrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitService(ctx *ServiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitFn(ctx *FnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitChannel(ctx *ChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitPrimitive(ctx *PrimitiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLitebufVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
