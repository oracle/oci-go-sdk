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
	"strconv"
	"golang.org/x/tools/go/ast/astutil"
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

func parseFile(filename string) (fset *token.FileSet, parsedFile *ast.File, e error)  {
	fset = token.NewFileSet()
	parsedFile, e = parser.ParseFile(fset, filename, nil, parser.ParseComments )
	return
}

func MergeImports(intoFile *ast.File, fromFile *ast.File, tokenSet *token.FileSet) {
	setOfImports := make(map[string]*ast.ImportSpec)

	for _, intoImports := range intoFile.Imports {
		setOfImports[intoImports.Path.Value] = intoImports
	}

	for _, fromImport := range fromFile.Imports {
		if setOfImports[fromImport.Path.Value] == nil {
			unquotedImport, _ := strconv.Unquote(fromImport.Path.Value)
			astutil.AddImport(tokenSet, intoFile, unquotedImport)
		}
	}
}

func DeleteComment(file *ast.File, set *token.FileSet, commentToRemove string) {
	for _, commentGroups := range file.Comments {
		newComments := make([]*ast.Comment, 0)
		for _, comment := range commentGroups.List {
			if comment.Text != commentToRemove {
				newComments = append(newComments, comment)
			}
		}
		if len(newComments) == 0 {
			commentGroups = nil
		} else {
			commentGroups.List = newComments
		}
	}
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
	 MergeImports(modelFile, pSourceFile, mf)

	fmt.Println(string(buffer.Bytes()))
}



