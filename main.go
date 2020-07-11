// go-custom-search
// https://github.com/4OH4/go-custom-search
//
// 4OH4, July 2020
//
// Basic Golang API, with Google Cloud Custom Search API
// Listens on port 3000, creates and endpoint at /search
//
// GET /search?query=search%20terms
// Runs a query for the search terms through Custom Search API
// and returns a list of links ins JSON
//
// Requires a Google Cloud API Key and an active Custom Search
// Engine ID - create one here: https://cse.google.com/cse/all

package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // Autoload environment variables (credentials) from file
)

func main() {

	gin.SetMode(gin.ReleaseMode) // Release mode, comment out for debug

	router := gin.Default()
	router.GET("/search", func(c *gin.Context) {

		query := c.Query("query") // Extract the search query terms

		links := customSearchMain(query) // Run the custom search, get back a slice of related links

		// Just bounce back the query terms as a string
		// c.String(http.StatusOK, "Query: %s", query)

		// Reply to the request with links as JSON:
		c.JSON(200, gin.H{
			"links": links,
		})
	})
	router.Run(":3000")
}
