package main

import (
	"os"

	"github.com/emicklei/proto"
	"ms/sun/shared/helper"
)

func main() {
    wd, _ := os.Getwd()
    print(wd)
    print("\n")
	reader, e1 := os.Open(`C:\Go\_gopath\src\ms\ant\play\2.proto`)
	defer reader.Close()
	parser := proto.NewParser(reader)
	definition, e2 := parser.Parse()
	formatter := proto.NewFormatter(os.Stdout, " ")
	formatter.Format(definition)

    for _, each := range definition.Elements {
        if msg, ok := each.(*proto.Message); ok {
            helper.PertyPrint(msg)

        }

        if msg, ok := each.(*proto.Option); ok {
            helper.PertyPrint(msg)

        }

        if msg, ok := each.(*proto.Enum); ok {
            helper.PertyPrint(msg)

        }
    }


	helper.NoErr(e1)
	helper.NoErr(e2)
}
