package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"log"
)

func main() {
	r := gin.Default()

	type Proof struct {
		HashProof string
	}

	type Proofs []Proof

	// getProof(assetID, txNumber) - достает merkleproof
	// http://localhost:8080/getProof/g/g
	r.GET("/getProof/:assetID/:txNumber", func(c *gin.Context) {
		var m []string
		m = append(m, c.Param("assetID"))
		m = append(m, c.Param("txNumber"))

		var d = getMerkleProof()

		var proofs = Proofs{}

		for i := 0; i < len(d); i++ {
			var a = d[i]
			proofs = append(proofs, Proof{
				a,
			})
		}

		myJson, err := json.Marshal(proofs)
		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}

		c.Data(200, "JSON", myJson)
	})

	// get(assetID, txNumber) - достает (timestamp, dataHash)
	// http://localhost:8080/get/1/11
	r.GET("/get/:assetID/:txNumber", func(c *gin.Context) {
		var m []string
		m = append(m, c.Param("assetID"))
		m = append(m, c.Param("txNumber"))

		var d = getTimestampAndDataHash(m)

		c.JSON(200, gin.H{
			"timestamp": d[0],
			"dataHash":  d[1],
		})
	})

	// post(assetID, dataHash) - добавляет данные для данного assetId, автоинкрементит txNumber. Возвращает txNumber.
	// http://localhost:8080/post/1/11
	r.POST("/post/:assetID/:dataHash", func(c *gin.Context) {
		var m []string
		m = append(m, c.Param("assetID"))
		m = append(m, c.Param("dataHash"))

		c.JSON(200, gin.H{
			"txNumber": getTxNumber(),
		})
		if len(m) == 0 {
			c.JSON(200, gin.H{
				"txNumber": " ",
			})
		}

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func getMerkleProof() []string {
	var m []string

	m = append(m, "1")
	m = append(m, "1")
	m = append(m, "1")
	m = append(m, "1")
	m = append(m, "1")
	return m
}

func getTxNumber() string {

	return "1111"
}

func getTimestampAndDataHash(m []string) []string {

	var answer []string
	answer = append(answer, m[0])
	answer = append(answer, m[1])
	return answer
}
