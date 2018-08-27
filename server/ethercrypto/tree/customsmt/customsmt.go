package customsmt

import (
	"../smt"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"encoding/hex"
)

func InitTree() (*smt.SparseMerkleTree) {
	return smt.NewSparseMerkleTree(smt.NewSimpleMap(), sha3.NewKeccak256())
}

func UpdateTree(key string, data []byte, currentTree *smt.SparseMerkleTree) {
	d, _ := hex.DecodeString(key)
	currentTree.Update(d, data)
}
