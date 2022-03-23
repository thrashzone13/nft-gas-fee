package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
	"github.com/thrashzone13/nft-gas-fee/service"
)

func main() {
	crp := service.NewCryptonatorService()
	gas := service.NewEthgasService()

	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/price", func(c *gin.Context) {
		gwei := gas.Get()
		eth := crp.Get()
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"eth":    eth.ETH,
			"gwei":   prepareGweiData(gwei),
			"contracts": map[string]map[string]int{
				"art_blocks": {
					"buy_nft": 503542,
				},
				"async_art": {
					"use_control_toke":    71472,
					"setup_control_token": 383225,
					"mint_nft":            710405,
					"bid_on_nft":          183398,
				},
				"ephimera": {
					"create_nft":     298600,
					"accept_nft_bid": 151091,
					"list_nft":       70135,
					"withdarw_order": 58891,
					"make_purchase":  145933,
				},
				"foundation": {
					"list_nft":          314488,
					"mint_nft":          315069,
					"change_list_price": 32745,
					"accept_bid":        264108,
					"create_bid":        70620,
				},
				"makers_place": {
					"create_release_series":     402377,
					"create_release_with_media": 495429,
					"accept_bid":                242427,
					"create_bid":                74214,
				},
				"opensea": {
					"register_account": 389335,
					"buy_nft":          394790,
				},
				"zora": {
					"mint_and_list": 539160,
					"set_ask_price": 81630,
					"accept_price":  285940,
					"remove_bid":    89853,
					"add_bid":       169131,
				},
			},
		})
	})
	r.Use(spa.Middleware("/", "./app/build/"))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func prepareGweiData(gwei *service.EthgasResponse) map[string]int {
	gweiData := make(map[string]int)

	gweiData["low"] = gwei.SafeLow
	gweiData["average"] = gwei.Average
	gweiData["fast"] = gwei.Fast
	gweiData["fastest"] = gwei.Fastest

	return gweiData
}
