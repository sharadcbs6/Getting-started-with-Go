package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)
var (
	g errgroup.Group
)
func router01() http.Handler{
	e:=gin.New()
	e.Use(gin.Recovery())
	e.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"code":http.StatusOK,
			"message":"welcome server01",
		})
	})
	return e;
}
func router02() http.Handler{
	e:=gin.New()
	e.Use(gin.Recovery())
	e.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"code":http.StatusOK,
			"message":"welcome server02",
		})
	})
	return e;
}
func main(){
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	g.Go(func()error{
		return server01.ListenAndServe()
	})
	g.Go(func() error{
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}