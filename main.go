package main

import (
	"log"

	"weixinapi/handler"

	"github.com/buaazp/fasthttprouter"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func main() {
	r := fasthttprouter.New()
	r.GET("/ping", handler.Ping)
	r.GET("/qywx/callback", handler.CallBackCheckHandler)
	r.POST("/qywx/callback", handler.CallBackHandler)

	logrus.Info("Starting server on port 80")
	log.Fatal(fasthttp.ListenAndServe("0.0.0.0:80", r.Handler))
}
