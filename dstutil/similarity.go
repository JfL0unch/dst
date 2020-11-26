package dstutil

import (
	"github.com/JfL0unch/dst"
	"go/token"
)

// similarity return similarity between node and n.
func similarity(node dst.Node,targetNode dst.Node) (int,int) {
	switch x:= node.(type){
	case *dst.GenDecl:
		if n,ok := targetNode.(*dst.GenDecl);ok{
			sim,hit := similarityGenDecl(x,n)
			return sim,hit
		}else{
			return 0,0
		}

	}

	return 0,0
}

func similarityGenDecl(n1,n2 *dst.GenDecl) (int,int){
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

func similarityToken(n1,n2 token.Token) (int,int){
	if n1 == n2 {
		return 1,1
	}else{
		return 0,1
	}
}


func similaritySpec(n1,n2 dst.Spec) (int,int){
	switch x1:= n1.(type){
	case *dst.TypeSpec:
		if x2,ok:= n2.(*dst.TypeSpec);ok{
			return similarityTypeSpec(*x1,*x2)
		}else{
			return 0,1
		}

	case *dst.ValueSpec:
		if x2,ok:= n2.(*dst.ValueSpec);ok{
			return similarityValueSpec(*x1,*x2)
		}else{
			return 0,1
		}

	case *dst.ImportSpec:
		if x2,ok:= n2.(*dst.ImportSpec);ok{
			return similarityImportSpec(*x1,*x2)
		}else{
			return 0,1
		}

	}

	return 0,0
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

func similarityExpr(n1,n2 dst.Expr)(int,int){
	return similarity(n1,n2)
}