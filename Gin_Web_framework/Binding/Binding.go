/*
							Context of Binding
			Gin provides a powerful system that parses the request data into struct and validates it . 
			Instead of calling manually ctx.Postform() or reading ctx.Request.Body,need to define a
			struct with tags and gin do the rest work.

*/
package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type Form struct{
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main(){
		router:=gin.Default()

		router.POST("/loginform",func(ctx *gin.Context) {
			var form Form
			if err:=ctx.ShouldBind(&form);err!=nil{
				ctx.JSON(http.StatusBadRequest,gin.H{
					"error":err.Error()})
					return
			}
			if form.User!="Sharad"&&form.Password!="123"{ 
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"status":"Not Authorized"})
				return
			}
			ctx.JSON(http.StatusOK,gin.H{
				"message":"Successfully logged in",
				"user":form.User,
				
			})
		})
		router.Run(":8080")
}