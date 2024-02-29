package middleware

import "github.com/gin-gonic/gin"

// ExampleMiddleware is an example of how you might set up a middleware.
func ExampleMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Middleware logic before controller handler...
        c.Next()
        // Middleware logic after controller handler...
    }
}
