/*				context
				ShouldBindQuery
ShouldBindQuery bind only the URL query string parameters to a struct ,ignoring the request body.
This is useful when you want to ensure that post data don't overwrite the query parameters.


*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct{
	Name string `form:"name"`
	Address string `form:"address"`
}

func startPage(ctx *gin.Context){
	var person Person
	if err:=ctx.ShouldBindQuery(&person);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error()})
			return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"name":person.Name,
		"addaress":person.Address,
	})
}

func main(){
	router:=gin.Default()
	router.Any("/testing",startPage)
	router.Run(":8081")
}
