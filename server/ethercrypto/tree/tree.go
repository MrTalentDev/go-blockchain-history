package tree

import (
	"./customsmt"
	)


func Tree(content []string) []string {
	sContent := customsmt.CreateContent(content)
	t := customsmt.CreateTree(sContent)
	return customsmt.Hashes(t)
}

