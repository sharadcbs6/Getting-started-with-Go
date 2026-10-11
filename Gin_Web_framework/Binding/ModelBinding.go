/*    Context

Model Binding and validation

To bind a request body into a type , use model binding.
Supported binding supports JSON, YAML, Standard form values (Foo=bar&boo=baz)

Binding tag should be there on all fileds you want to bind.
For eg: when binding from json, set json:"filedname"

Gin provides two Methods of Binding .
	Type- Must Bind
		Methods-> Bind, BindJSON, Bind XML, BindQuery,
		BindYAML

		Behavior-. these methods use MustBindWith under hood.
		If there is a Binding error , the request is aborted 
		witn c.AbortWithError(400,err).SetType(ErrorBindType).
		This set the response status code to 400 and 
		the content-type header is set to text/plain: character=utf-8.

	Type- Should Bind
		Method - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYaml.
		
		Behaviour - These methods use ShouldBindWith under the  hood. If there is binding error,
		the error is returned and its the  reponsibiliity of developers to handle the request and error 

If a field is decorated with binding: "required" and has a empty value
when binding , an error will be returned

if one of the struct fields is itself a struct (nested struct). the fields of that struct will
also need to be decroated with binding:"required"
to validate correctly

*/


package main
import(
		"net/http"
		"github.com/gin-gonic/gin"
)

//Binding from JSON
type login struct{
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main(){
router:=gin.Default()
	
	router.POST("/loginJSON",func(ctx *gin.Context) {
		var json login
		if err:=ctx.ShouldBindJSON(&json);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error()})
			return
		}
			//sample for json binding ({"user":"sharad", "passowrd":"123"})
		if json.User!="Sharad" && json.Password!="123" {
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"status":"unauthorized access"})
				return
		}
		ctx.JSON(http.StatusOK,gin.H{"status":"You are logged in :>"})
	})


	/*
			Example for binding YAML
			(
				<?xml version="1.0" encoding="UTF-8"?>
				<root>
					<user>Sharad</user>
					<password>123</password>
			
				</root>
			)
	
	
	*/

	router.POST("/loginXML",func(ctx *gin.Context) {
		var xml login
		if err:=ctx.ShouldBindXML(&xml);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error()})
				return
		}
		if xml.User!="Sharad" && xml.Password!="123"{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"status":"Unauthorised access"})
				return
		}
		ctx.JSON(http.StatusOK,gin.H{
			"status":"you are logged in :>"})
		
	})
	//binding html form
	router.POST("/loginForm",func(ctx *gin.Context) {
		var form login
		if err:=ctx.ShouldBind(&form);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		if form.User!="Sharad" && form.Password!="123"{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"status":"Unauthorized",
			})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{
			"status":"you are logged in : >",
		})

	})
	router.Run(":8082")
}