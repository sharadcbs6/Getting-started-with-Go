/*
			Context - JSONP (json with padding)
			JSONP (JSON with Padding) is a technique for making cross-domain requests from browsers that predate CORS support. I
			It works by wrapping a JSON response in a JavaScript function call.

*/

package main
import(

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default();
	router.GET("/JSONP",func(ctx *gin.Context) {
		data:=map[string]interface{}{
			"foo":"bar",
		}
		ctx.JSONP(200,data);
	})
	router.Run(":8080");

}