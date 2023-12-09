package api

import (
	"fmt"
	"gasPrices/gasPrice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")
	api.GET("/getGasPrice", getGasPrice)

	return r
}

func Run() {
	r := setupRouter()
	r.Run(":8080")
}

func getGasPrice(c *gin.Context) {
	gas, err := gasPrice.GetCurrentGasPrice()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"value": nil,
			"ok":    false,
			"error": "failed to fetch gas price",
		})
	}

	formatedGas := gasPrice.FormatGasPrice(gas)

	c.JSON(http.StatusOK, gin.H{
		"value": fmt.Sprintf("%.2f", formatedGas),
		"ok":    true,
		"error": nil,
	})
}
