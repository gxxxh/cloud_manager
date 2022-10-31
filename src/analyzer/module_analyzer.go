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

//todo save constants into a filter struct

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
				parseFieldList := func(fieldList []*ast.Field, kind string) bool {
					log.Printf("-----------------%v:%d------------------/n", kind, len(fieldList))
					for _, expr := range fieldList {
						//a field may contain two name with the same type
						names := pa.parseFieldNames(expr)
						typeName, packagePath := pa.parseExprTypeInfo(expr.Type, pkg)
						if typeName == "*gophercloud.ServiceClient" {
							continue
						}
						//todo support action with array parameter
						if strings.HasPrefix(typeName, "[]") {
							return false
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
					return true
				}
				if parseFieldList(fn.Type.Params.List, "parameters") && parseFieldList(fn.Type.Results.List, "returns") {
					resourceInfo.AddAction(actionInfo)
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
	interfaceName = GetStructName(interfaceName)
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
func (pa *PackageAnalyzer) interface2struct(ifaceType *types.Type, iface *types.Interface, tinfo *types.Info) (string, string) {
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

func (pa *PackageAnalyzer) parseExprTypeInfo(expr ast.Expr, pkg *packages.Package) (tyName string, packagePath string) {
	ty := pkg.TypesInfo.Types[expr].Type
	//check if  type is interface,
	tyName, packagePath = pa.parseTypeInfo(ty)
	isSlice := false
	if strings.HasPrefix(tyName, "[]") {
		isSlice = true
	}
	ifaceType, iface := pa.getInterface(tyName, pkg.Types)
	if ifaceType != nil {
		tyName, packagePath = pa.interface2struct(ifaceType, iface, pkg.TypesInfo)
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
func (pa *PackageAnalyzer) checkValidFunction(fn *ast.FuncDecl) bool {
	funcName := fn.Name.String()
	//check if the function is exported
	if utils.IsLower(funcName) {
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
parse field's name and type
*/
func (pa *PackageAnalyzer) parseFieldNames(field *ast.Field) []string {
	names := make([]string, len(field.Names), len(field.Names))
	for idx, fieldName := range field.Names {
		names[idx] = fieldName.Name
	}
	return names
}
