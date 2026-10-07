package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)
type albums struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`

}
var album=[]albums{
	{ID:"1", Title:"Believer", Artist: "Imagine Dragons", Price:100.24},
	{ID:"2", Title:"No Fluke", Artist: "Dhanda Nyoliwala", Price:38.25},
	{ID:"3", Title:"Not Guilty", Artist: "Dhanda Nyoliwala", Price:50.25},
}
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK,album);
}
func main(){
	router:= gin.Default();
	router.GET("/albums",getAlbums);
	router.Run("localhost:8080");
}
