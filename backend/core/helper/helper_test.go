package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dir = `D:\GoProject\baize\backend`
var macOsDir = `/Users/panjiangming/Project/baize/backend`
var moduleName = `github.com/pjimming/baize`

func TestGetFileImports(t *testing.T) {
	ast := assert.New(t)
	imports, err := GetFileImports("D:\\GoProject\\baize\\core\\core.go")
	ast.Nil(err)

	fmt.Println("GetFileImports:")
	for _, item := range imports {
		fmt.Println(item)
	}
}

func TestGetLocalPackages(t *testing.T) {
	ast := assert.New(t)
	packages, err := GetLocalPackages(macOsDir)
	ast.Nil(err)

	fmt.Println("GetLocalPackages:")
	for _, item := range packages {
		fmt.Println(item)
	}
}

func TestGetAllGoFiles(t *testing.T) {
	ast := assert.New(t)
	goFiles, err := GetAllGoFiles(macOsDir)
	ast.Nil(err)

	fmt.Println("GetAllGoFiles:")
	for _, item := range goFiles {
		fmt.Println(item.Name, item.Size)
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
	modulePath, err := GetModulePath(macOsDir)
	ast.Nil(err)
	t.Log(modulePath)
}

func TestShowDependency(t *testing.T) {
	ast := assert.New(t)
	//packages, err := GetLocalPackages("D:\\GoProject\\baize")
	//ast.Nil(err)

	goFiles, err := GetAllGoFiles(dir)
	ast.Nil(err)

	for _, goFile := range goFiles {
		imports, err := GetFileImports(goFile.Name)
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
	packageName, err := GetFullPackageName(`/Users/panjiangming/Project/baize/backend/go.mod`,
		`/Users/panjiangming/Project/baize/backend/core/helper/helper.go`)
	ast.Nil(err)
	t.Log("GetFullPackageName:", packageName)
}

func TestGetThirdPackages(t *testing.T) {
	ast := assert.New(t)
	thirdPackages, err := GetThirdPackages(`/Users/panjiangming/Project/baize/backend`)
	ast.Nil(err)
	fmt.Println("GetThirdPackages:")
	for _, item := range thirdPackages {
		fmt.Println(item.Name, item.Version)
	}
}

func TestGetModuleVersion(t *testing.T) {
	ast := assert.New(t)
	moduleVersion, err := GetModuleVersion(moduleName, macOsDir)
	ast.Nil(err)
	t.Log(moduleVersion)
}
