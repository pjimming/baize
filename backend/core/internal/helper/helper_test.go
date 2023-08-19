package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dir = `D:\GoProject\baize\backend`

func TestGetFileImports(t *testing.T) {
	ast := assert.New(t)
	imports, err := GetFileImports("D:\\GoProject\\baize\\core\\core.go")
	ast.Nil(err)

	fmt.Println("GetFileImports:")
	for _, item := range imports {
		fmt.Println(item)
	}
}

func TestGetAllPackages(t *testing.T) {
	ast := assert.New(t)
	packages, err := GetAllPackages(dir)
	ast.Nil(err)

	fmt.Println("GetAllPackages:")
	for _, item := range packages {
		fmt.Println(item)
	}
}

func TestGetAllGoFiles(t *testing.T) {
	ast := assert.New(t)
	goFiles, err := GetAllGoFiles(dir)
	ast.Nil(err)

	fmt.Println("GetAllGoFiles:")
	for _, item := range goFiles {
		fmt.Println(item)
	}
}

func TestGetPackageName(t *testing.T) {
	ast := assert.New(t)

	packageName, err := GetPackageName("D:\\GoProject\\baize\\core\\internal\\helper\\helper.go")
	ast.Nil(err)
	t.Log("GetPackageName", packageName)
}

func TestGetModulePath(t *testing.T) {
	ast := assert.New(t)
	modulePath, err := GetModulePath(dir)
	ast.Nil(err)
	t.Log(modulePath)
}

func TestShowDependency(t *testing.T) {
	ast := assert.New(t)
	//packages, err := GetAllPackages("D:\\GoProject\\baize")
	//ast.Nil(err)

	goFiles, err := GetAllGoFiles(dir)
	ast.Nil(err)

	for _, goFile := range goFiles {
		imports, err := GetFileImports(goFile)
		ast.Nil(err)

		fmt.Println(goFile, "GetAllImports")
		for _, item := range imports {
			fmt.Println(item)
		}
		fmt.Println("")
	}
}

func TestGetModuleName(t *testing.T) {
	ast := assert.New(t)
	moduleName, err := GetModuleName("D:\\GoProject\\baize\\go.mod")
	ast.Nil(err)
	t.Log("GetModuleName:", moduleName)
}

func TestGetFullPackageName(t *testing.T) {
	ast := assert.New(t)
	packageName, err := GetFullPackageName("D:\\GoProject\\baize\\go.mod",
		"D:\\GoProject\\baize\\core\\internal\\helper\\helper.go")
	ast.Nil(err)
	t.Log("GetFullPackageName:", packageName)
}
