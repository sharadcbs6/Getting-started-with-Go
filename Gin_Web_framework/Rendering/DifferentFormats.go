package main
import(
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/testdata/protoexample"
)
func main(){
	router:=gin.Default()//middleware include logger, recovery
	router.GET("/someJSON",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message": http.StatusOK,
		})
	
	})
	router.GET("/moreJSON",func(ctx *gin.Context) {
		var msg struct{
			Author string `json:"user"`
			Message string
			Number int
		}
		msg.Author="Emiliano Zapata"
		msg.Message="Better to die standing than to live on your knees"
		msg.Number=120
		ctx.JSON(200,msg)
	})
	router.GET("/someXML",func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	router.GET("/someYAML",func(ctx *gin.Context) {
		ctx.YAML(http.StatusOK,gin.H{"message":"hey","status":http.StatusOK})
	})
	router.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		
		data := &protoexample.Test{
		  Label: &label,
		  Reps:  reps,
		}
		
		c.ProtoBuf(http.StatusOK, data)
	  })
	
	 
	  router.Run(":8080")
}