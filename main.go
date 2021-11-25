package main

import (
	"log"

	"weixinapi/handler"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	r := fasthttprouter.New()
	r.GET("/qywx/callback", handler.CallBackCheckHandler)
	r.POST("/qywx/callback", handler.CallBackHandler)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
