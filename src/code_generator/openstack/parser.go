package openstack

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
)

// 1. package name 用于定位参数
// todo. 函数1：若原参数/返回值类型不包含包名，则添加包名，否则直接返回
type PackageParser struct {
	Path        string
	FileName    string
	PackageName string              // package name is resource name
	Parameters  map[string][]string //参数
	Returns     map[string]string   //返回值
}

func NewPackageParser(path string) *PackageParser {
	pp := &PackageParser{
		Path:     path,
		FileName: filepath.Base(path),
	}
	return pp
}

func (pp *PackageParser) doParse() {
	set := token.NewFileSet()
	//todo using fileter to filter go file
	packs, err := parser.ParseDir(set, pp.Path, nil, 0)
	if err != nil {
		log.Fatal("Failed to parse package: ", err)
	}
	for _, pack := range packs {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					fmt.Println(fn.Name)
					fmt.Println(fn.Type)
					fmt.Println(fn.Recv)
				}
			}
		}
	}
}
