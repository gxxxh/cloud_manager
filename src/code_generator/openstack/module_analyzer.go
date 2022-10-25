package openstack

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"log"
	"strings"
)

type ModuleAnalyzer struct {
	*packages.Config
}

func NewModuleAnalyzer(dir string) *ModuleAnalyzer {
	ma := &ModuleAnalyzer{
		&packages.Config{
			Mode: packages.LoadSyntax |
				packages.LoadTypes |
				packages.LoadFiles |
				packages.LoadAllSyntax |
				packages.LoadImports,
			Context:    nil,
			Dir:        dir,
			Env:        nil,
			BuildFlags: nil,
			Fset:       token.NewFileSet(),
			ParseFile:  nil,
			Tests:      false,
		},
	}
	return ma
}

// parse all the packages in the module
func (ma *ModuleAnalyzer) DoAnalyze() error {
	pkgs, err := packages.Load(ma.Config, "./...")
	if err != nil {
		log.Println(err)
		return err
	}
	for _, pkg := range pkgs {
		requestAST := ma.GetRequestAST(pkg)
		if requestAST != nil {
			ma.AnalyzeAST(requestAST)
		}
	}
	return err
}

/*
get the ast for file requests.go
*/
func (ma *ModuleAnalyzer) GetRequestAST(pkg *packages.Package) *ast.File {
	for idx, file := range pkg.CompiledGoFiles {
		if strings.HasSuffix(file, "requests.go") {
			return pkg.Syntax[idx]
		}
	}
	return nil
}

/*
check if the function is required(List, Create, Delete, Get...)
1. the first parameter should be * gophercloud.ServiceClient
*/
func (ma *ModuleAnalyzer) FunctionParameterFilter(fn *ast.FuncDecl) bool {
	if fn.Recv == nil { //function's Recv filed is nil, method is not
		if len(fn.Type.Params.List) != 0 {
			//the first parameter should be a star expr(pointer)
			paraStarExpr, isStarExpr := (fn.Type.Params.List[0].Type).(*ast.StarExpr) //*gophercloud.ServiceClient
			if isStarExpr {
				//the start expr should contain a selector expr(A.B)
				paraSelectorExpr, isSelectorExpr := (paraStarExpr.X).(*ast.SelectorExpr)
				if isSelectorExpr {
					if paraSelectorExpr.Sel.Name == "ServiceClient" {
						return true
					}
				}
			}
		}
	}
	return false
}

// analyze packages and parse info
func (ma *ModuleAnalyzer) AnalyzeAST(f *ast.File) {
	for _, d := range f.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if ma.FunctionParameterFilter(fn) {
				for _, paraExpr := range fn.Type.Params.List {
					paraName, paraTypeName := ma.ParseParameters(paraExpr)
					fmt.Println(paraName, paraTypeName)
				}
			}
		}
	}
}

/*
get the parameter/return name and type
*/

func (ma *ModuleAnalyzer) ParseExprName(paraExpr ast.Expr) string {
	switch tyExpr := (paraExpr).(type) {
	case *ast.StarExpr:
		return "*" + ma.ParseExprName(tyExpr.X)
	case *ast.Ident:
		return tyExpr.Name
	case *ast.SelectorExpr:
		return ma.ParseExprName(tyExpr.X) + "." + ma.ParseExprName(tyExpr.Sel)
	default:
		return ""
	}
}

func (ma *ModuleAnalyzer) ParseParameters(paraExpr *ast.Field) (string, string) {
	paraName := paraExpr.Names[0].Name
	paraTypeName := ma.ParseExprName(paraExpr.Type)
	return paraName, paraTypeName
}

func (ma *ModuleAnalyzer) AnalyzeFunction(fn *ast.FuncDecl) {
	fmt.Println(fn)
}
