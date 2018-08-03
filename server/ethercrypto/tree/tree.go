package tree

import (
	"./customsmt"
	"log"
)


func Tree() {

	var content []string
	content = append(content, "a")
	content = append(content, "b")
	content = append(content, "c")
	content = append(content, "d")

	sContent := customsmt.CreateContent(content)

	t := customsmt.CreateTree(sContent)

	log.Println(customsmt.Hashes(t))

	var content2 []string
	content2 = append(content2, "hv1ly")
	content2 = append(content2, "cl444ivv7l")

	sContent2 := customsmt.CreateContent(content2)

	customsmt.RewriteTree(sContent2, t)

	log.Println(customsmt.Hashes(t))
	log.Println(customsmt.GetMerkleRoot(t))
	log.Println(customsmt.VerifySpecificLeaf(t, sContent2[1]))
	log.Println(customsmt.VerifyAll(t))

}

