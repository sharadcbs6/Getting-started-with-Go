/*
			Context - Custom autocert Manager
	For more control — such as specifying a certificate cache directory or restricting allowed hostnames — use autotls.RunWithManager with a custom autocert.Manager:

*/

package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main(){

router:=gin.Default()
router.GET("/ping",func (ctx *gin.Context)  {
	ctx.String(200,"pong")
})
m:=autocert.Manager{
	Prompt: autocert.AcceptTOS,
	HostPolicy: autocert.HostWhitelist("example1.com","example2.com"),
	Cache: autocert.DirCache("/var/www/.cache"),
}
log.Fatal(autotls.RunWithManager(router,&m))

}