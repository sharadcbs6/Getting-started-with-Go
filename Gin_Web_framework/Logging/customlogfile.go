package main
import(
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)
// LoggerWithFormatter middleware  will write the logs to gin.DefaultWriter
//By default gin.DefaultWriter=os.Stdout
func main(){
	router:=gin.New()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		params.ClientIP,
		params.TimeStamp.Format(time.RFC1123),
		params.Method,
        params.Path,
        params.Request.Proto,
        params.StatusCode,
        params.Latency,
        params.Request.UserAgent(),
        params.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
	  c.String(200, "pong")
	})
	router.Run(":8080")
}