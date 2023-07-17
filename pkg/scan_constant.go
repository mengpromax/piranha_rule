package pkg

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"

	utils "github.com/mengpromax/piranha_rule/pkg/util"
)

func ScanAllConstantsMapping(sourcePath string, grayKeys []string, fSet *token.FileSet) (constantMapping map[string]string, err error) {
	if fSet == nil {
		fSet = token.NewFileSet()
	}

	grayKeySet := utils.CreateStringSet(grayKeys)

	dirEntries, err := os.ReadDir(sourcePath)
	if err != nil {
		return nil, err
	}

	constantMapping = make(map[string]string)

	for _, entry := range dirEntries {
		curEntryPath := path.Join(sourcePath, entry.Name())
		var curEntryMapping map[string]string

		if entry.IsDir() {
			subDirConstantsMapping, err := ScanAllConstantsMapping(curEntryPath, grayKeys, fSet)
			if err != nil {
				return nil, err
			}
			curEntryMapping = subDirConstantsMapping
		} else {
			ext := path.Ext(entry.Name())
			if ext != ".go" {
				continue
			}

			// 需要处理的文件
			subFileConstantsMapping, err := scanConstantsMappingForEachFile(curEntryPath, fSet)
			if err != nil {
				return nil, err
			}
			curEntryMapping = subFileConstantsMapping
		}

		for constantName, constantValue := range curEntryMapping {
			if grayKeySet.Contains(constantValue) {
				constantMapping[constantName] = constantValue
			}
		}
	}

	return constantMapping, nil
}

func scanConstantsMappingForEachFile(filePath string, fSet *token.FileSet) (constantMapping map[string]string, err error) {
	f, err := parser.ParseFile(fSet, filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	constantMapping = make(map[string]string)
	for k, v := range f.Scope.Objects {
		if v.Kind == ast.Con {
			d := v.Decl.(*ast.ValueSpec)

			if len(d.Values) == 0 {
				continue
			}

			lit, ok := d.Values[0].(*ast.BasicLit)
			if !ok {
				continue
			}

			if lit.Kind != token.STRING {
				continue
			}

			// 只记录 string 类型，并 trim 掉双引号
			constantMapping[k] = strings.Trim(lit.Value, "\"")
		}
	}

	return constantMapping, nil
}
