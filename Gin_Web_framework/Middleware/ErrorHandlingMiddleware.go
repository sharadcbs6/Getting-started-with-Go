/*
	Context-> Error Handling Middleware(EHM)

Story behind EHM
In a typical RESTful application, you might encounter errors in any route — invalid input,
database failures, unauthorized access, or internal bugs.
Handling errors individually in each handler leads to repetitive code and inconsistent responses.

Solution
A centralized error-handling middleware solves this by running after each request
and checking for any errors added to the Gin context via c.Error(err).
If errors are found, it sends a structured JSON response with a proper status code.
*/
package main

import (
  "errors"
  "net/http"

  "github.com/gin-gonic/gin"
)


func ErrorHandler() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Next() 

    
    if len(c.Errors) > 0 {
      err := c.Errors.Last().Err

      c.JSON(http.StatusInternalServerError, gin.H{
        "success": false,
        "message": err.Error(),
      })
    }
  }
}

func main() {
  router := gin.Default()
  router.Use(ErrorHandler())

  router.GET("/ok", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "success": true,
      "message": "Everything is fine!",
    })
  })

  router.GET("/error", func(c *gin.Context) {
    c.Error(errors.New("something went wrong"))
  })

  router.Run(":8080")
}