/*
				Context-basicAuthMiddleware
				Gin ships with a built-in gin.BasicAuth() middleware that implements HTTP Basic Authentication. 
				It accepts a gin.Accounts map (a shortcut for map[string]string) of username/password pairs ,
				and protects any route group it is applied to.
*/

package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)


var secrets = gin.H{
  "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
  "austin": gin.H{"email": "austin@example.com", "phone": "666"},
  "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
  router := gin.Default()

  authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
    "foo":    "bar",
    "austin": "1234",
    "lena":   "hello2",
    "manu":   "4321",
  }))

  authorized.GET("/secrets", func(c *gin.Context) {
    
    user := c.MustGet(gin.AuthUserKey).(string)
    if secret, ok := secrets[user]; ok {
      c.JSON(http.StatusOK, gin.H{" user ": user, "secret": secret})
    } else {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
    }
  })

  router.Run(":8080")
}