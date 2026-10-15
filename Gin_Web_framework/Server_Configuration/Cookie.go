/*
	SetCookie parameters
The c.SetCookie() method signature is:

c.SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)


*/

package main
import(
	"fmt"
	"github.com/gin-gonic/gin"
)
func main()  {
		router:=gin.Default()
		router.GET("/cookie",func(ctx *gin.Context) {
			cookie,err:=ctx.Cookie("gin_cookie")
			if err!=nil{
				cookie="Not Set"
				ctx.SetCookie("gin_cookie","test",3600,"/","locahost",true,true)
			}
			fmt.Printf("Cookie values:%s\n",cookie)
		})
		router.Run(":8080")

}