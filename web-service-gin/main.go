package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)
type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`

}
var albums=[]album{
	{ID:"1", Title:"Believer", Artist: "Imagine Dragons", Price:100.24},
	{ID:"2", Title:"No Fluke", Artist: "Dhanda Nyoliwala", Price:38.25},
	{ID:"3", Title:"Not Guilty", Artist: "Dhanda Nyoliwala", Price:50.25},
}


func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK,albums);
}
func postAlbums(c *gin.Context){
	var newAlbums album;
	if err:=c.BindJSON(&newAlbums);err!=nil{
			return;
	}
	albums=append(albums,newAlbums);
	c.IndentedJSON(http.StatusCreated,newAlbums);
}
func main(){
	router:= gin.Default();
	router.GET("/albums",getAlbums);
	router.POST("/addalbums",postAlbums);
	router.Run("localhost:8080");

}
