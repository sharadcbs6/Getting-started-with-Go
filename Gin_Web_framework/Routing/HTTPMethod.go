package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"Method":"GET"});
}
func putting(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"Method":"PUT"});
}
func patching(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"Method":"PATCH"});
}
func deleting(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"Method":"DELETE"});
}


func main(){
	router:=gin.Default();
	router.GET("/somereq",getting);
	router.GET("/user/:name",func (c *gin.Context)  {
		name:=c.Param("name")
		c.String(http.StatusOK,"Hello %s",name);
	})
	router.Run("localhost:8082");
}