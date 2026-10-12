package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type Filters struct{
	Tags []string `forms:"tags" collection_format:"csv"`
	Labels []string `forms:"labels" collection_format:"multi"`
	IdSSV  []int	`forms:"id_ssv" collection_format:"ssv"`
	IdsTSV  []int  	`forms:"id_tsv" collection_format:"tsv"`
	Levels  []int 	`forms:"levels" collection_format:"pipes"`
}
func main(){
	router:=gin.Default();
	router.GET("/search",func(ctx *gin.Context){
		var filter Filters
		if err:=ctx.ShouldBind(&filter);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK,filter);
	})
	router.Run(":8080")
}