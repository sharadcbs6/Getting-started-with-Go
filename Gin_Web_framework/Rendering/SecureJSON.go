package main
import(
	
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	router.GET("/someJSON",func(ctx *gin.Context) {
		names:=[]string{"amit","aman","anurag"}
		ctx.SecureJSON(200,names);
	})
	router.Run(":8080")
}