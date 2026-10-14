/*
					Context- Goroutines
				Gin uses a sync.Pool to reuse gin.Context objects across requests for performance. 
				Once a handler returns, the gin.Context is returned to the pool 
				and may be assigned to an entirely different request. 
				If a goroutine still holds a reference to the original context at that point, 
				it will read or write fields that now belong to another request.
				 This leads to race conditions, data corruption, or panics.

				Calling c.Copy() creates a snapshot of the context that is safe to use after the handler returns.
				 The copy includes the request,URL, keys, and other read-only data, but is detached from the pool lifecycle.
*/
package main

import (
  "log"
  "time"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/long_async", func(c *gin.Context) {
    // create copy to be used inside the goroutine
    cCp := c.Copy()
    go func() {
      time.Sleep(5 * time.Second)

      log.Println("Done! in path " + cCp.Request.URL.Path)
    }()
  })

  router.GET("/long_sync", func(c *gin.Context) {
    time.Sleep(5 * time.Second)
    log.Println("Done! in path " + c.Request.URL.Path)
  })
  router.Run(":8080")
}