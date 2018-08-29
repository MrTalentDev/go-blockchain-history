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

	//a, _:= hex.DecodeString("bc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98a")
	//b, _:= hex.DecodeString("f40b9896a9a2e201b669f87a4fd323d0374f41de7ec3911eff7967b3de9b42bd")
	//
	//println(hex.EncodeToString(checkSum(a,b)))

	//Check(proofs, c)
	return proofs, d, dataHash, r
}

func Proof(proofs [][]byte, root []byte, value []byte, key []byte) {
	if smt.VerifyProof(proofs, root, key, value, sha3.NewKeccak256()) {
		fmt.Println("Proof verification succeeded.")
	} else {
		fmt.Println("Proof verification failed.")
	}
}

//type Info struct {
//	Key []byte
//}
//
func checkSum(hashOne []byte, hashTwo []byte)  []byte {
	hasher:= sha3.NewKeccak256()
	hasher.Write(hashOne)
	hasher.Write(hashTwo)
	path := hasher.Sum(nil)
	hasher.Reset()
	return path
}
//
//func Check(proofs [][]byte, c *gin.Context) {
//	d := checkSum(proofs[0], proofs[1])
//
//	var info = Info{}
//	info.Key = d
//
//	myJson, err := json.Marshal(info)
//
//	if err != nil {
//		log.Fatal("Cannot encode to JSON ", err)
//	}
//
//	c.Data(http.StatusOK, "JSON", myJson)
//}

func makeTree(c *gin.Context, contents map[string][]byte) *smt.SparseMerkleTree {
	t := customsmt.InitTree()
	keys := content.GetAllKeys(c)
	for i := 0; i < len(contents); i++ {
		customsmt.UpdateTree(keys[i], contents[keys[i]], t)
	}
	return t
}
