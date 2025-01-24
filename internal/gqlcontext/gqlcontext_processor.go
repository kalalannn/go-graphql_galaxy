package gqlcontext

import (
	"fmt"
	"go-graphql_galaxy/pkg/log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type key string

const PreloadContextKey key = "gorm_preload"

func countMaxDepth(fieldNode *ast.Field, currentDepth, maxDepth int) (int, error) {
	if currentDepth > maxDepth {
		return 0, fmt.Errorf("max depth limix exceeded: %d", maxDepth)
	}

	depth := 0
	for _, selection := range fieldNode.SelectionSet {
		if selection, ok := selection.(*ast.Field); ok {
			tempDepth, err := countMaxDepth(selection, currentDepth+1, maxDepth)
			if err != nil {
				return 0, err
			}
			if tempDepth > depth {
				depth = tempDepth
			}
		}
	}
	return depth + 1, nil
}

func DFSMaxDepth(opCtx *graphql.OperationContext, maxDepth int) error {
	rootNodes := opCtx.Operation.SelectionSet
	depth := 0
	for _, rootNode := range rootNodes {
		tempDepth, err := countMaxDepth(rootNode.(*ast.Field), 1, maxDepth)

		if err != nil {
			return err
		}

		if tempDepth > depth {
			depth = tempDepth
		}
	}
	log.Debug("GQLMaxDepth: %d", depth)
	return nil
}

func findPreloads(fieldNode *ast.Field, currentPath []string) [][]string {
	paths := make([][]string, 0)

	for _, selection := range fieldNode.SelectionSet {
		if selection, ok := selection.(*ast.Field); ok {
			newPath := append(currentPath,
				cases.Title(language.Und).String(selection.Name))

			if len(selection.SelectionSet) != 0 {
				childPaths := findPreloads(selection, newPath)
				if len(childPaths) > 0 {
					paths = append(paths, childPaths...)
				} else {
					paths = append(paths, newPath)
				}
			}
		}
	}
	return paths
}

func DFSPreload(rootCtx *graphql.RootFieldContext) [][]string {
	preloads := findPreloads(rootCtx.Field.Field, []string{})

	log.Debug("Preloads for (%s): %v", rootCtx.Field.Field.Name, preloads)
	return preloads
}
