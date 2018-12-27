package handlers

import (
	"../assets"
	"../ethercrypto/web3history"
	_ "../responses"
	"../tree/content"
	"../tree/tree"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Proof struct {
	Hash []byte `json:"Hash" example:"qNCllA0uMdgEPSVQBYzD4JESEECY2NyjbJgGjy0NP6c="`
}
type Proofs []Proof

type Info struct {
	Key  []byte `json:"Key" example:"qNCllA0uMdgEPSVQBYzD4JESEECY2NyjbJgGjy0NP6c="`
	Hash []byte `json:"Hash" example:"qNCllA0uMdgEPSVQBYzD4JESEECY2NyjbJgGjy0NP6c="`
	Root []byte `json:"Root" example:"qNCllA0uMdgEPSVQBYzD4JESEECY2NyjbJgGjfwNP6c="`
}

type TotalValues struct {
	Data Proofs
	Info Info
}

// @Summary Add new asset to assetId (if assetId exists)
// @Description add hash by assetId
// @Accept  text/plain
// @Produce  application/json
// @Param   assetId     path    string     true        "assetId"
// @Param   hash        path    string     true        "Hash of file"
// @Success 200 {array} responses.UpdateResponse
// @Security BasicAuth
// @Router /a/update/{assetId}/{hash} [post]
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

// @Summary Add a new asset to assetId
// @Description add hash by assetId
// @Accept  text/plain
// @Produce  application/json
// @Param   assetId     path    string     true        "assetId"
// @Param   hash        path    string     true        "Hash of file"
// @Security BasicAuth
// @Success 200 {array} responses.CreateResponse
// @Failure 400 {array} responses.CreateResponseError
// @Router /a/new/{assetId}/{hash} [post]
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
			"merkleRoot": root,
		})
		return
	}
}

// @Summary Give info about all assets and all meta information
// @Description Lists all assets
// @Accept  text/plain
// @Produce  application/json
// @Success 200 {array} responses.ListResponse
// @Router /list [get]
//List Lists all assets in DB
func List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"assets": assets.FindALlAssets(c),
	})
}

// @Summary Get total Merkle proof
// @Description Merkle proof for assetId, txNumber, hash, timestamp (Actually send a JSON File with two arrays Data and Info
// Data is a list of merkle proofs leaves from end to start
// Info has parameters: Key - array key, Hash - array value, Root - current merkle tree Root Hash
// @Accept  text/plain
// @Produce  application/json
// @Param   assetId     path    string     true        "assetId"
// @Param   txNumber    path    string     true        "txNumber"
// @Param   hash        path    string     true         "hash"
// @Param   timestamp   path    string     true        "timestamp"
// @Success 200 {string} string "test it"
// @Router /proof/{assetId}/{txNumber}/{hash}/{timestamp} [get]
//GetTotalProof Get total Merkle proof
func GetTotalProof(c *gin.Context) {
	d, key, data, root := tree.GetProofs(c)
	var proofs = Proofs{}
	for i := 0; i < len(d); i++ {
		proofs = append(proofs,
			Proof{
				d[i],
			})
	}

	var info = Info{}
	info.Key = key
	info.Root = root
	info.Hash = data

	var final = TotalValues{}

	final.Info = info
	final.Data = proofs

	myJson, err := json.Marshal(final)

	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	c.Data(http.StatusOK, "JSON", myJson)

}

// @Summary Return asset by assetId
// @Description Lists all assets by assetId
// @Accept  text/plain
// @Produce  application/json
// @Param   assetId     path    string     true        "assetId"
// @Param   txNumber    path    string     true        "txNumber"
// @Success 200 {array} responses.AssetsResponse
// @Router /get/{assetId}/{txNumber} [get]
//GetData Get timestamp and hash of specified asset in assetId
func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"assets": hex.EncodeToString(assets.GetAssetByAssetIdAndTxNumber(c)),
	})
}
