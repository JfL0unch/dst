package dstutil

import (
	"github.com/JfL0unch/dst"
	"github.com/JfL0unch/dst/decorator"
	"go/parser"
	"go/token"
	"testing"
)

func TestCursor_Contain_Case1(t *testing.T) {

	input := `
// x
package service


// y
var cnt int32
`
	pre := func(c *Cursor) bool {

		// 'type CommonSvcService interface{}'
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
		_=x


		// 'var cnt int32'
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

		//if !c.Contain(x){
		//	t.Errorf("got %s,expect %s","non-contained","contained")
		//
		//}

		if c.Contain(y){
			t.Errorf("got %s,%s",c.Node().Decorations(),"contained")
			return false
		}else{
			t.Errorf("got %s,%s",c.Node().Decorations(),"non-contained")
			return true
		}

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



func TestCursor_Contain_Case2(t *testing.T) {

	input := `
package service

import (
	"context"

	"svcGenerator/data/proto/v1"
)

type CommonSvcService interface {
	IdentifyFetch(ctx context.Context, reqproto *commonproto.IdentifyFetchReqProto) (*commonproto.IdentifyFetchRespProto, error)
}
`
	pre := func(c *Cursor) bool {

		// `IdentifyFetch(ctx context.Context, reqproto *commonproto.IdentifyFetchReqProto){}'
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
				Params: fieldListParams,
			},
		}

		if !c.Contain(field){
			t.Errorf("got %s,expect %s","non-contained","contained")
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
