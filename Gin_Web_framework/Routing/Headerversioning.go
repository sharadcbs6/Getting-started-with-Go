package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func VersionMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    version := c.GetHeader("Accept-Version")
    if version == "" {
      version = "v1" 
    }
    c.Set("api_version", version)
    c.Next()
  }
}

func main() {
  r := gin.Default()
  r.Use(VersionMiddleware())

  r.GET("/api/users", func(c *gin.Context) {
    version := c.GetString("api_version")

    switch version {
    case "v2":
      c.JSON(http.StatusOK, gin.H{"version": "v2", "data": []gin.H{}})
    default:
      c.JSON(http.StatusOK, gin.H{"version": "v1", "users": []string{}})
    }
  })

  r.Run(":8080")
}