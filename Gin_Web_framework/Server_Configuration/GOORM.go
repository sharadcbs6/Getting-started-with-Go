package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Name  string  `json:"name"`
  Price float64 `json:"price"`
}

func main() {
  dsn := "host=localhost user=app dbname=mydb port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect to database")
  }

  // Auto-migrate the schema
  db.AutoMigrate(&Product{})

  r := gin.Default()

  r.POST("/products", func(c *gin.Context) {
    var p Product
    if err := c.ShouldBindJSON(&p); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    if err := db.Create(&p).Error; err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
      return
    }
    c.JSON(http.StatusCreated, p)
  })

  r.GET("/products/:id", func(c *gin.Context) {
    var p Product
    if err := db.First(&p, c.Param("id")).Error; err != nil {
      c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
      return
    }
    c.JSON(http.StatusOK, p)
  })

  r.Run(":8080")
}