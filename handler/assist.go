package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

func MakeButtonStatusChangeResponse(toUsers, receiverId string, timeStamp int64) string {
	buttonChangeMsg := fmt.Sprintf(`<xml><ToUserName><![CDATA[%v]]></ToUserName><FromUserName><![CDATA[%v]]></FromUserName><CreateTime>%d</CreateTime><MsgType><![CDATA[update_button]]></MsgType><Button><ReplaceName><![CDATA[审批驳回]]></ReplaceName></Button></xml>`, toUsers, receiverId, timeStamp)
	return buttonChangeMsg
}

func MakeResponseData(ctx *fasthttp.RequestCtx, msg string, timestamp int64) (string, error) {
	reqNonce := fmt.Sprintf("Nonce_%v", wxcpt.GetRandString(16))
	timeStamp := strconv.Itoa(int(timestamp))
	encrypt, sig, err := wxcpt.GetEncryptMsg(msg, timeStamp, reqNonce)
	if err != nil {
		return "", errors.New(err.ErrMsg)
	}
	respBody := fmt.Sprintf(`<xml><Encrypt><![CDATA[%v]]></Encrypt><MsgSignature><![CDATA[%v]]></MsgSignature><TimeStamp>%v</TimeStamp><Nonce><![CDATA[%v]]></Nonce></xml>`, encrypt, sig, timeStamp, reqNonce)
	return respBody, nil
}
