package analyzer

import (
	"cloud_manager/src/utils"
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

	for _, pkg := range pkgs {
		packageAnalyzer := NewPackageAnalyzer(pkg)
		if resourceInfo := packageAnalyzer.DoAnalyze(); resourceInfo != nil {
			resourceInfos = append(resourceInfos, resourceInfo)
		}
	}
	return resourceInfos, err
}

type PackageAnalyzer struct {
	pkg *packages.Package
}

func NewPackageAnalyzer(pkg *packages.Package) *PackageAnalyzer {
	pa := &PackageAnalyzer{
		pkg: pkg,
	}
	return pa
}

// get the ast for of fileName
func (pa *PackageAnalyzer) GetASTFile(fileName string) *ast.File {
	for idx, file := range pa.pkg.CompiledGoFiles {
		if strings.HasSuffix(file, fileName) {
			return pa.pkg.Syntax[idx]
		}
	}
	return nil
}

// todo save constants into a filter struct
func (pa *PackageAnalyzer) ParseFieldList(fieldList []*ast.Field, actionInfo *OpenStackActionInfo, kind string) ([]string, bool) {
	log.Printf("-----------------parse %v:%d------------------/n", kind, len(fieldList))
	importPaths := make([]string, 0)
	for _, expr := range fieldList {
		names := pa.parseFieldNames(expr)
		typeName, packagePath := pa.parseExprTypeInfo(expr.Type)
		if typeName == "*gophercloud.ServiceClient" {
			continue
		}
		//todo support action with array parameter
		if strings.HasPrefix(typeName, "[]") {
			return importPaths, false
		}
		log.Println(names, typeName, packagePath)
		actionInfo.AddVarInfos(names, typeName, kind)
		importPaths = append(importPaths, packagePath)
	}
	return importPaths, true
}

// analyze packages and parse info
func (pa *PackageAnalyzer) DoAnalyze() *OpenstackResourceInfo {
	log.Printf("-----------analyze requestfile:  %s-----------\n", pa.pkg.Name)
	resourceInfo := NewOpenstackResourceInfo(pa.pkg.Name, pa.pkg.PkgPath)
	requestAST := pa.GetASTFile("requests.go")
	if requestAST == nil {
		return nil
	}
	for _, d := range requestAST.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if pa.checkRequestFunc(fn) {
				actionInfo := NewOpenstackActionInfo(fn.Name.String())
				log.Println("******************handle function***************** :", fn.Name)
				paramsImportPaths, parseParamsResult := pa.ParseFieldList(fn.Type.Params.List, actionInfo, "parameters")
				returnsImportPaths, parseReturnsResult := pa.ParseFieldList(fn.Type.Results.List, actionInfo, "returns")
				if parseParamsResult && parseReturnsResult {
					resourceInfo.AddAction(actionInfo)
					resourceInfo.AddImportPaths(paramsImportPaths)
					resourceInfo.AddImportPaths(returnsImportPaths)
				} else {
					log.Println("Warning: unsupport action: ", actionInfo.ActionName)
				}
			}
		}
	}
	return resourceInfo
}

// get interface type from the pkg
func (pa *PackageAnalyzer) getInterface(interfaceName string, pkg *types.Package) (*types.Type, *types.Interface) {
	if pkg == nil || !strings.Contains(pkg.Path(), "openstack") {
		return nil, nil
	}
	interfaceName = utils.GetStructName(interfaceName)
	obj := pkg.Scope().Lookup(interfaceName)
	if obj != nil {
		objType := obj.Type()
		ifaceType, ok := objType.Underlying().(*types.Interface)
		if ok {
			return &objType, ifaceType
		}
	}
	return nil, nil
}

// find the struct type that implement the interface
func (pa *PackageAnalyzer) interface2struct(ifaceType *types.Type, iface *types.Interface) (string, string) {
	tinfo := pa.pkg.TypesInfo
	log.Println("find struct for interface ", *ifaceType)
	for _, ty := range tinfo.Types {
		if types.Implements(ty.Type, iface) {
			//if ty.Type.String() != (*ifaceType).String() {
			log.Println(ty.Type)
			_, isInterface := ty.Type.Underlying().(*types.Interface)
			if !isInterface {
				log.Println(ty.Type.String())
				log.Println((*ifaceType).String())
				log.Printf("struct %v implements interface %v\n", ty.Type, *ifaceType)
				return pa.parseTypeInfo(ty.Type)
			}
		}
	}
	return "", ""
}

func (pa *PackageAnalyzer) GetPackage(ty types.Type) *types.Package {
	switch tyType := ty.(type) {
	case *types.Pointer:
		return pa.GetPackage(tyType.Elem())
	case *types.Named:
		tmp := tyType.Obj()
		if tmp.Pkg() != nil {
			return tmp.Pkg()
		} else {
			return nil
		}
	case *types.Basic:
		return nil
	case *types.Slice:
		return pa.GetPackage(tyType.Elem())
	default:
		return nil
	}
}

func (pa *PackageAnalyzer) parseExprTypeInfo(expr ast.Expr) (tyName string, packagePath string) {
	ty := pa.pkg.TypesInfo.Types[expr].Type
	//check if  type is interface,
	tyName, packagePath = pa.parseTypeInfo(ty)
	isSlice := false
	if strings.HasPrefix(tyName, "[]") {
		isSlice = true
	}
	typesPkg := pa.GetPackage(ty)
	ifaceType, iface := pa.getInterface(tyName, typesPkg)
	//ifaceType, iface := pa.getInterface(tyName, pkg.Types)
	if ifaceType != nil {
		tyName, packagePath = pa.interface2struct(ifaceType, iface)
	}
	if isSlice {
		tyName = "[]" + tyName
	}
	return
}

/*
check if the function is required(List, Create, Delete, Get...)
1. the first parameter should be * gophercloud.ServiceClient
*/
func (pa *PackageAnalyzer) checkRequestFunc(fn *ast.FuncDecl) bool {
	funcName := fn.Name.String()
	//check if the function is exported
	if utils.IsLower(funcName) {
		return false
	}
	if fn.Recv == nil { //function's Recv filed is nil, method is not
		if len(fn.Type.Params.List) != 0 {
			typeName, _ := pa.parseExprTypeInfo(fn.Type.Params.List[0].Type)
			if typeName == "*gophercloud.ServiceClient" {
				return true
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
		log.Println("error! unhandled type: ", tyType)
		return "", ""
	}
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
parse field's name and type,
a field may contain two var names
*/
func (pa *PackageAnalyzer) parseFieldNames(field *ast.Field) []string {
	names := make([]string, len(field.Names), len(field.Names))
	for idx, fieldName := range field.Names {
		names[idx] = fieldName.Name
	}
	//return filed may have no names
	if len(field.Names) == 0 {
		names = []string{""}
	}
	return names
}
