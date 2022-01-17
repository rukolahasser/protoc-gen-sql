package main

import (
	"flag"
	"fmt"
	"github.com/rukolahasser/protoc-gen-sql/sqlgen"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
)

var (
	fileName     string
	version      string
	contractName string
	packageName  string
	owner        string
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Fprintf(os.Stdout, "See README.md for usage information.\n")
		os.Exit(0)
	}

	flag.StringVar(&fileName, "file", "", "Generated file name.")
	flag.StringVar(&version, "version", "0_0_0", "version number.")
	flag.StringVar(&contractName, "contract", "", "contract name.")
	flag.StringVar(&packageName, "package", "", "package name.")
	flag.StringVar(&owner, "owner", "", "owner.")

	protogen.Options{
		ParamFunc: flag.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			sqlgen.GenerateUpFile(gen, f, fileName, packageName, owner, contractName, version)
			sqlgen.GenerateDownFile(gen, f, fileName, packageName, owner, contractName, version)
		}
		return nil
	})
}
