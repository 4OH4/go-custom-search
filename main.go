package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/search", func(c *gin.Context) {

		query := c.Query("query")

		customSearchMain(query)

		// Just bounce back the query terms as a string
		c.String(http.StatusOK, "Query: %s", query)

		// Or, as JSON:
		// c.JSON(200, gin.H{
		// 		"query": "query",
		// })
	})
	router.Run(":3000")
}
