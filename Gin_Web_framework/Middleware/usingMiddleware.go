package main

import (
	"github.com/gin-gonic/gin"
)


func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}


func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "login endpoint"})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "submit endpoint"})
}

func readEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "read endpoint"})
}

func analyticsEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "analytics endpoint"})
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authorized := router.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	router.Run(":8080")
}