package dstutil

import (
	"fmt"
	"github.com/JfL0unch/dst"
	"go/token"
	"reflect"
)

// similarity return similarity between node and n.
func similarity(node dst.Node,targetNode dst.Node) (int,int) {
	switch x:= node.(type){
	case *dst.File:
		// do nothing
	case *dst.Package:
		// do nothing
	case *dst.GenDecl:
		if n,ok := targetNode.(*dst.GenDecl);ok{
			sim,hit := similarityGenDecl(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.InterfaceType:
		if n,ok := targetNode.(*dst.InterfaceType);ok{
			sim,hit := similarityInterfaceType(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}

	case *dst.Field:
		if n,ok := targetNode.(*dst.Field);ok{
			sim,hit := similarityField(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.FieldList:
		if n,ok := targetNode.(*dst.FieldList);ok{
			sim,hit := similarityFieldList(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.StructType:
		if n,ok := targetNode.(*dst.StructType);ok{
			sim,hit := similarityStructType(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.Ident:
		if n,ok := targetNode.(*dst.Ident);ok{
			sim,hit := similarityIdent(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.BasicLit:
		if n,ok := targetNode.(*dst.BasicLit);ok{
			sim,hit := similarityBasicLit(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.SelectorExpr:
		if n,ok := targetNode.(*dst.SelectorExpr);ok{
			sim,hit := similaritySelectorExpr(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.UnaryExpr:
		if n,ok := targetNode.(*dst.UnaryExpr);ok{
			sim,hit := similarityUnaryExpr(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.ImportSpec:
		if n,ok := targetNode.(*dst.ImportSpec);ok{
			sim,hit := similarityImportSpec(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.TypeSpec:
		if n,ok := targetNode.(*dst.TypeSpec);ok{
			sim,hit := similarityTypeSpec(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.StarExpr:
		if n,ok := targetNode.(*dst.StarExpr);ok{
			sim,hit := similarityStarExpr(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}

	case *dst.ValueSpec:
		if n,ok := targetNode.(*dst.ValueSpec);ok{
			sim,hit := similarityValueSpec(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.FuncDecl:
		if n,ok := targetNode.(*dst.FuncDecl);ok{
			sim,hit := similarityFuncDecl(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.FuncType:
		if n,ok := targetNode.(*dst.FuncType);ok{
			sim,hit := similarityFuncType(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.CompositeLit:
		if n,ok := targetNode.(*dst.CompositeLit);ok{
			sim,hit := similarityCompositeLit(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.ReturnStmt:
		if n,ok := targetNode.(*dst.ReturnStmt);ok{
			sim,hit := similarityReturnStmt(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}
	case *dst.BlockStmt:
		if n,ok := targetNode.(*dst.BlockStmt);ok{
			sim,hit := similarityBlockStmt(*x,*n)
			return sim,hit
		}else{
			return 0,0
		}

	default:
		if node != nil{
			v := reflect.ValueOf(node)
			panic(fmt.Errorf("undefined dst.Node %s",v.Type()))
		}

	}

	return 0,0
}

func similarityGenDecl(n1,n2 dst.GenDecl) (int,int){
	sim,hit := similarityToken(n1.Tok, n2.Tok)

	for k,v := range n1.Specs{
		if k >= len(n2.Specs){
			break
		}
		s,h := similaritySpec(v,n2.Specs[k])
		sim,hit = sim+s,hit+h
	}

	return sim,hit
}


func similarityFuncDecl(n1,n2 dst.FuncDecl) (int,int){
	sim,hit :=0,0
	if n1.Name != nil && n2.Name != nil{
		s,h := similarityIdent(*n1.Name,*n2.Name)
		sim,hit = sim+s,hit+h
	}

	if n1.Type != nil && n2.Type != nil{
		s,h := similarityFuncType(*n1.Type,*n2.Type)
		sim,hit = sim+s,hit+h
	}

	if n1.Body != nil && n2.Body != nil{
		s,h := similarityBlockStmt(*n1.Body,*n2.Body)
		sim,hit = sim+s,hit+h
	}

	if n1.Recv != nil && n2.Recv != nil{
		for k,v := range n1.Recv.List{
			if k >= len(n2.Recv.List){
				break
			}
			s,h := similarityField(*v,*n2.Recv.List[k])
			sim,hit = sim+s,hit+h
		}
	}


	return sim,hit
}


func similarityFuncType(n1,n2 dst.FuncType) (int,int){
	sim,hit :=0,0
	if n1.Params != nil && n2.Params != nil{
		for k,v := range n1.Params.List{
			if k >= len(n2.Params.List){
				break
			}
			s,h := similarityField(*v,*n2.Params.List[k])
			sim,hit = sim+s,hit+h
		}
	}

	if n1.Results != nil && n2.Results != nil{
		for k,v := range n1.Results.List{
			if k >= len(n2.Results.List){
				break
			}
			s,h := similarityField(*v,*n2.Results.List[k])
			sim,hit = sim+s,hit+h
		}
	}

	return sim,hit
}


func similarityBlockStmt(n1,n2 dst.BlockStmt) (int,int){
	sim,hit :=0,0
	for k,v := range n1.List{
		if k >= len(n2.List){
			break
		}
		s,h := similarityStmt(v,n2.List[k])
		sim,hit = sim+s,hit+h
	}

	return sim,hit
}

func similarityToken(n1,n2 token.Token) (int,int){
	if n1 == n2 {
		return 1,1
	}else{
		return 0,1
	}
}


func similarityDeclStmt(n1,n2 dst.DeclStmt) (int,int){
	sim,hit := similarityDecl(n1.Decl,n2.Decl)

	return sim,hit
}


func similarityAssignStmt(n1,n2 dst.AssignStmt) (int,int){
	sim,hit := 0,0
	for k,v := range n1.Lhs{
		if k >= len(n2.Lhs){
			break
		}
		s,h := similarityExpr(v,n2.Lhs[k])
		sim,hit = sim+s,hit+h
	}

	return sim,hit
}


func similarityDecl(n1,n2 dst.Decl) (int,int) {
	switch x1 := n1.(type) {
	case *dst.GenDecl:
		if x2, ok := n2.(*dst.GenDecl); ok {
			return similarityGenDecl(*x1, *x2)
		} else {
			return 0, 1
		}

	default:
		if n1 != nil{
			v := reflect.ValueOf(n1)
			panic(fmt.Errorf("undefined Decl: %s",v.Type()))
		}

	}
	return 0,0
}

func similarityStmt(n1,n2 dst.Stmt) (int,int){
	return similarity(n1,n2)
	//switch x1:= n1.(type){
	//case *dst.DeclStmt:
	//	if x2,ok:= n2.(*dst.DeclStmt);ok{
	//		return similarityDeclStmt(*x1,*x2)
	//	}else{
	//		return 0,1
	//	}
	//
	//case *dst.AssignStmt:
	//	if x2,ok:= n2.(*dst.AssignStmt);ok{
	//		return similarityAssignStmt(*x1,*x2)
	//	}else{
	//		return 0,1
	//	}
	//
	//default:
	//	panic(errors.New("undefined stmt"))
	//
	//}
	//
	//return 0,0
}

func similaritySpec(n1,n2 dst.Spec) (int,int){
	//switch x1:= n1.(type){
	//case *dst.TypeSpec:
	//	if x2,ok:= n2.(*dst.TypeSpec);ok{
	//		return similarityTypeSpec(*x1,*x2)
	//	}else{
	//		return 0,1
	//	}
	//
	//case *dst.ValueSpec:
	//	if x2,ok:= n2.(*dst.ValueSpec);ok{
	//		return similarityValueSpec(*x1,*x2)
	//	}else{
	//		return 0,1
	//	}
	//
	//case *dst.ImportSpec:
	//	if x2,ok:= n2.(*dst.ImportSpec);ok{
	//		return similarityImportSpec(*x1,*x2)
	//	}else{
	//		return 0,1
	//	}
	//default:
	//	v := reflect.ValueOf(n1)
	//	panic(fmt.Errorf("undefined spec:%s",v.Type()))
	//
	//}
	//
	//return 0,0

	return similarity(n1,n2)
}


func similarityExpr(n1,n2 dst.Expr)(int,int){

	return similarity(n1,n2)
	//switch x1 := n1.(type){
	//case *dst.UnaryExpr:
	//	if x2,ok:= n2.(*dst.UnaryExpr);ok{
	//		return similarityUnaryExpr(*x1,*x2)
	//	}else{
	//		return 0,1
	//	}
	//default:
	//	v := reflect.ValueOf(n1)
	//	panic(fmt.Errorf("undefined expr:%s",v.Type()))
	//}
	//return similarity(n1,n2)
}


func similarityTypeSpec(n1,n2 dst.TypeSpec) (int,int){
	sim,hit :=0,0
	if n1.Name != nil && n2.Name != nil{
		s,h := similarityIdent(*n1.Name,*n2.Name)
		sim,hit = sim+s,hit+h
	}

	s,h := similarityBool(n1.Assign,n2.Assign)
	sim,hit = sim+s,hit+h

	s,h = similarityExpr(n1.Type,n2.Type)
	sim,hit = sim+s,hit+h

	return sim,hit
}

func similarityValueSpec(n1,n2 dst.ValueSpec) (int,int){
	sim,hit := similarityExpr(n1.Type,n2.Type)

	for k,v := range n1.Names{
		if k >= len(n2.Names){
			break
		}
		s,h := similarityIdent(*v,*n2.Names[k])
		sim,hit = sim+s,hit+h
	}

	for k,v := range n1.Values{
		if k >= len(n2.Values){
			break
		}
		s,h := similarityExpr(v,n2.Names[k])
		sim,hit = sim+s,hit+h
	}

	return sim,hit
}

func similarityImportSpec(n1,n2 dst.ImportSpec) (int,int){
	sim,hit := 0,0
	if n1.Name != nil && n2.Name != nil {
		s,h := similarityIdent(*n1.Name,*n2.Name)
		sim,hit = sim+s,hit+h
	}

	if n1.Path != nil && n2.Path != nil{
		s,h := similarityBasicLit(*n1.Path,*n2.Path)
		sim,hit = sim+s,hit+h
	}

	return sim,hit
}

func similaritySelectorExpr(n1,n2 dst.SelectorExpr) (int,int){
	sim,hit := 0,0

	if n1.Sel !=nil && n2.Sel != nil{
		s,h := similarityIdent(*n1.Sel,*n2.Sel)
		sim,hit = sim+s,hit+h
	}

	s,h := similarityExpr(n1.X,n2.X)
	sim,hit = sim+s,hit+h

	return sim,hit
}

func similarityStarExpr(n1,n2 dst.StarExpr)(int,int){
	sim,hit := 0,0

	s,h := similarityExpr(n1.X,n2.X)
	sim,hit = sim+s,hit+h

	return sim,hit
}

func similarityIdent(n1,n2 dst.Ident) (int,int){
	sim,hit := similarityString(n1.Name,n2.Name)

	s,h := similarityString(n1.Path,n2.Path)
	sim,hit = sim+s,hit+h

	if n1.Obj != nil && n2.Obj !=nil{
		s,h = similarityObject(*n1.Obj,*n2.Obj)
		sim,hit = sim+s,hit+h
	}

	return sim,hit
}


func similarityCompositeLit(n1,n2 dst.CompositeLit) (int,int){
	sim,hit := similarityExpr(n1.Type,n2.Type)

	for k,v := range n1.Elts{
		if k >= len(n2.Elts){
			break
		}
		s,h := similarityExpr(v,n2.Elts[k])
		sim,hit = sim+s,hit+h
	}

	return sim,hit
}


func similarityReturnStmt(n1,n2 dst.ReturnStmt) (int,int){
	sim,hit := 0,0
	for k,v := range n1.Results{
		if k >= len(n2.Results){
			break
		}
		s,h := similarityExpr(v,n2.Results[k])
		sim,hit = sim+s,hit+h
	}
	return sim,hit
}


func similarityUnaryExpr(n1,n2 dst.UnaryExpr) (int,int){
	sim,hit := similarityExpr(n1.X,n2.X)

	s,h := similarityToken(n1.Op,n2.Op)
	sim,hit = sim+s,hit+h

	return sim,hit
}

func similarityBasicLit(n1,n2 dst.BasicLit) (int,int){
	sim,hit := similarityInt(int(n1.Kind),int(n2.Kind))

	s,h := similarityString(n1.Value,n2.Value)
	sim,hit = sim+s,hit+h

	return sim,hit
}

func similarityString(n1,n2 string)(int,int){
	if n1==n2 {
		return 1,1
	}else{
		return 0,1
	}
}

func similarityInt(n1,n2 int)(int,int){
	if n1==n2 {
		return 1,1
	}else{
		return 0,1
	}
}

func similarityBool(n1,n2 bool)(int,int){
	if n1==n2 {
		return 1,1
	}else{
		return 0,1
	}
}

func similarityObject(n1,n2 dst.Object)(int,int){
	sim,hit := similarityString(n1.Name,n2.Name)

	s,h := similarityInt(int(n1.Kind),int(n2.Kind))
	sim,hit = sim+s,hit+h

	return sim,hit
}

func similarityInterfaceType(n1,n2 dst.InterfaceType)(int,int){
	if n1.Methods != nil && n2.Methods!=nil{
		sim,hit := similarityFieldList(*n1.Methods,*n2.Methods)
		return sim,hit
	}
	return 0,0
}

func similarityFieldList(n1,n2 dst.FieldList)(int,int){

	sim,hit := 0,0
	for k,v := range n1.List{
		if k >= len(n2.List){
			break
		}
		s,h := similarityField(*v,*n2.List[k])
		sim,hit = sim+s,hit+h
	}

	return sim,hit
}


func similarityStructType(n1,n2 dst.StructType)(int,int){
	sim,hit := 0,0
	if n1.Fields != nil && n2.Fields != nil{
		for k,field := range n1.Fields.List{
			if k >= len(n2.Fields.List){
				break
			}
			s,h := similarityField(*field, *n2.Fields.List[k])
			sim,hit = sim+s,hit+h
		}
	}


	return sim,hit
}


func similarityField(n1,n2 dst.Field)(int,int){

	sim,hit := 0,0

	if n1.Tag != nil && n2.Tag != nil{
		s,h := similarityBasicLit(*n1.Tag,*n2.Tag)
		sim,hit = sim+s,hit+h
	}

	if len(n1.Names)>0{
		for k,ident := range n1.Names{
			if k >= len(n2.Names){
				break
			}
			s,h := similarityIdent(*ident, *n2.Names[k])
			sim,hit = sim+s,hit+h
		}
	}

	s,h := similarityExpr(n1.Type, n2.Type)
	sim,hit = sim+s,hit+h

	return sim,hit
}