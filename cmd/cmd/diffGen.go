package cmd

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"os"
	"strings"

	"github.com/oapi-codegen/oapi-codegen/v2/pkg/codegen"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/util"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/ast/astutil"
)

var diffGen = &cobra.Command{
	Use:   "diff-gen <path to go package> <path to openapi/swagger doc>",
	Short: "generates go models from an openapi/swagger doc that are not already declared in the given go package",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}
		info, err := os.Stat(args[0])
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				return errors.New("First positional argument is not a valid local filepath")
			}
			return fmt.Errorf("error checking filepath:%w", err)
		}
		if !info.IsDir() {
			return fmt.Errorf("First positional argument %s is not a directory", args[0])
		}
		if !strings.HasSuffix(args[1], ".json") {
			return errors.New("Second positional argument must be a '.json' swagger file")
		}
		return nil
	},
	Long: `Accepts two file paths: 1)go package directory & 2)swagger/openAPI file
Command generates go models from the given swagger doc then removes the Types and Consts 
from the newly generated models that were already declared in the given go package.

By default the diff-models are written to a file named the same as the openAPI/swagger 
file, with the '.json' extension changed to '.gen.go'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		targetFilePath, _ := cmd.Flags().GetString("target")
		packagePath := args[0]
		swaggerPath := args[1]

		pack, err := parsePackage(packagePath)
		if err != nil {
			return fmt.Errorf("parsing package:%w", err)
		}

		code, err := genSwaggerModels(swaggerPath, pack.packageName)
		if err != nil {
			return fmt.Errorf("generating models from swagger file:%w", err)
		}

		fs := token.NewFileSet()
		replaceNode, err := parser.ParseFile(fs, "models.gen.go", code, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("Error parsing file:%w", err)
		}

		//Create commentmap from newly generated file
		cmap := ast.NewCommentMap(fs, replaceNode, replaceNode.Comments)

		newnode := filterNodes(replaceNode, pack.exportedTypes, pack.exportedConsts)

		//filter original comment groups by nodes that we keep
		nf, ok := newnode.(*ast.File)
		if !ok {
			return fmt.Errorf("cast ast.Node as *ast.File:%w", err)
		}
		nf.Comments = cmap.Filter(newnode).Comments()

		if targetFilePath == "" {
			targetFilePath = strings.Replace(swaggerPath, ".json", ".gen.go", 1)
		}
		outputFile, err := os.Create(targetFilePath)
		if err != nil {
			return fmt.Errorf("Error creating output file:%w", err)
		}
		defer outputFile.Close()
		return printer.Fprint(outputFile, fs, newnode)
	},
}

type packageModel struct {
	packageName    string
	exportedTypes  map[string]bool
	exportedConsts map[string]bool
}

func NewPackageModel() packageModel {
	return packageModel{
		exportedTypes:  map[string]bool{},
		exportedConsts: map[string]bool{},
	}
}

func parsePackage(packagePath string) (packageModel, error) {

	pm := NewPackageModel()
	fileSet := token.NewFileSet()

	// Parse all Go files in the package directory
	packages, err := parser.ParseDir(fileSet, packagePath, nil, 0)
	if err != nil {
		return pm, fmt.Errorf("Error parsing package:%w", err)
	}

	for _, pkg := range packages {
		if pm.packageName == "" && !strings.HasSuffix(pkg.Name, "_test") {
			pm.packageName = pkg.Name
		}
		for _, file := range pkg.Files {

			// Inspect the AST of file and record exported Types and Const's
			ast.Inspect(file, func(n ast.Node) bool {
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
								pm.exportedConsts[name.Name] = true
							}
						}
					} else {
						if ast.IsExported(typeSpec.Name.Name) {
							pm.exportedTypes[typeSpec.Name.Name] = true
						}
					}
				}

				return true
			})
		}
	}
	return pm, nil
}

func genSwaggerModels(filepath string, packageName string) (string, error) {
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

func filterNodes(node *ast.File, typesToFilter map[string]bool, constsToFilter map[string]bool) ast.Node {

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
	diffGen.Flags().StringP("target", "t", "", "target file - defaults to {swagger_file}.gen.go")
}
