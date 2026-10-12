/* 
			Context-AsciiJSON
			AsciiJSON converts data to JSON but escapes all non-ASCII charater into \uXXXX unicode escape characters.
			HTML-special characters like < and > are also escaped.

*/
package main
import(
  	"github.com/gin-gonic/gin"
)
func main(){
	router:=gin.Default()
	router.GET("/someJSON",func(ctx *gin.Context) {
		data:=map[string]interface{}{
			"lang":"Golang",
			"tag":"<br>",
		}
		ctx.AsciiJSON(200,data);
	})
	router.Run(":8080");
}