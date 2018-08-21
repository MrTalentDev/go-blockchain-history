package handlers

import (
	"github.com/gin-gonic/gin"
	"../assets"
	"../tree"
	"../../ethercrypto/web3history"
	"net/http"
	"encoding/hex"
	"log"
	"encoding/json"
	"strconv"
)

type Proof struct {
	Number    string
	Hash      string
	MerkleRoot string
}
type Proofs []Proof

//UpdateAssetId Add new asset to assetId and change merkle tree
func UpdateAssetId(c *gin.Context) {
	if assets.Check(c) {
		assets.UpdateAssetsByAssetId(c)
		tree.RebuildOrCreateTree(c)
		root := tree.GetMerkleRoot(c)
		web3history.SendNewRootHash(root)
		defer assets.IncrementAssetTx(c)
	} else {

	}

}

//CreateAssetId Create new assetId with asset
func CreateAssetId(c *gin.Context) {
	id, er, try := assets.CheckAndReturn(c)
	if try {
		tree.RebuildOrCreateTree(c)
		root := tree.GetMerkleRoot(c)
		web3history.SendNewRootHash(root)
		if er == "err" {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"assetId": id[0],
			"hash":    id[1],
		})
		return
	}
}

//List Lists all assets in DB
func List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"assets": assets.FindALlAssets(c),
	})
}

//GetTotalProof Get total Merkle proof
func GetTotalProof(c *gin.Context) {
	d, root := tree.GetTotalProof(c)
	var proofs = Proofs{}
	var i int
	for i = 0; i < len(d); i++ {
		var a = d[i]
		proofs = append(proofs,
			Proof{
				strconv.FormatInt(int64(i), 10),a, root,
			})
	}

	myJson, err := json.Marshal(proofs)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	c.Data(http.StatusOK, "JSON", myJson)

}

//GetData Get timestamp and hash of specified asset in assetId
func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"assets": hex.EncodeToString(assets.GetAssetByAssetIdAndTxNumber(c)),
	})
}
