/*
			context
			Gin’s built-in gin.Recovery() middleware catches any panic that happens while handling a request,
			 writes a 500 response, and keeps the server running. 
			When you need to control what happens on recovery – for example to report the panic to the user, 
			save it to a database, or send it to an error-tracking service – use gin.CustomRecovery() instead.

*/
package main
import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
func main() {
	// Creates a router without any middleware by default
	router := gin.New()
  
	
	router.Use(gin.Logger())
  
	
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
	  if err, ok := recovered.(string); ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	  }
	  c.AbortWithStatus(http.StatusInternalServerError)
	}))
  
	router.GET("/panic", func(c *gin.Context) {
	
	  panic("foo")
	})
  
	router.GET("/", func(c *gin.Context) {
	  c.String(http.StatusOK, "ohai")
	})
  
	
	router.Run(":8080")
  }