package main

import (
	"fmt"
	"flag"
	"os"
	"go/token"
	"go/parser"
	"go/ast"
	"bytes"
	"go/printer"
)

var (
	flagFileName = flag.String("partOf", "", "file to merge the model into")
)

func dieGracefully(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(0)
	}
}

func createOrOpen(filename string) (file *os.File, e error) {
	if _ , err := os.Stat(filename); os.IsNotExist(err) {
		//file doesn't exist
		file, e = os.Create(filename)
		return
	}
	file, e = os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0600)
	return
}

func parseFile(filename string) (fset *token.FileSet, parsedFile *ast.File, e error)  {
	fset = token.NewFileSet()
	parsedFile, e = parser.ParseFile(fset, filename, nil, parser.ParseComments )
	return
}

func mergeImports(intoFile *ast.File, fromFile *ast.File) *ast.File  {
	setOfImports := make(map[string]*ast.ImportSpec)
	newImports := make([]*ast.ImportSpec,0)

	for _, intoImports := range intoFile.Imports {
		setOfImports[intoImports.Path.Value] = intoImports
		newImports = append(newImports, intoImports)
	}

	for _, fromImport := range fromFile.Imports {
		if setOfImports[fromImport.Path.Value] == nil {
			newImports = append(newImports, fromImport)
		}
	}
	intoFile.Imports = newImports
	return intoFile
}


func main() {
	flag.Parse()
	fmt.Println("Initializing")
	sourceFile := os.Getenv("GOFILE")
	mergeIntoFile := *flagFileName + ".go"
	if sourceFile == mergeIntoFile {
		return
	}
	fmt.Println(sourceFile, "->", mergeIntoFile)
	// Parses files
	_, pSourceFile, e := parseFile(sourceFile)
	dieGracefully(e)

	mf, modelFile, e := parseFile(mergeIntoFile)
	dieGracefully(e)

	//MergeImports
	buffer := bytes.NewBuffer(nil)
	mergedFile := mergeImports(modelFile, pSourceFile)

	//This prints the ast for debugging
	//ast.Fprint(buffer, mf, mergedFile, nil)

	//This prints the ast as code
	printer.Fprint(buffer, mf, mergedFile)
	fmt.Println(string(buffer.Bytes()))

	//MergeBody
}



