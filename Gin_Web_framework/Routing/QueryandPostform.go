package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	
	ID   string `form:"id"`
	Page string `form:"page"`

	
	Name    string `form:"name"`
	Message string `form:"message"`
}

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		var req PostRequest

		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", req.ID, req.Page, req.Name, req.Message)

		c.String(
			http.StatusOK,
			"id: %s; page: %s; name: %s; message: %s",
			req.ID, req.Page, req.Name, req.Message,
		)
	})

	router.Run(":8080")
}