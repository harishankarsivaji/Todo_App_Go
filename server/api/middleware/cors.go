package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS Middleware
func CORS(c *gin.Context) {

	//Add the headers with need to enable CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Handle OPTIONS method
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		c.AbortWithStatus(http.StatusOK)
	}
}
