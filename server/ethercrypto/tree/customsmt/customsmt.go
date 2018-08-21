package customsmt

import (
	"log"
	"../smerkletree"
	"github.com/miguelmota/go-solidity-sha3"
	"bytes"
)

//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	X []byte
}

//CalculateHashBytes hashes the values of a TestContent
func (t TestContent) CalculateHashBytes() ([]byte, error) {
	hash := solsha3.SoliditySHA3(
		solsha3.Bytes32(t.X),
	)
	return hash, nil
}

//Init hashes the values of a TestContent
func (t TestContent) Init() ([]byte) {
	return t.X
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other smerkletree.Content) (bool, error) {
	return bytes.Equal(t.X, other.(TestContent).X), nil
}

//CreateContent make a serialized content for tree
func CreateContent(content [][]byte) []smerkletree.Content {
	var list []smerkletree.Content
	for i := 0; i < len(content); i++ {
		list = append(list, TestContent{X: content[i]})
	}
	return list
}

//CreateTree make a tree and returns a pointer at memory
func CreateTree(list []smerkletree.Content) (*smerkletree.MerkleTree) {
	//Create a new Merkle Tree from the list of Content
	t, err := smerkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

//RewriteTree remake a tree
func RewriteTree(content []smerkletree.Content, tree *smerkletree.MerkleTree) {
	tree.RebuildTreeWith(content)
	//tree.RebuildTree()
}

//GetMerkleRoot returns a merkle root of the tree
func GetMerkleRoot(tree *smerkletree.MerkleTree) []byte {
	return tree.MerkleRoot()
}

//VerifySpecificLeaf verify specific content
func VerifySpecificLeaf(tree *smerkletree.MerkleTree, content smerkletree.Content) bool {
	res, err := tree.VerifyContent(content)
	if err != nil {
		log.Fatal(err)
	}
	return res
}


//VerifyAll proof all the data of the tree
func VerifyAll(tree *smerkletree.MerkleTree) bool {
	res, err := tree.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	return res
}


//Strings show all hashes in string type
func Strings(tree *smerkletree.MerkleTree) []string {
	return tree.Strings()
}
