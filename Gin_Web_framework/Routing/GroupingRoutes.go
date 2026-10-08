package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func logicEndpoint(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{"action":"login"});

}
func submitEndpoint(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"action":"submit",
	});
}

func readEndpoint(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"action":"read",
	});
}
func main(){
	router:=gin.Default();
	{
		v1:=router.Group("/v1")
		v1.POST("/login",logicEndpoint)
		v1.POST("/submit",submitEndpoint)
		v1.POST("/read",readEndpoint)

	}
	{
		v2:=router.Group("/v2")
		v2.POST("/login",logicEndpoint)
		v2.POST("/submit",submitEndpoint)
		v2.POST("/read",readEndpoint)
	}
	router.Run(":8080");
}
