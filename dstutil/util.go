package dstutil

import (
	"github.com/JfL0unch/dst"
	"go/token"
)

// Unparen returns e with any enclosing parentheses stripped.
func Unparen(e dst.Expr) dst.Expr {
	for {
		p, ok := e.(*dst.ParenExpr)
		if !ok {
			return e
		}
		e = p.X
	}
}


type ExprType string

const (
	ExprTypeIdent        ExprType = "Ident"
	ExprTypeStartExpr             = "StartExpr"
	ExprTypeSelectorExpr          = "SelectorExpr"
)

func NewIdent(name string) *dst.Ident {
	return &dst.Ident{
		Name:    name,
		Obj:     nil,
	}
}

/*
type svc interface {
  UserGet()
}
*/
func NewFuncType(params, results *dst.FieldList) *dst.FuncType {
	return &dst.FuncType{
		Results: results,
		Params:  params,
		Func:    false,
	}
}

//new valueSpec
func NewValueSpec(ident string, expr dst.Expr) *dst.ValueSpec {
	return &dst.ValueSpec{
		Names: []*dst.Ident{NewIdent(ident)},
		Type:  expr,
		//Type: NewIdent(typeName),
		Values:  nil,
	}
}

func NewCallExpr(f dst.Expr, args []dst.Expr) *dst.CallExpr {
	return &dst.CallExpr{
		Fun:      f,
		Args:     args,
		Ellipsis: false,
	}
}

func NewSelectExp(x dst.Expr, sel *dst.Ident) *dst.SelectorExpr {
	return &dst.SelectorExpr{
		X:   x,
		Sel: sel,
	}
}

func NewStarExp(exp dst.Expr) *dst.StarExpr {
	return &dst.StarExpr{
		X: exp,
		//X:NewIdent(ident),
	}
}

func NewKeyValueExp(key, val dst.Expr) *dst.KeyValueExpr {
	return &dst.KeyValueExpr{
		Key:   key,
		Value: val,
	}
}

/*
example ====>
type svc interface {
  UserGet(ctx context.Context, reqproto *partnerproto.StaffAuthFetchReqProto)(*partnerproto.StaffAuthFetchRespProto, error)
}

====>
typeVals: context.Context
typeName: StartExpr
fieldNames: ctx
tag: nil
*/
func NewField(fieldNames []string, typeVal dst.Expr, typeName ExprType, tag *dst.BasicLit) *dst.Field {
	names := make([]*dst.Ident, 0)
	for _, v := range fieldNames {
		names = append(names, NewIdent(v))
	}
	ret := &dst.Field{
		Names: names,
		Tag:   tag,
	}
	switch typeName {
	case ExprTypeStartExpr:
		ret.Type = typeVal
	case ExprTypeIdent:
		ret.Type = typeVal
	case ExprTypeSelectorExpr:
		//ret.Type = NewSelectExp(NewIdent(typeVals[0]), NewIdent(typeVals[1]))
		ret.Type = typeVal
	//todo:    other expr
	default:
		ret.Type = typeVal
	}

	return ret

}

func NewFieldOfFuncType(fieldNames []string, functype *dst.FuncType, tag *dst.BasicLit) *dst.Field {
	names := make([]*dst.Ident, 0)
	for _, v := range fieldNames {
		names = append(names, NewIdent(v))
	}
	ret := &dst.Field{
		Names: names,
		Type:  functype,
		Tag:   tag,
	}

	return ret
}
func NewFieldList(fields ...*dst.Field) *dst.FieldList {
	fieldList := make([]*dst.Field, 0)
	for _, v := range fields {
		fieldList = append(fieldList, v)
	}
	return &dst.FieldList{
		List: fieldList,
	}
}

//new funcDecl
func NewFuncDecl(funcName string, blkStmt *dst.BlockStmt, recv, params *dst.FieldList) *dst.FuncDecl {
	return &dst.FuncDecl{
		Name: &dst.Ident{Name: funcName},
		Body: blkStmt,
		Type: &dst.FuncType{Func: false,
			Params:  params,
			Results: nil,
		},
		Recv: recv,
	}
}

//new GenDecl
func NewGenDecl(tok token.Token, specs ...dst.Spec) *dst.GenDecl {
	return &dst.GenDecl{
		Tok:    tok,
		Specs:  specs,
	}
}

//new basicLit
func NewBasicLit(kind token.Token, value string) *dst.BasicLit {
	return &dst.BasicLit{
		Kind:     kind,
		Value:    value,
	}
}

//new compositeLit
func NewCompositeLit(typ dst.Expr, elts []dst.Expr) *dst.CompositeLit {
	return &dst.CompositeLit{
		Type:   typ,
		Elts:   elts,
	}
}

func NewAssignStmt(lhs, rhs []dst.Expr) *dst.AssignStmt {
	return &dst.AssignStmt{
		Lhs:    lhs,
		Rhs:    rhs,
		Tok:    token.ASSIGN,
	}
}

func NewShortAssignStmt(lhs, rhs []dst.Expr) *dst.AssignStmt {
	return &dst.AssignStmt{
		Lhs:    lhs,
		Rhs:    rhs,
		Tok:    token.DEFINE,
	}
}

/*
实例:
package main
func demo(){
	x := 3
}

========>
package main
func demo(){
	x := 3
	var guessCreateEndpoint kitendpoint.Endpoint
}
*/

func NewDeclStmt(decl dst.Decl) *dst.DeclStmt {
	return &dst.DeclStmt{
		Decl: decl,
	}
}

func NewReturnStmt(results []dst.Expr) *dst.ReturnStmt {
	return &dst.ReturnStmt{
		Results: results,
	}
}

func NewExpStmt(x dst.Expr) *dst.ExprStmt {
	return &dst.ExprStmt{
		X: x,
	}
}
func NewBlockStmt(expStmt ...*dst.ExprStmt) *dst.BlockStmt {
	ret := &dst.BlockStmt{}
	ret.List = make([]dst.Stmt, 0)
	for _, v := range expStmt {
		if expStmt == nil {
			continue
		}
		ret.List = append(ret.List, v)
	}
	return ret
}

func NewEmptyStmt(pos token.Pos) *dst.EmptyStmt {
	ret := &dst.EmptyStmt{}
	ret.Implicit = false

	return ret
}

//new import spec
func NewImportSpec(name string, path string) *dst.ImportSpec {
	return &dst.ImportSpec{
		Path: NewBasicLit(token.STRING, path),
		Name: NewIdent(name),
	}
}

func NewInterface(fieldList *dst.FieldList) *dst.InterfaceType {
	return &dst.InterfaceType{
		Methods:    fieldList,
		Incomplete: false,
	}
}
