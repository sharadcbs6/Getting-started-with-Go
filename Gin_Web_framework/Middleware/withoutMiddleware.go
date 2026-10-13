/*
						Context-Middleware
			gin offers two ways to create a route engine.
			1. gin.Default()-> comes with logger and recovery
				recovery-> prevents from server crashes,recovers from any panic in handlers
				and returns 500 reponse.
				logger->write logs to stdout(method,path,statuscode.)
			2. gin.New()-> a blank engine, a bare  router with no middleware attached,
				can customize panic behaviour
				while building microservice ,need  a minimal or special middleware stack.
				
*/
package main
import (
	"log"	
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default();
	router.Use(gin.Recovery()) // this middleware prevents from server crashes
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.String(200,"pong");

	})
	log.Fatal(router.Run(":8080"));
}