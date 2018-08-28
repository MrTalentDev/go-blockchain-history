package handlers

import (
	"github.com/gin-gonic/gin"
	"../assets"
	"../tree/tree"
	"../tree/content"
	"../../ethercrypto/web3history"
	"net/http"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"log"
)

type Proof struct {
	//Number string
	Hash string
}
type Proofs []Proof

type Info struct {
	Key  string
	Hash string
	Root string
}

type TotalValues struct {
	Data Proofs
	Info Info
}

//UpdateAssetId Add new asset to assetId and change merkle tree
func UpdateAssetId(c *gin.Context) {
	if assets.Check(c) {
		timing := assets.UpdateAssetsByAssetId(c)
		tx := assets.GetTxNumber(c)
		tx++
		content.AddContent(c, tx, timing)
		root := tree.GetRoot(c)
		web3history.SendNewRootHash(root)
		defer assets.IncrementAssetTx(c, timing)
	} else {
	}
}

//CreateAssetId Create new assetId with asset
func CreateAssetId(c *gin.Context) {
	id, er, try := assets.CheckAndReturn(c)
	if try {
		tx := assets.GetTxNumber(c)
		timestamp, err := strconv.Atoi(id[2])
		if err != nil {
			panic(err)
		}
		content.AddContent(c, tx, int64(timestamp))
		root := tree.GetRoot(c)
		web3history.SendNewRootHash(root)
		if er == "err" {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"assetId":    id[0],
			"txNumber":   tx,
			"timstamp":   id[2],
			"hash":       id[1],
			"merkleRoot": hex.EncodeToString(root),
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
	d, key, data, root := tree.GetProofs(c)
	var proofs = Proofs{}
	for i := 0; i < len(d); i++ {
		proofs = append(proofs,
			Proof{
				"0x" + hex.EncodeToString(d[i]),
			})
	}

	var info = Info{}
	info.Key = "0x" + hex.EncodeToString(key)
	info.Root = "0x" + hex.EncodeToString(root)
	info.Hash = "0x" + hex.EncodeToString(data)

	var final = TotalValues{}

	final.Info = info
	final.Data = proofs

	myJson, err := json.Marshal(final)

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
