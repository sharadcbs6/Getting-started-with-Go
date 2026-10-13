/*
			Context-customMiddlewareHandler.go
		Gin Middleware  is a fxn that returns a gin.handlerFunc .
		Middleware runs before/after the main handler, which makes it more powerful for logging,
		authenication,error handling and cross cutting platform

		Middleware execution flow
		A middleware function has two phases, divided by the call to c.Next():

		Before c.Next() – Code here runs before the request reaches the main handler. 
			Use this phase for setup tasks such as recording the start time, validating tokens, 
			or setting context values with c.Set().
		c.Next() – This calls the next handler in the chain (which may be another middleware 
			or the final route handler).
			Execution pauses here until all downstream handlers have completed.
		After c.Next() – Code here runs after the main handler has finished. 
		Use this phase for cleanup, logging response status, or measuring latency.
		If you want to stop the chain entirely (for example, when authentication fails), 
		call c.Abort() instead of c.Next(). This prevents any remaining handlers from executing. 
		
*/

package main

import (
  "log"
  "time"

  "github.com/gin-gonic/gin"
)
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
	  t := time.Now()
	  c.Set("example", "12345")
	  c.Next()
	  latency := time.Since(t)
	  log.Print(latency)  
	  status := c.Writer.Status()
	  log.Println(status)
	}
  }
  
func main(){
	router := gin.New()
	router.Use(Logger())
  
	router.GET("/test", func(c *gin.Context) {
	  example := c.MustGet("example").(string)
	  log.Println(example)
	})
	router.Run(":8080")
  }
