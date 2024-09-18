/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/util"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/ast/astutil"
)

var diffGen = &cobra.Command{
	Use:   "diff-gen",
	Short: "generates go models from the given swagger doc that are not already declared in the given .go file",
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}
		if !strings.HasSuffix(args[0], ".go") {
			return errors.New("First positional argument must be a '.go' file")
		}
		if !strings.HasSuffix(args[1], ".json") {
			return errors.New("Second positional argument must be a '.json' swagger file")
		}
		return nil
	},
	Long: `Accepts two file paths: .go file & swagger file
Command generates go models from the given swagger doc and removes the Types and Consts from 
the newly generated models that were declared in the given .go file before writing to stdOut`,
	RunE: func(cmd *cobra.Command, args []string) error {
		packageName, _ := cmd.Flags().GetString("package")
		// Parse the first Go file into an AST
		node, err := parser.ParseFile(token.NewFileSet(), args[0], nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("Error parsing file:%w", err)
		}

		exportedTypes, exportedConsts := getExportedModels(node)

		code, err := generateModels(args[1], packageName)
		if err != nil {
			return fmt.Errorf("generating models from swagger file:%w", err)
		}

		fs := token.NewFileSet()
		replaceNode, err := parser.ParseFile(fs, "models.gen.go", code, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("Error parsing file:%w", err)
		}

		newnode := filterNodes(replaceNode, exportedTypes, exportedConsts)

		return format.Node(os.Stdout, fs, newnode)
	},
}

func getExportedModels(node ast.Node) (map[string]bool, map[string]bool) {
	// Inspect the AST of file and record exported Types and Const's
	exportedConsts, exportedTypes := map[string]bool{}, map[string]bool{}
	ast.Inspect(node, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE && genDecl.Tok != token.CONST {
			return true
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				//got to be a const if it's not a token
				valueSpec := spec.(*ast.ValueSpec)
				for _, name := range valueSpec.Names {
					if ast.IsExported(name.Name) {
						exportedConsts[name.Name] = true
					}
				}
			} else {
				if ast.IsExported(typeSpec.Name.Name) {
					exportedTypes[typeSpec.Name.Name] = true
				}
			}
		}

		return true
	})
	return exportedConsts, exportedTypes
}

func generateModels(filepath string, packageName string) (string, error) {
	swagger, err := util.LoadSwagger(filepath)
	if err != nil {
		return "", fmt.Errorf("error loading swagger spec in %s\n: %s\n", filepath, err)
	}
	return codegen.Generate(swagger, codegen.Configuration{
		PackageName: packageName,
		Generate: codegen.GenerateOptions{
			Models: true,
		},
	})
}

func filterNodes(node ast.Node, typesToFilter map[string]bool, constsToFilter map[string]bool) ast.Node {
	return astutil.Apply(node, func(c *astutil.Cursor) bool {
		n := c.Node()
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || (genDecl.Tok != token.TYPE && genDecl.Tok != token.CONST) {
			return true
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				//got to be a const if it's not a token
				valueSpec := spec.(*ast.ValueSpec)
				for _, name := range valueSpec.Names {
					if constsToFilter[name.Name] {
						c.Delete()
						return true
					}
				}
			} else {
				if ast.IsExported(typeSpec.Name.Name) && typesToFilter[typeSpec.Name.Name] {
					c.Delete()
					return true
				}
			}
		}
		return true
	}, nil)
}

func init() {
	rootCmd.AddCommand(diffGen)
	diffGen.Flags().StringP("package", "p", "energy", "target package name")
}
