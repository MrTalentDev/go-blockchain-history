package tree

import (
	"../../../ethercrypto/tree/customsmt"
	"../../../ethercrypto/tree/smt"
	"../content"
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"strconv"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"fmt"
)

func GetRoot(c *gin.Context) []byte {
	t := makeTree(c, content.GetAllContent(c))
	return t.Root()
}

func GetProofs(c *gin.Context) ([][]byte, []byte, []byte, []byte) {
	st := c.Param("txNumber")
	tp := c.Param("timestamp")
	timestamp, err := strconv.Atoi(tp)
	dataHash := content.GenContent(c, int64(timestamp))
	m := make(map[string][]byte)

	txNumber, err := strconv.ParseInt(st, 10, 64)
	if err != nil {
		println(err)
	}
	key := content.CreateKey(c, txNumber)
	m[key] = dataHash
	t := makeTree(c, content.GetAllContent(c))
	d, _ := hex.DecodeString(key)
	proofs, _ := t.Prove(d)
	defer Proof(proofs, t.Root(), m[key], d)
	r := GetRoot(c)
	return proofs, d, dataHash, r
}

func Proof(proofs [][]byte, root []byte, value []byte, key []byte) {
	if smt.VerifyProof(proofs, root, key, value, sha3.NewKeccak256()) {
		fmt.Println("Proof verification succeeded.")
	} else {
		fmt.Println("Proof verification failed.")
	}
}

func makeTree(c *gin.Context, contents map[string][]byte) *smt.SparseMerkleTree {
	t := customsmt.InitTree()
	keys := content.GetAllKeys(c)
	for i := 0; i < len(contents); i++ {
		customsmt.UpdateTree(keys[i], contents[keys[i]], t)
	}
	return t
}
