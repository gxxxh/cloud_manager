package openstack

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
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
		packageAnalyzer := NewPackageAnalyzer(pkg)
		if packageAnalyzer.ASTFile != nil {
			packageAnalyzer.DoAnalyze()
		}
	}
	return err
}

type PackageAnalyzer struct {
	Pkg     *packages.Package
	ASTFile *ast.File
}

func NewPackageAnalyzer(pkg *packages.Package) *PackageAnalyzer {
	pa := &PackageAnalyzer{
		Pkg:     pkg,
		ASTFile: nil,
	}
	//get the ast for file requests.go
	for idx, file := range pkg.CompiledGoFiles {
		if strings.HasSuffix(file, "requests.go") {
			pa.ASTFile = pkg.Syntax[idx]
		}
	}
	return pa
}

// analyze packages and parse info
// todo 1. 判断是否是基本类型
// todo 2. 基于拼接生成参数和返回值类型
// todo 3. 生成代码
func (pa *PackageAnalyzer) DoAnalyze() {
	for _, d := range pa.ASTFile.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if pa.checkValidFunction(fn) {
				fmt.Println("handle function :", fn.Name)
				for _, paraExpr := range fn.Type.Params.List {
					paraName, paraTypeName := pa.parseField(paraExpr)
					fmt.Println(paraName, paraTypeName)
					typeName, packagePath := pa.parseExprTypeInfo(paraExpr.Type, pa.Pkg.TypesInfo)
					fmt.Println(typeName, packagePath)
				}
				fmt.Println("-----------------return------------------")
				for _, returnExpr := range fn.Type.Results.List {
					returnName, returnTypeName := pa.parseField(returnExpr)
					fmt.Println(returnName, returnTypeName)
					typeName, packagePath := pa.parseExprTypeInfo(returnExpr.Type, pa.Pkg.TypesInfo)
					fmt.Println(typeName, packagePath)
				}
			}
		}
	}
}

/*
check if the function is required(List, Create, Delete, Get...)
1. the first parameter should be * gophercloud.ServiceClient
*/
func (pa *PackageAnalyzer) checkValidFunction(fn *ast.FuncDecl) bool {
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

/*
type Name, type package path
*/
func (pa *PackageAnalyzer) parseTypeInfo(ty types.Type) (string, string) {
	switch tyType := ty.(type) {
	case *types.Pointer:
		typeName, packagePath := pa.parseTypeInfo(tyType.Elem())
		return "*" + typeName, packagePath
	case *types.Named:
		tmp := tyType.Obj()
		typeName := tmp.Pkg().Name() + "." + tmp.Name()
		return typeName, tmp.Pkg().Path()
	case *types.Basic:
		return tyType.Name(), ""
	default:
		fmt.Println("error! unhandled type: ", tyType)
		return "", ""
	}
}

func (pa *PackageAnalyzer) parseExprTypeInfo(expr ast.Expr, tinfo *types.Info) (string, string) {
	ty := tinfo.Types[expr].Type
	return pa.parseTypeInfo(ty)
}

/*
get the parameter/return name and type
*/
func (pa *PackageAnalyzer) parseExprName(expr ast.Expr) string {
	switch tyExpr := (expr).(type) {
	case *ast.StarExpr:
		return "*" + pa.parseExprName(tyExpr.X)
	case *ast.Ident:
		return tyExpr.Name
	case *ast.SelectorExpr:
		return pa.parseExprName(tyExpr.X) + "." + pa.parseExprName(tyExpr.Sel)
	default:
		return ""
	}
}

/*
parse field's name and type
*/
func (pa *PackageAnalyzer) parseField(field *ast.Field) (string, string) {
	paraName := ""
	if field.Names != nil {
		paraName = field.Names[0].Name
	}
	paraTypeName := pa.parseExprName(field.Type)
	return paraName, paraTypeName
}

func (ma *ModuleAnalyzer) AnalyzeFunction(fn *ast.FuncDecl) {
	fmt.Println(fn)
}
