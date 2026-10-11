package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct{
	Name string `form:"name,default=William"`
	Age int		`form:"age,default=10"`
	Friends []string `form:"friends,default=Will;Bill"`
	Address [2]string `form:"addresses,default=foo bar" collection_format:"ssv"`
	LapTimes []int    `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

func main(){
	router:=gin.Default()
	router.POST("/person",func (ctx *gin.Context)  {
			var person Person
			if err:=ctx.ShouldBind(&person); err!=nil{
				ctx.JSON(http.StatusBadGateway,gin.H{
					"error":err.Error(),
				})
				return 
			}
			ctx.JSON(http.StatusOK,person)
	})
	router.Run(":8080");
}