package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/api/data", func(c *gin.Context) {
    ctx := c.Request.Context()

    log.Println("request context deadline:", ctx.Done())

    c.JSON(http.StatusOK, gin.H{"status": "ok"})
  })

  r.Run(":8080")
}