package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	router.GET("/twitter", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://x.com")
	  })
	router.POST("/submit", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/result")
	  })
	  router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/final"
		router.HandleContext(c)
	  })
	  router.GET("/final",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"hello":"world",
		})
		})
	  router.GET("/result",func(ctx *gin.Context) {
		ctx.String(http.StatusOK,"Redirected here!")
	  })
	  router.Run(":8080")

	
}