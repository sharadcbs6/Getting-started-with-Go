/*
			Context - Bind custom unmarshaler
	To override Gin’s default binding logic, define a function on your type that satisfies the encoding.
	TextUnmarshaler interface from the Go standard library.
	 Then specify parser=encoding.TextUnmarshaler in the uri/form tag of the field being bound.
*/

package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Birthday string

func (b *Birthday) UnmarshalText(text []byte) error {
  *b = Birthday(strings.Replace(string(text), "-", "/", -1))
  return nil
}

var _ encoding.TextUnmarshaler = (*Birthday)(nil)
func main() {
	route := gin.Default()
	var request struct {
	  Birthday         Birthday   `form:"birthday"`
	  Birthdays        []Birthday `form:"birthdays" collection_format:"csv"`
	  BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02" collection_format:"csv"`
	}
	route.GET("/test", func(ctx *gin.Context) {
	  err := ctx.BindQuery(&request);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"err":err.Error()})
		return;
	  }
	  ctx.JSON(200, request)
	})
	route.Run(":8088")
  }