/* Map as query string or  postform parameters */
package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
func main(){
	router:=gin.Default();
	router.POST("/post",func(ctx *gin.Context) {
		ids:=ctx.QueryMap("ids");
		names:=ctx.PostFormMap("names");
		fmt.Printf("ids: %v; names: %v\n",ids, names)
		ctx.JSON(http.StatusOK, gin.H{
			"ids": ids,
			"names": names,
		})
	})
	router.Run(":8080");
}