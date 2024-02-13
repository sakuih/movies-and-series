package middlewares

import (
  "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    apikey := c.GetHeader("xyz-api")
    if apikey == "" {
      c.AbortWithStatusJSON(401, gin.H{"error" : "error use a valid API-key"})
      return
    }
    c.Next()
  }
}
  func RegisterMiddlewares(router *gin.Engine) {
    router.Use(AuthMiddleware())
  }


