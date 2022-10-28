package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"strings"
)

// using to analyze openstack module staticly
// info was saved to openstack_resource_info
type ModuleAnalyzer struct {
	*packages.Config
}

func NewModuleAnalyzer() *ModuleAnalyzer {
	ma := &ModuleAnalyzer{
		&packages.Config{
			Mode: packages.LoadSyntax |
				packages.LoadTypes |
				packages.LoadFiles |
				packages.LoadAllSyntax |
				packages.LoadImports,
			Context:    nil,
			Dir:        "",
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
func (ma *ModuleAnalyzer) DoAnalyze(dir string) ([]*OpenstackResourceInfo, error) {
	resourceInfos := make([]*OpenstackResourceInfo, 0)
	ma.Config.Dir = dir
	pkgs, err := packages.Load(ma.Config, "./...")
	if err != nil {
		log.Println(err)
		return resourceInfos, err
	}
	packageAnalyzer := NewPackageAnalyzer()
	for _, pkg := range pkgs {
		if resourceInfo := packageAnalyzer.DoAnalyze(pkg); resourceInfo != nil {
			resourceInfos = append(resourceInfos, resourceInfo)
		}
	}
	return resourceInfos, err
}

type PackageAnalyzer struct {
}

func NewPackageAnalyzer() *PackageAnalyzer {
	pa := &PackageAnalyzer{}
	return pa
}

// get the ast for file requests.go
func (pa *PackageAnalyzer) GetASTFile(pkg *packages.Package) *ast.File {
	for idx, file := range pkg.CompiledGoFiles {
		if strings.HasSuffix(file, "requests.go") {
			return pkg.Syntax[idx]
		}
	}
	return nil
}

// analyze packages and parse info
func (pa *PackageAnalyzer) DoAnalyze(pkg *packages.Package) *OpenstackResourceInfo {
	log.Printf("*********************************analyze %s*****************************", pkg.Name)
	resourceInfo := NewOpenstackResourceInfo(pkg.Name, pkg.PkgPath)
	astFile := pa.GetASTFile(pkg)
	if astFile == nil {
		return nil
	}
	for _, d := range astFile.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if pa.checkValidFunction(fn) {
				actionInfo := NewOpenstackActionInfo(fn.Name.String())
				log.Println("******************handle function***************** :", fn.Name)
				if fn.Name.String() == "ListAddressesByNetwork" {
					fmt.Println(fn.Name)
				}
				parseFieldList := func(fieldList []*ast.Field, kind string) {
					log.Printf("-----------------%v:%d------------------/n", kind, len(fieldList))
					for _, expr := range fieldList {
						//a field may contain two name with the same type
						names, _ := pa.parseFieldInfo(expr)
						typeName, packagePath := pa.parseExprTypeInfo(expr.Type, pkg.TypesInfo)
						if typeName == "*gophercloud.ServiceClient" {
							continue
						}
						for _, name := range names {
							actionInfo.AddVarInfo(name, typeName, kind)
							log.Println(name, typeName, packagePath)
						}
						if len(names) == 0 { //return has no names
							actionInfo.AddVarInfo("", typeName, kind)
						}
						if packagePath != "" {
							resourceInfo.ImportPaths.Insert(packagePath)
						}
					}
				}
				parseFieldList(fn.Type.Params.List, "parameters")
				parseFieldList(fn.Type.Results.List, "returns")
				resourceInfo.AddAction(actionInfo)
			}
		}
	}
	return resourceInfo
}

/*
check if the function is required(List, Create, Delete, Get...)
1. the first parameter should be * gophercloud.ServiceClient
*/
func (pa *PackageAnalyzer) checkValidFunction(fn *ast.FuncDecl) bool {
	funcName := fn.Name.String()
	//check if the function is exported
	if funcName[0] >= 'a' && funcName[0] <= 'z' {
		return false
	}
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
		if tmp.Pkg() != nil {
			return tmp.Pkg().Name() + "." + tmp.Name(), tmp.Pkg().Path()
		} else {
			return tmp.Name(), ""
		}
	case *types.Basic:
		return tyType.Name(), ""
	case *types.Slice:
		typeName, packagePath := pa.parseTypeInfo(tyType.Elem())
		return "[]" + typeName, packagePath
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
func (pa *PackageAnalyzer) parseFieldInfo(field *ast.Field) ([]string, []string) {
	names := make([]string, len(field.Names), len(field.Names))
	typeNames := make([]string, len(field.Names), len(field.Names))
	for idx, fieldName := range field.Names {
		names[idx] = fieldName.Name
		typeNames[idx] = pa.parseExprName(field.Type)
	}
	return names, typeNames
}
