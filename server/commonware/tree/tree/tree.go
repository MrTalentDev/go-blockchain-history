package tree

import (
	"../../../ethercrypto/tree/customsmt"
	"../../../ethercrypto/tree/smt"
	"../content"
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"fmt"
	"strconv"
)

func GetRoot(c *gin.Context) []byte {
	t := makeTree(c, content.GetAllContent(c))
	return t.Root()
}

func GetProofs(c *gin.Context) ([][]byte, []byte, []byte, []byte) {
	st := c.Param("txNumber")
	hash := c.Param("hash")
	dataHash, _ := hex.DecodeString(hash)
	m := make(map[string][]byte)
	txNumber, err := strconv.ParseInt(st, 10, 64)
	if err != nil {
		println(err)
	}
	key := content.CreateKey(c, txNumber)
	m[key] = dataHash
	t := makeTree(c, m)
	d, _ := hex.DecodeString(key)
	proofs, _ := t.Prove(d)
	if smt.VerifyProof(proofs, t.Root(), d, m[key], sha3.NewKeccak256()) {
		fmt.Println("Proof verification succeeded.")
	} else {
		fmt.Println("Proof verification failed.")
	}
	r := GetRoot(c)
	return proofs, d, dataHash, r
}

func makeTree(c *gin.Context, contents map[string][]byte) *smt.SparseMerkleTree {
	t := customsmt.InitTree()
	keys := content.GetAllKeys(c)
	for i := 0; i < len(contents); i++ {
		customsmt.UpdateTree(keys[i], contents[keys[i]], t)
	}
	return t
}
