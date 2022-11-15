package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"multicloud_service/src/utils"
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
		////analyzing request file
		resourceInfo := packageAnalyzer.AnalyzeRequestFile()
		if resourceInfo == nil {
			continue
		}

		if pageExtractInfos := packageAnalyzer.AnalyseResultFile(); len(pageExtractInfos) != 0 {
			ma.MapPageExtractInfo2Action(pkg.PkgPath, pkg.Name, pageExtractInfos, resourceInfo.ActionInfos)
		}
		resourceInfo.RemoveInvalidActions()
		resourceInfo.GenRequestImportPaths()
		resourceInfo.GenResultImportPaths()
		resourceInfos = append(resourceInfos, resourceInfo)
	}
	return resourceInfos, err
}

func (ma *ModuleAnalyzer) MapPageExtractInfo2Action(pkgPath string, pkgName string, pageExtractInfos []*PageExtractInfo, actionInfos []*OpenStackActionInfo) bool {
	if len(pageExtractInfos) == 0 {
		return false
	}
	//筛选返回值是page的action
	resource2actionInfo := make(map[string]*OpenStackActionInfo, 0)
	validActionNames := make([]string, 0)
	for _, actionInfo := range actionInfos {
		for _, varInfo := range actionInfo.ActionReturns {
			if strings.Contains(varInfo.TypeName, "pagination.Page") {
				actionName := utils.ParseResourceName(actionInfo.ActionName, "List")
				if actionName == "list" {
					actionName = strings.ToLower(pkgName)
				}
				resource2actionInfo[actionName] = actionInfo
				validActionNames = append(validActionNames, actionName)
				break
			}
		}
	}
	resource2pageExtractInfo := make(map[string]*PageExtractInfo, 0)
	validPageFuncNames := make([]string, 0)
	for _, pageExtractInfo := range pageExtractInfos {
		pageFuncName := utils.ParseResourceName(pageExtractInfo.FuncName, "Extract")
		resource2pageExtractInfo[pageFuncName] = pageExtractInfo
		validPageFuncNames = append(validPageFuncNames, pageFuncName)
	}
	if len(resource2actionInfo) == 0 {
		return false
	}
	//1. 若只有一个extract function 直接赋值
	if len(resource2pageExtractInfo) == 1 && len(resource2actionInfo) == 1 {
		for _, tmpAction := range resource2actionInfo {
			tmpAction.PageExtractInfo = pageExtractInfos[0]
		}
		return true
	}
	//2. 一个extract对应多个action
	if len(resource2pageExtractInfo) == 1 {
		for _, tmpAction := range resource2actionInfo {
			tmpAction.PageExtractInfo = pageExtractInfos[0]
		}
		return true
	}
	//2. 若能产生的后缀相同，则直接赋值。
	if utils.CompareSlice(validActionNames, validPageFuncNames) {
		for _, pageFuncName := range validPageFuncNames {
			resource2actionInfo[pageFuncName].PageExtractInfo = resource2pageExtractInfo[pageFuncName]
		}
		return true
	}
	//3. 若只有一个名称不同，则也可以匹配
	if diffNames := utils.DiffSlice(validActionNames, validPageFuncNames); len(diffNames) == 2 {
		for _, pageFuncName := range validPageFuncNames {
			if !diffNames.Has(pageFuncName) {
				resource2actionInfo[pageFuncName].PageExtractInfo = resource2pageExtractInfo[pageFuncName]
			} else {
				for name, _ := range diffNames {
					if name != pageFuncName {
						resource2actionInfo[name.(string)].PageExtractInfo = resource2pageExtractInfo[pageFuncName]
					}
				}
			}
		}
		return true
	}
	//3. 手动map
	switch pkgPath {
	//case "github.com/gophercloud/gophercloud/openstack/compute/v2/servers":
	//	resource2actionInfo["servers"].PageExtractInfo = resource2pageExtractInfo["servers"]
	//	resource2actionInfo["addresses"].PageExtractInfo = resource2pageExtractInfo["addresses"]
	//	resource2actionInfo["addressbynetwork"].PageExtractInfo = resource2pageExtractInfo["networkaddresses"]
	//case "github.com/gophercloud/gophercloud/openstack/clustering/v1/clusters":
	//	resource2actionInfo["clusters"].PageExtractInfo = resource2pageExtractInfo["clusters"]
	//	resource2actionInfo["clusterpolicies"].PageExtractInfo = resource2pageExtractInfo["clusterpolicies"]
	case "github.com/gophercloud/gophercloud/openstack/db/v1/configurations":
		resource2actionInfo["configurations"].PageExtractInfo = resource2pageExtractInfo["configs"]
		resource2actionInfo["datastoreparams"].PageExtractInfo = resource2pageExtractInfo["params"]
		resource2actionInfo["globalparams"].PageExtractInfo = resource2pageExtractInfo["params"]
		//todo 有一个action listinstance需要调用instance包的extract
		return true
	case "github.com/gophercloud/gophercloud/openstack/identity/v3/roles":
		resource2actionInfo["roles"].PageExtractInfo = resource2pageExtractInfo["roles"]
		resource2actionInfo["assignments"].PageExtractInfo = resource2pageExtractInfo["roleassignments"]
		resource2actionInfo["assignmentsonresource"].PageExtractInfo = resource2pageExtractInfo["roles"]
		return true
	case "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/bgp/speakers":
		resource2actionInfo["speakers"].PageExtractInfo = resource2pageExtractInfo["bgpspeakers"]
		resource2actionInfo["getadvertisedroutes"].PageExtractInfo = resource2pageExtractInfo["advertisedroutes"]
		return true
	case "github.com/gophercloud/gophercloud/openstack/objectstorage/v1/containers":
		resource2actionInfo["containers"].PageExtractInfo = resource2pageExtractInfo["info"]
		return true
	case "github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects":
		resource2actionInfo["objects"].PageExtractInfo = resource2pageExtractInfo["info"]
		return true
	case "github.com/gophercloud/gophercloud/openstack/orchestration/v1/stackresources":
		resource2actionInfo["stackresources"].PageExtractInfo = resource2pageExtractInfo["resources"]
		resource2actionInfo["types"].PageExtractInfo = resource2pageExtractInfo["resourcetypes"]
		return true
	default:
		log.Println("no rules mapped for package ", pkgPath)
	}
	return false
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

func (pa *PackageAnalyzer) Field2VarInfos(fieldList []*ast.Field) VarInfos {
	varInfos := NewVarInfos()
	for _, expr := range fieldList {
		names := pa.parseFieldNames(expr)
		typeName, packagePath := pa.parseExprTypeInfo(expr.Type)
		varInfos.Add(names, typeName, packagePath)
	}
	return varInfos
}

// analyze packages and parse info
func (pa *PackageAnalyzer) AnalyzeRequestFile() *OpenstackResourceInfo {
	log.Printf("-----------analyze requestfile:  %s-----------\n", pa.pkg.Name)
	resourceInfo := NewOpenstackResourceInfo(pa.pkg.Name, pa.pkg.PkgPath)
	requestAST := pa.GetASTFile("requests.go")
	if requestAST == nil {
		return nil
	}
	for _, d := range requestAST.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if pa.checkValidFunc(fn, "gophercloud.ServiceClient") {
				log.Println("******************handle function***************** :", fn.Name)
				actionInfo := NewOpenstackActionInfo(fn.Name.String())
				paramsVarInfos := pa.Field2VarInfos(fn.Type.Params.List)
				//no need to import
				//paramsImportPaths.Delete("github.com/gophercloud/gophercloud")
				returnsVarInfos := pa.Field2VarInfos(fn.Type.Results.List)
				actionInfo.AddVarInfos(paramsVarInfos, "parameters")
				actionInfo.AddVarInfos(returnsVarInfos, "returns")
				// analyze result extract func
				// todo check if all return var info's len is 1
				if len(returnsVarInfos) >= 1 &&
					strings.HasSuffix(returnsVarInfos[0].TypeName, "Result") {
					actionInfo.ResultExtractInfo = pa.ParseResultExtractInfo(fn.Type.Results.List[0].Type)
				}
				//todo 生成importPath, 注意paramsImportPaths要删除一个
				resourceInfo.AddAction(actionInfo)
			}
		}
	}

	return resourceInfo
}

// get result's extract function's return info
func (pa *PackageAnalyzer) ParseResultExtractReturns(ty *types.Named) *ResultExtractInfo {
	resultExtractInfo := NewResultExtractInfo()
	for i := 0; i < ty.NumMethods(); i++ {
		tmpMethod := ty.Method(i)
		if tmpMethod.Name() == "Extract" ||
			tmpMethod.Name() == "ExtractErr" ||
			strings.Contains(tmpMethod.Name(), "Extract") && !strings.HasSuffix(tmpMethod.Name(), "Into") {
			methodType, ok := tmpMethod.Type().(*types.Signature)
			if !ok {
				log.Println("error, the extract function can't turned to a signaure", ty)
			}
			if resultExtractInfo.FuncName != "" {
				log.Println("error, not only one Extract/ExtractErr func for thies result ", ty)
			}
			varInfos := NewVarInfos()
			for i := 0; i < methodType.Results().Len(); i++ {
				result := methodType.Results().At(i)
				typeName, typeImportPath := pa.parseTypeInfo(result.Type())
				varInfo := NewVarInfo("", typeName, typeImportPath)
				varInfos.AddVarInfo(varInfo)
			}
			resultExtractInfo.FuncName = tmpMethod.Name()
			resultExtractInfo.ReturnInfo = varInfos
		}
	}
	if resultExtractInfo.FuncName == "" {
		log.Printf("error, no extract func for type %v in package %v", ty, pa.pkg.PkgPath)
	}
	return resultExtractInfo
}

// input: result Type
// ouput: result Type's extract function
func (pa *PackageAnalyzer) ParseResultExtractInfo(expr ast.Expr) *ResultExtractInfo {
	log.Println("parse result extract info for ", expr)
	ty := pa.pkg.TypesInfo.Types[expr].Type
	tyNamed, ok := ty.(*types.Named)
	if !ok {
		log.Println("error, the result type should be a named type", ty)
		return nil
	}
	if tyNamed.NumMethods() != 0 {
		resultExtractInfo := pa.ParseResultExtractReturns(tyNamed)
		log.Println("find result extract info for ", expr)
		return resultExtractInfo
	} else {
		tyStruct, ok := tyNamed.Underlying().(*types.Struct)
		if !ok {
			log.Println("error, the result underlying type is not struct", ty)
		}
		for i := 0; i < tyStruct.NumFields(); i++ {
			fieldType := tyStruct.Field(i).Type()
			fieldTypeNamed, ok := fieldType.(*types.Named)
			if ok && fieldTypeNamed.NumMethods() > 0 && strings.Contains(fieldTypeNamed.Obj().Name(), "Result") {
				resultExtractInfo := pa.ParseResultExtractReturns(fieldTypeNamed)
				log.Println("find result extract info for ", expr)
				return resultExtractInfo
			}
		}
	}
	return nil
}

func (pa *PackageAnalyzer) AnalyseResultFile() []*PageExtractInfo {
	log.Printf("-----------analyze resultfile:  %s-----------\n", pa.pkg.Name)
	pageExtractInfos := make([]*PageExtractInfo, 0)
	resultAST := pa.GetASTFile("results.go")
	if resultAST == nil {
		return pageExtractInfos
	}
	for _, d := range resultAST.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			fnName := fn.Name.String()
			log.Println("handle result function", fnName)
			if pa.checkValidFunc(fn, "pagination.Page") &&
				strings.HasPrefix(fnName, "Extract") &&
				!strings.HasSuffix(fnName, "Into") &&
				!strings.HasSuffix(fnName, "Base") &&
				len(fn.Type.Params.List) == 1 &&
				fnName != "Extract" {
				pageExtractInfo := NewPageExtractInfo(fnName)
				returnVarInfos := pa.Field2VarInfos(fn.Type.Results.List)
				pageExtractInfo.ReturnInfo = returnVarInfos
				pageExtractInfos = append(pageExtractInfos, pageExtractInfo)
			}
		}
	}

	return pageExtractInfos
}

// get interface type from the pkg
func (pa *PackageAnalyzer) getInterface(interfaceName string, pkg *types.Package) *types.Interface {
	if pkg == nil || !strings.Contains(pkg.Path(), "openstack") {
		return nil
	}
	interfaceName = utils.GetStructName(interfaceName)
	obj := pkg.Scope().Lookup(interfaceName)
	if obj != nil {
		objType := obj.Type()
		ifaceType, ok := objType.Underlying().(*types.Interface)
		if ok {
			return ifaceType
		}
	}
	return nil
}

// find the struct type that implement the interface
func (pa *PackageAnalyzer) interface2struct(iface *types.Interface) (string, string) {
	tinfo := pa.pkg.TypesInfo
	log.Println("find struct for interface ", iface)
	for _, ty := range tinfo.Types {
		if types.Implements(ty.Type, iface) {
			//if ty.Type.String() != (*ifaceType).String() {
			log.Println(ty.Type)
			_, isInterface := ty.Type.Underlying().(*types.Interface)
			if !isInterface {
				log.Printf("struct %v implements interface %v\n", ty.Type, iface)
				//return pa.parseTypeInfo(ty.Type)
				tyName, packagePath := pa.parseTypeInfo(ty.Type)
				if tyName != "" {
					return tyName, packagePath
				}
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
	iface := pa.getInterface(tyName, typesPkg)
	//ifaceType, iface := pa.getInterface(tyName, pkg.Types)
	if iface != nil {
		tyName, packagePath = pa.interface2struct(iface)
		if tyName == "" {
			log.Println("error, not find struct for interface", iface)
		}
		if isSlice {
			tyName = "[]" + tyName
		}
	}
	return
}

/*
check if the function is required
1. it should be an exported function
2. the first parameter should be required
*/
func (pa *PackageAnalyzer) checkValidFunc(fn *ast.FuncDecl, paraName string) bool {
	funcName := fn.Name.String()
	//check if the function is exported
	if utils.IsLower(funcName) {
		return false
	}
	if fn.Recv == nil { //function's Recv filed is nil, method is not
		if len(fn.Type.Params.List) != 0 {
			typeName, _ := pa.parseExprTypeInfo(fn.Type.Params.List[0].Type)
			//if typeName == paraName {
			if strings.Contains(typeName, paraName) {
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
	case *types.Map:
		keyName, keyPath := pa.parseTypeInfo(tyType.Key())
		valueName, valuePath := pa.parseTypeInfo(tyType.Elem())
		//todo keyPath and valuePath may be different
		return "map[" + keyName + "]" + valueName, keyPath + valuePath
	case *types.Interface:
		log.Println("interface type: ", tyType)
		return "interface{}", ""
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
