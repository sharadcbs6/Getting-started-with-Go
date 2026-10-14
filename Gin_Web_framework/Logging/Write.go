/*
							Context - Logging
			Writing logs to a file is essential for production applications where you need to retain request history for debugging, auditing, or monitoring. 
			By default, Gin writes all log output to os.Stdout. 
			You can redirect this by setting gin.DefaultWriter before creating the router.		
*/
package main
import(
	"io"
	"os"
	"github.com/gin-gonic/gin"
)
func main(){
	gin.DisableConsoleColor()
	f,_:=os.Create("gin.log")
	gin.DefaultWriter=io.MultiWriter(f)
	router:=gin.Default()
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.String(200,"pong")
	})
	router.Run(":8080")
}