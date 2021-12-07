package api

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestCallBackRequestUnmarshal(t *testing.T) {
	msg := ``
	fmt.Println(msg)
	cbMsg := new(ButtonClickCallbackRequestBody)
	err := xml.Unmarshal([]byte(msg), cbMsg)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("msg:", cbMsg)
}
