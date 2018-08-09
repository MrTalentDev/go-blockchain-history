package customsmt

import (
	"log"
	"../smerkletree"
	"github.com/miguelmota/go-solidity-sha3"
)

//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	X string
}

//CalculateHashBytes hashes the values of a TestContent
func (t TestContent) CalculateHashBytes() ([]byte, error) {
	hash := solsha3.SoliditySHA3(
		solsha3.Bytes32([]byte(t.X)),
	)
	return hash, nil
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other smerkletree.Content) (bool, error) {
	return t.X == other.(TestContent).X, nil
}

func CreateContent(content []string) []smerkletree.Content {
	var list []smerkletree.Content
	for i := 0; i < len(content); i++ {
		list = append(list, TestContent{X: content[i]})
	}
	return list
}

func CreateTree(list []smerkletree.Content) (*smerkletree.MerkleTree) {
	//Create a new Merkle Tree from the list of Content
	t,  err := smerkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func RewriteTree(content []smerkletree.Content, tree *smerkletree.MerkleTree) {
	tree.RebuildTreeWith(content)
	//tree.RebuildTree()
}

func GetMerkleRoot(tree *smerkletree.MerkleTree) []byte {
	return tree.MerkleRoot()
}

func VerifySpecificLeaf(tree *smerkletree.MerkleTree, content smerkletree.Content) bool {
	res, err := tree.VerifyContent(content)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func ShowLeafs(tree *smerkletree.MerkleTree) string {

	return tree.String()
}

func VerifyAll(tree *smerkletree.MerkleTree) bool {
	res, err := tree.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func Hashes(tree *smerkletree.MerkleTree) []string {
	return tree.GetHash()
}
func Strings(tree *smerkletree.MerkleTree) []string {
	return tree.Strings()
}
