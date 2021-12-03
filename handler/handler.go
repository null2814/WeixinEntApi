package handler

import (
	"encoding/xml"
	"fmt"
	"time"
	"weixinapi/api"
	"weixinapi/utils/wxbizjsonmsgcrypt"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

const (
	token          = ""
	receiverId     = ""
	encodingAeskey = ""
	toUsers        = ""
)

var wxcpt *wxbizjsonmsgcrypt.WXBizMsgCrypt

func init() {
	wxcpt = wxbizjsonmsgcrypt.NewWXBizMsgCrypt(token, encodingAeskey, receiverId, wxbizjsonmsgcrypt.JsonType)
	logrus.SetLevel(logrus.TraceLevel)
}

func CallBackCheckHandler(ctx *fasthttp.RequestCtx) {
	query := ctx.Request.URI().QueryArgs()
	verifyMsgSign := string(query.Peek("msg_signature"))
	verifyTimestamp := string(query.Peek("timestamp"))
	verifyNonce := string(query.Peek("nonce"))
	verifyEchoStr := string(query.Peek("echostr"))
	echoStr, cryptErr := wxcpt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	if nil != cryptErr {
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprint(ctx, "")
		return
	}
	fmt.Fprint(ctx, string(echoStr))
}

func CallBackHandler(ctx *fasthttp.RequestCtx) {
	logrus.Info("CallBack has been called")
	query := ctx.Request.URI().QueryArgs()
	reqMsgSign := string(query.Peek("msg_signature"))
	reqTimestamp := string(query.Peek("timestamp"))
	reqNonce := string(query.Peek("nonce"))
	body := new(api.CallbackRequestBody)
	err := xml.Unmarshal(ctx.Request.Body(), body)
	ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
	if err != nil {
		logrus.Error("body unmarshal failed  with:", err.Error())
		fmt.Fprint(ctx, "")
		return
	}
	reqDataString := fmt.Sprintf(`{"tousername":"%s","encrypt":"%s","agentid":"%s"}`, body.ToUser, body.MsgEncrypt, body.ToAgentID)
	reqData := []byte(reqDataString)
	msg, cryptErr := wxcpt.DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, reqData)
	if nil != cryptErr {
		logrus.Error("Decrypt failed  with:", cryptErr.ErrMsg)
		fmt.Fprint(ctx, "")
		return
	}
	Msg := new(api.ButtonClickCallbackRequestBody)
	err = xml.Unmarshal(msg, Msg)
	if nil != err {
		logrus.Error("msg unmarshal failed  with:", err.Error())
		fmt.Fprint(ctx, "")
		return
	}
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	logrus.Info("unmarshal successed with:", Msg.EventKey)
	timeStamp := time.Now().Unix()
	ButtonChangeMsg := MakeButtonStatusChangeResponse(toUsers, receiverId, timeStamp)
	respStr, err := MakeResponseData(ctx, ButtonChangeMsg, timeStamp)
	if err != nil {
		fmt.Fprint(ctx, "")
		return
	}
	logrus.Info("resp with:", respStr)
	fmt.Fprint(ctx, respStr)
}

func Ping(ctx *fasthttp.RequestCtx) {
	logrus.Info("Ping has been called")
	query := ctx.Request.URI().QueryArgs()
	msg := string(query.Peek("ping"))
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	if msg != "" {
		fmt.Fprint(ctx, fmt.Sprintf("Got msg:%v", msg))
		return
	}
	fmt.Fprint(ctx, "pong!")
}
