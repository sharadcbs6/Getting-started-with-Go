package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

)
const(
	MaxUploadSize=1<<20;//1MB
)
func uploadedHandler(ctx *gin.Context){
	ctx.Request.Body=http.MaxBytesReader(ctx.Writer,ctx.Request.Body,MaxUploadSize);
	if err:= ctx.Request.ParseMultipartForm(MaxUploadSize);err!=nil{
		if _, ok:=err.(*http.MaxBytesError);ok{
			ctx.JSON(http.StatusRequestEntityTooLarge,gin.H{
				"error":fmt.Sprintf("file too large (max: %d bytes)",MaxUploadSize),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()});
		return
	}
	file,_,err:= ctx.Request.FormFile("file")
	if err!= nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"file form required"})
		return
	}
	defer file.Close()
	ctx.JSON(http.StatusOK,gin.H{
		"message":"upload successful",
	})
}
func main(){
	router:=gin.Default();
	router.POST("/upload",uploadedHandler)
	router.Run(":8080")
}