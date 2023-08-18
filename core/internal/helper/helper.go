package helper

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetFileImports(filename string) ([]string, error) {
	ret := make([]string, 0)

	// Create a new file set
	fset := token.NewFileSet()

	// Parse the Go source file
	node, err := parser.ParseFile(fset, filename, nil, parser.ImportsOnly)
	if err != nil {
		return nil, err
	}

	// Get the imported packages
	for _, importSpec := range node.Imports {
		ret = append(ret, importSpec.Path.Value)
	}
	return ret, nil
}

func GetAllPackages(dir string) ([]string, error) {
	output, err := exec.Command("go", "list", fmt.Sprintf("%s/...", dir)).CombinedOutput()
	if err != nil {
		return nil, err
		os.Exit(1)
	}

	return strings.Fields(string(output)), nil
}

func GetAllGoFiles(dir string) (goFiles []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return goFiles, nil
}

func GetFullPackageName(modulePath, filePath string) (string, error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	relPath, err := filepath.Rel(modulePath, absFilePath)
	if err != nil {
		return "", err
	}

	// Convert file path separators to dots
	pkgPath := filepath.ToSlash(relPath)

	// Remove .go extension and replace slashes with dots
	pkgPath = strings.TrimSuffix(pkgPath, ".go")
	pkgPath = strings.TrimPrefix(pkgPath, "/")
	pkgPath = strings.Replace(pkgPath, "/", ".", -1)

	// Get the directory name of the package
	pkgDir := filepath.Dir(pkgPath)

	// Construct the complete package name
	modulePath = strings.TrimPrefix(modulePath, "module ")
	fullPkgName := modulePath + "/" + pkgDir

	return fullPkgName, nil
}

func GetPackageName(filePath string) (string, error) {
	fileSet := token.NewFileSet()

	// 解析 Go 文件，仅解析 package 语句
	node, err := parser.ParseFile(fileSet, filePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}

	return node.Name.Name, nil
}

func GetModulePath(dir string) (modulePath string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mod") {
			modulePath = path
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return modulePath, nil
}

func GetModuleName(goModPath string) (string, error) {
	// 读取 go.mod 文件的内容
	modFileContent, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	// 解析第一行获取模块路径
	firstLine := strings.SplitN(string(modFileContent), "\n", 2)[0]
	modulePath := strings.TrimPrefix(firstLine, "module ")

	return modulePath, nil
}
