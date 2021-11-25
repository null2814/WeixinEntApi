package handler

import (
	"encoding/xml"
	"fmt"
	"weixinapi/api"
	"weixinapi/utils/wxbizjsonmsgcrypt"

	"github.com/valyala/fasthttp"
)

const (
	token          = "FQ1tdOcVjsGneinxwnFnvp0wtAX"
	receiverId     = "receiverId"
	encodingAeskey = "hi2vvsGmd9GfTfUbsOarsEuGfCs4GUrCc3Jsc8IjpAD"
)

var wxcpt *wxbizjsonmsgcrypt.WXBizMsgCrypt

func init() {
	wxcpt = wxbizjsonmsgcrypt.NewWXBizMsgCrypt(token, encodingAeskey, receiverId, wxbizjsonmsgcrypt.JsonType)
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
	query := ctx.Request.URI().QueryArgs()
	reqMsgSign := string(query.Peek("msg_signature"))
	reqTimestamp := string(query.Peek("timestamp"))
	reqNonce := string(query.Peek("nonce"))
	body := new(api.CallbackRequestBody)
	err := xml.Unmarshal(ctx.Request.Body(), body)
	ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
	if err != nil {
		fmt.Fprint(ctx, "")
		return
	}
	reqDataString := fmt.Sprintf(`{"tousername":"%s","encrypt":"%s","agentid":"%s"}`, body.ToUser, body.MsgEncrypt, body.ToAgentID)
	reqData := []byte(reqDataString)
	msg, cryptErr := wxcpt.DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, reqData)
	if nil != cryptErr {
		fmt.Fprint(ctx, "")
		return
	}
	Msg := new(api.ButtonClickCallbackRequestBody)
	err = xml.Unmarshal(msg, Msg)
	if nil != err {
		fmt.Fprint(ctx, "")
		return
	}
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprint(ctx, Msg.EventKey)
}
