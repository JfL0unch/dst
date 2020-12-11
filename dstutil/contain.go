package dstutil

import (
	"fmt"
	"github.com/JfL0unch/dst"
	"go/token"
	"reflect"
)

// 对于targetNode 的每一个属性 ，是否node都有对等的属性
// 有任意一个属性不相等，则返回true
func contain(node dst.Node,targetNode dst.Node) bool {
	switch x:= targetNode.(type){
	case *dst.File:
		// do nothing
	case *dst.Package:
		// do nothing
	case *dst.GenDecl:
		if n,ok := node.(*dst.GenDecl);ok{
			return containGenDecl(n,x)
		}else{
			return false
		}
	case *dst.InterfaceType:
		if n,ok := node.(*dst.InterfaceType);ok{
			contain := containInterfaceType(n,x)
			return contain
		}else{
			return false
		}

	case *dst.Field:
		if n,ok := node.(*dst.Field);ok{
			contain := containField(n,x)
			return contain
		}else{
			return false
		}
	case *dst.FieldList:
		if n,ok := node.(*dst.FieldList);ok{
			contain := containFieldList(n,x)
			return contain
		}else{
			return false
		}
	case *dst.StructType:
		if n,ok := node.(*dst.StructType);ok{
			contain := containStructType(n,x)
			return contain
		}else{
			return false
		}
	case *dst.Ident:
		if n,ok := node.(*dst.Ident);ok{
			contain := containIdent(n,x)
			return contain
		}else{
			return false
		}
	case *dst.BasicLit:
		if n,ok := node.(*dst.BasicLit);ok{
			contain := containBasicLit(n,x)
			return contain
		}else{
			return false
		}
	case *dst.SelectorExpr:
		if n,ok := node.(*dst.SelectorExpr);ok{
			contain := containSelectorExpr(n,x)
			return contain
		}else{
			return false
		}
	case *dst.UnaryExpr:
		if n,ok := node.(*dst.UnaryExpr);ok{
			contain := containUnaryExpr(n,x)
			return contain
		}else{
			return false
		}
	case *dst.ImportSpec:
		if n,ok := node.(*dst.ImportSpec);ok{
			contain := containImportSpec(n,x)
			return contain
		}else{
			return false
		}
	case *dst.TypeSpec:
		if n,ok := node.(*dst.TypeSpec);ok{
			contain := containTypeSpec(n,x)
			return contain
		}else{
			return false
		}
	case *dst.StarExpr:
		if n,ok := node.(*dst.StarExpr);ok{
			contain := containStarExpr(n,x)
			return contain
		}else{
			return false
		}

	case *dst.ValueSpec:
		if n,ok := node.(*dst.ValueSpec);ok{
			contain := containValueSpec(n,x)
			return contain
		}else{
			return false
		}
	case *dst.FuncDecl:
		if n,ok := node.(*dst.FuncDecl);ok{
			contain := containFuncDecl(n,x)
			return contain
		}else{
			return false
		}
	case *dst.FuncType:
		if n,ok := node.(*dst.FuncType);ok{
			contain := containFuncType(n,x)
			return contain
		}else{
			return false
		}
	case *dst.CompositeLit:
		if n,ok := node.(*dst.CompositeLit);ok{
			contain := containCompositeLit(n,x)
			return contain
		}else{
			return false
		}
	case *dst.ReturnStmt:
		if n,ok := node.(*dst.ReturnStmt);ok{
			contain := containReturnStmt(n,x)
			return contain
		}else{
			return false
		}
	case *dst.BlockStmt:
		if n,ok := node.(*dst.BlockStmt);ok{
			contain := containBlockStmt(n,x)
			return contain
		}else{
			return false
		}
	case *dst.AssignStmt:
		if n,ok := node.(*dst.AssignStmt);ok{
			contain := containAssignStmt(n,x)
			return contain
		}else{
			return false
		}
	case *dst.DeclStmt:
		if n,ok := node.(*dst.DeclStmt);ok{
			contain := containDeclStmt(n,x)
			return contain
		}else{
			return false
		}
	default:
		if node != nil{
			v := reflect.ValueOf(node)
			panic(fmt.Errorf("undefined dst.Node %s",v.Type()))
		}

	}

	return false
}

func containGenDecl(n1,n2 *dst.GenDecl) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	if !containInt(int(n1.Tok),int(n2.Tok)){
		return false
	}

	if len(n1.Specs) < len(n2.Specs){
		return false
	}

	if len(n2.Specs)>0{
		for k,v := range n2.Specs{
			if !contain(n1.Specs[k],v){
				return false
			}
		}
	}

	return true
}


func containFuncDecl(n1,n2 *dst.FuncDecl) bool{

	if n1 != nil && n2 ==nil{
		return false
	}

	if n1 == nil && n2 !=nil{
		return false
	}

	if n1 == nil && n2 ==nil {
		return true
	}
	
	if !containIdent(n1.Name,n2.Name){
		return false
	}

	if !containFuncType(n1.Type,n2.Type){
		return false
	}

	if !containBlockStmt(n1.Body,n2.Body){
		return false
	}

	if !containFieldList(n1.Recv,n2.Recv){
		return false
	}


	return true
}


func containFuncType(n1,n2 *dst.FuncType) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}
	
	if !containFieldList(n1.Params,n2.Params){
		return false
	}

	if !containFieldList(n1.Results,n2.Results){
		return false
	}

	return true
}


func containBlockStmt(n1,n2 *dst.BlockStmt) bool{

	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}
	

	if len(n1.List) < len(n2.List){
		return false
	}

	for k,v := range n2.List{
		if !contain(n1.List[k],v){
			return false
		}
	}

	return true
}

func containToken(n1,n2 token.Token) bool{
	if n1 != n2 {
		return false
	}else{
		return true
	}
}


func containDeclStmt(n1,n2 *dst.DeclStmt) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	return contain(n1.Decl,n2.Decl)
}


func containAssignStmt(n1,n2 *dst.AssignStmt) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	if len(n1.Lhs) < len(n2.Lhs){
		return false
	}

	for k,v := range n2.Lhs{
		if !contain(n1.Lhs[k],v){
			return false
		}
	}

	return true
}

func containTypeSpec(n1,n2 *dst.TypeSpec) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	if !containIdent(n1.Name,n2.Name){
		return false
	}

	if !n1.Assign != n2.Assign{
		return false
	}

	if !contain(n1.Type,n2.Type){
		return false
	}

	return true
}

func containValueSpec(n1,n2 *dst.ValueSpec) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	if !contain(n1.Type,n2.Type){
		return false
	}

	if len(n1.Names) < len(n2.Names){
		return false
	}

	for k,v := range n2.Names{
		if !containIdent(n1.Names[k],v){
			return false
		}
	}

	if len(n1.Values) < len(n2.Values){
		return false
	}

	for k,v := range n2.Values{
		if !contain(n1.Values[k],v){
			return false
		}
	}

	return true
}

func containImportSpec(n1,n2 *dst.ImportSpec) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	if !containIdent(n1.Name,n2.Name){
		return false
	}

	if !containBasicLit(n1.Path,n2.Path){
		return false
	}

	return true
}

func containSelectorExpr(n1,n2 *dst.SelectorExpr) bool{
	if n1 != nil && n2 ==nil{
		return false
	}

	if n1 == nil && n2 !=nil{
		return false
	}

	if n1 == nil && n2 ==nil {
		return true
	}
	
	if !containIdent(n1.Sel,n2.Sel){
		return false
	}

	if !contain(n1.X,n2.X){
		return false
	}

	return true
}

func containStarExpr(n1,n2 *dst.StarExpr)bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	if !contain(n1.X,n2.X){
		return false
	}

	return true
}

func containIdent(n1,n2 *dst.Ident) bool{
	if n1 != nil && n2 ==nil{
		return false
	}

	if n1 == nil && n2 !=nil{
		return false
	}

	if n1 == nil && n2 ==nil {
		return true
	}

	if n1 != nil && n2 !=nil {
		if !containString(n1.Name,n2.Name){
			return false
		}
	}

	if !containString(n1.Path,n2.Path){
		return false
	}

	if !containObject(n1.Obj,n2.Obj){
		return false
	}
	return true
}

func containCompositeLit(n1,n2 *dst.CompositeLit) bool{
	if n1 != nil && n2 ==nil{
		return false
	}

	if n1 == nil && n2 !=nil{
		return false
	}

	if n1 == nil && n2 ==nil {
		return true
	}
	
	if !contain(n1.Type,n2.Type){
		return false
	}

	if len(n1.Elts) < len(n2.Elts){
		return false
	}

	for k,v := range n2.Elts {
		if !contain(n1.Elts[k],v){
			return false
		}
	}
	
	return true
}


func containReturnStmt(n1,n2 *dst.ReturnStmt) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	if len(n1.Results) < len(n2.Results){
		return false
	}
	
	for k,v := range n2.Results{
		if !contain(n1.Results[k],v){
			return false
		}
	}

	return true
}


func containUnaryExpr(n1,n2 *dst.UnaryExpr) bool{
	if n1 !=nil && n2 == nil{
		return false
	}

	if n1 ==nil && n2 == nil{
		return true
	}

	if n1 ==nil && n2 != nil{
		return false
	}

	if !contain(n1.X,n2.X){
		return false
	}

	if !containToken(n1.Op,n2.Op){
		return false
	}

	return true
}

func containBasicLit(n1,n2 *dst.BasicLit) bool{
	if n1 == nil && n2==nil{
		return true
	}

	if n1 != nil && n2==nil{
		return false
	}


	if n1 == nil && n2!=nil{
		return false
	}
	if n1!=nil && n2!=nil{
		if !containInt(int(n1.Kind),int(n2.Kind)){
			return false
		}

		if !containString(n1.Value,n2.Value){
			return false
		}

	}
	
	return true
}

func containString(n1,n2 string)bool{
	if n1==n2 {
		return true
	}else{
		return false
	}
}

func containInt(n1,n2 int)bool{
	if n1==n2 {
		return true
	}else{
		return false
	}
}


func containObject(n1,n2 *dst.Object)bool{
	if n1 == nil && n2 !=nil{
		return false
	}
	if n1 != nil && n2 ==nil{
		return false
	}
	if n1 == nil && n2 ==nil{
		return true
	}
	
	if n1 !=nil && n2 !=nil {
		if !containString(n1.Name,n2.Name){
			return false
		}

		if containInt(int(n1.Kind),int(n2.Kind)){
			return false
		}
	}

	return true
}

func containInterfaceType(n1,n2 *dst.InterfaceType)bool{

	if n1.Methods == nil && n2.Methods !=nil{
		return false
	}
	if n1.Methods != nil && n2.Methods ==nil{
		return false
	}
	if n1.Methods == nil && n2.Methods ==nil{
		return true
	}
	
	if n1.Methods != nil && n2.Methods!=nil{
		if !containFieldList(n1.Methods,n2.Methods){
			return false
		}
	}
	return true
}

func containFieldList(n1,n2 *dst.FieldList)bool{

	if n1 == nil && n2 !=nil{
		return false
	}
	if n1 != nil && n2 ==nil{
		return false
	}
	if n1 == nil && n2 ==nil{
		return true
	}

	if n1 != nil && n2 !=nil{
		if len(n1.List) < len(n2.List){
			return false
		}
		
		for k,v := range n2.List{
			if !containField(n1.List[k],v){
				return false
			}
		}
	}
	
	return true
}


func containStructType(n1,n2 *dst.StructType)bool{
	if n1 == nil && n2 !=nil{
		return false
	}
	if n1 != nil && n2 ==nil{
		return false
	}
	if n1 == nil && n2 ==nil{
		return true
	}

	if n1 != nil && n2 !=nil{
		if !containFieldList(n1.Fields,n2.Fields){
			return false
		}
	}

	return true
}


func containField(n1,n2 *dst.Field)bool{

	if n1 == nil && n2 !=nil{
		return false
	}
	if n1 != nil && n2 ==nil{
		return false
	}
	if n1 == nil && n2 ==nil{
		return true
	}

	if n1 != nil && n2 !=nil{
		if !containBasicLit(n1.Tag,n2.Tag){
			return false
		}

		if !contain(n1.Type, n2.Type){
			return false
		}

		if len(n1.Names) < len(n2.Names){
			return false
		}
		
		for k,v := range n2.Names{
			if !containIdent(n1.Names[k],v){
				return false
			}
		}
	}

	return true
}