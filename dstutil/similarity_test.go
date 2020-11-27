package dstutil

import (
	"fmt"
	"github.com/JfL0unch/dst"
	"github.com/JfL0unch/dst/decorator"
	"go/parser"
	"go/token"
	"testing"
)

func TestCursor_Similarity(t *testing.T) {

	input := `
package service

import (
	"context"

	"svcGenerator/data/proto/v1"
)

type CommonSvcService interface {
	IdentifyFetch(ctx context.Context, reqproto *commonproto.IdentifyFetchReqProto) (*commonproto.IdentifyFetchRespProto, error)

	//{{template9}}
}
var cnt int32
`
	pre := func(c *Cursor) bool {

		typeSpec := &dst.TypeSpec{
			Name: &dst.Ident{Name: "CommonSvcService"},
			Assign:  false,
			Type:  &dst.InterfaceType{},
		}
		xSpecs := make([]dst.Spec, 0)
		xSpecs = append(xSpecs, typeSpec)
		x := &dst.GenDecl{
			Tok:   token.TYPE,
			Specs: xSpecs,
		}

		names := make([]*dst.Ident,0)
		names = append(names,&dst.Ident{
			Name:"cnt",
		})
		valSpec := &dst.ValueSpec{
			Names: names,
			Type: &dst.Ident{Name: "int32"},
		}
		ySpecs := make([]dst.Spec, 0)
		ySpecs = append(ySpecs, valSpec)
		y := &dst.GenDecl{
			Tok: token.VAR,
			Specs: ySpecs,
		}

		if sim, hit := c.Similarity(x);sim >0{
			if sim != 4 {
				t.Errorf("got %d,expect %d",sim,4)
			}

			if hit != 4 {
				t.Errorf("got %d,expect %d",hit,3)
			}

		}

		if sim, hit := c.Similarity(y);sim>0{
			if sim != 3 {
				t.Errorf("got %d,expect %d",sim,3)
			}

			if hit != 3 {
				t.Errorf("got %d,expect %d",hit,3)
			}
		}

		return true
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", input, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	dstFile, err := decorator.DecorateFile(fset, f)
	if err != nil {
		t.Fatal(err)
	}

	dstFile = Apply(dstFile, pre, nil).(*dst.File)

}



func TestCursor_Similarity_Case1(t *testing.T) {

	input := `
package service

import (
	"context"

	"svcGenerator/data/proto/v1"
)

type CommonSvcService interface {
	IdentifyFetch(ctx context.Context, reqproto *commonproto.IdentifyFetchReqProto) (*commonproto.IdentifyFetchRespProto, error)

	//{{template9}}
}
var cnt int32
`
	pre := func(c *Cursor) bool {

		names :=make([]*dst.Ident,0)
		names = append(names,&dst.Ident{
			Name: "IdentifyFetch",
		})


		params := make([]*dst.Field,0)
		params = append(params,&dst.Field{
			Names: []*dst.Ident{NewIdent("ctx")},
			Type: &dst.SelectorExpr{
				X: &dst.Ident{Name: "context" },
				Sel: &dst.Ident{Name: "Context"},
			},
		})
		params = append(params,&dst.Field{
			Names: []*dst.Ident{NewIdent("ctx")},
			Type: &dst.StarExpr{
				X: &dst.SelectorExpr{
					X: &dst.Ident{Name: "commonproto"},
					Sel: &dst.Ident{Name: "IdentifyFetchReqProto"},
				},
			},
		})
		fieldListParams := &dst.FieldList{
			List: params,
		}

		field := &dst.Field{
			Names:  names,
			Type: &dst.FuncType{
				Func: false,
				Params: fieldListParams,
			},
			Tag: nil,
		}

		if sim, hit := c.Similarity(field);sim !=0{
			fmt.Printf("got %d,expect %d",sim,hit)

		}

		return true
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", input, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	dstFile, err := decorator.DecorateFile(fset, f)
	if err != nil {
		t.Fatal(err)
	}

	dstFile = Apply(dstFile, pre, nil).(*dst.File)

}
