package helper

import (
	"encoding/json"
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

func GetLocalPackages(dir string) ([]string, error) {
	output, err := exec.Command("go", "list", fmt.Sprintf("%s/...", dir)).CombinedOutput()
	if err != nil {
		return nil, err
		os.Exit(1)
	}

	return strings.Fields(string(output)), nil
}

func GetThirdPackages(dir string) ([]string, error) {
	cmd := exec.Command("go", "list", "-m", "-json", "all")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	output = append([]byte("["), output...)
	output = append(output, []byte("]")...)

	output = []byte(strings.ReplaceAll(string(output), "}\n{", "},\n{"))

	var modules []map[string]interface{}
	if err = json.Unmarshal(output, &modules); err != nil {
		return nil, err
	}

	fmt.Println("Direct Third-party Dependencies:")
	ret := make([]string, 0)
	for _, module := range modules {
		if main, ok := module["Main"].(bool); ok && main {
			continue
		}

		if indirect, ok := module["Indirect"].(bool); ok && indirect {
			continue
		}

		if depOnly, ok := module["DepOnly"].(bool); ok && depOnly {
			continue
		}

		if path, ok := module["Path"].(string); ok {
			ret = append(ret, path)
		}
	}
	return ret, nil
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
	moduleName, err := GetModuleName(modulePath)
	if err != nil {
		return "", err
	}

	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	moduleDir := filepath.Dir(modulePath)
	midPath := strings.TrimPrefix(absFilePath, moduleDir)
	midPath = strings.TrimPrefix(midPath, "\\")
	midPath = filepath.Dir(midPath)

	// Convert file path separators to dots
	midPath = filepath.ToSlash(midPath)

	fullPkgName := moduleName + "/" + midPath

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
		if !info.IsDir() && info.Name() == "go.mod" {
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
