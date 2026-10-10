
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  v1 := r.Group("/api/v1")
  {
    v1.GET("/users", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{"version": "v1", "users": []string{}})
    })
  }

  v2 := r.Group("/api/v2")
  {
    v2.GET("/users", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "version": "v2",
        "data":    []gin.H{},
        "meta":    gin.H{"total": 0},
      })
    })
  }

  r.Run(":8080")
}