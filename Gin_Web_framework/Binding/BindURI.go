/*
			Context - Bind URI
		ShouldBindURI binds URI path parameters directly into  a struct using uri struct tags.
		this combined with binding validations tags , helps to validate path parameters, such as
		uuid with a single call

		


*/

package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type People struct{
	ID string 	`uri:"id" binding:"required,uuid"`
	Name string  `uri:"name" binding:"required"`
}

func main(){
	router:=gin.Default()
	router.GET("/:name/:id",func(ctx *gin.Context) {
		var per People
		if err:=ctx.ShouldBindUri(&per);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{
			"name":per.Name,
			"uuid":per.ID,
		})
	})
	router.Run(":8081");
}