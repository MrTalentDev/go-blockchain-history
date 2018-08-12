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
)

type Proof struct {
	HashProof       string
	FinalMerkleRoot string
}
type Proofs []Proof

// Add new asset to assetId and change merkle tree
func UpdateAssetId(c *gin.Context) {
	assets.UpdateAssetsByAssetId(c)
	tree.RebuildOrCreateTree(c)
	root := tree.GetMerkleRoot(c)
	web3history.SendNewRootHash(root)
	defer assets.IncrementAssetTx(c)
}

// Create new assetId with asset
func CreateAssetId(c *gin.Context) {
	id, er := assets.CheckAndReturn(c)
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

}

// Lists all assets in DB
func List(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"assets": assets.FindALlAssets(c),
	})
}

// Get Merkle proof of specified asset
func GetSpecifiedProof(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": tree.GetSpecificProof(c),
	})
}

// Get total Merkle proof
func GetTotalProof(c *gin.Context) {
	d, root := tree.GetTotalProof(c)
	var proofs = Proofs{}
	for i := 0; i < len(d); i++ {
		var a = d[i]
		proofs = append(proofs,
			Proof{
				a, root,
			})
	}

	myJson, err := json.Marshal(proofs)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	c.Data(http.StatusOK, "JSON", myJson)

}

// Get timestamp and hash of specified asset in assetId
func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"assets": hex.EncodeToString(assets.GetAssetByAssetIdAndTxNumber(c)),
	})
}
