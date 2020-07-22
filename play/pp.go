package main

import (
	helper "../ant"
	"github.com/dsymonds/gotoc/parser"
	"io/ioutil"
	"strings"
)

func main() {

	const DIR_PROTOS = `C:\Go\_gopath\src\ms\ant\play\`
	files, err := ioutil.ReadDir(DIR_PROTOS)
	helper.NoErr(err)
	protos := make([]string, 0, len(files))
	for _, f := range files {
		if strings.Contains(f.Name(), ".proto") {
			protos = append(protos, f.Name())
		}
	}
    _ = protos

	ast, err := parser.ParseFiles([]string{"1.proto"}, []string{"./", DIR_PROTOS})
    helper.NoErr(err)
	helper.PertyPrint(ast.Files)
}
