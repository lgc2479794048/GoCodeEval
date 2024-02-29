package response

import "github.com/gin-gonic/gin"

// Error sends an error response with the provided message.
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

// Success sends a success response with the provided data.
func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
