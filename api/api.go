package api

type CallbackRequestBody struct {
	ToUser     string `xml:"ToUserName"`
	ToAgentID  string `xml:"AgentID"`
	MsgEncrypt string `xml:"Encrypt"`
}

type (
	ButtonClickCallbackRequestBody struct {
		ToUserName    string                                       `xml:"ToUserName"`   // 企业微信CorpID
		FromUserName  string                                       `xml:"FromUserName"` // 成员UserID
		CreateTime    int                                          `xml:"CreateTime"`   // 消息创建时间（整型）
		MsgType       string                                       `xml:"MsgType"`      // 消息类型 event
		Event         string                                       `xml:"Event"`        // 事件类型 template_card_event
		EventKey      string                                       `xml:"EventKey"`     // 按钮btn:key值
		TaskId        string                                       `xml:"TaskId"`       // task_id
		CardType      string                                       `xml:"CardType"`     // 通用模板卡片的类型
		ResponseCode  string                                       `xml:"ResponseCode"` // 用于调用更新卡片接口的ResponseCode
		AgentID       int                                          `xml:"AgentID"`      // 企业应用的id，整型
		SelectedItems []ButtonClickCallbackRequestBodySelectedItem `xml:"SelectedItems"`
	}
	ButtonClickCallbackRequestBodySelectedItem struct {
		QuestionKey string                                               `xml:"QuestionKey"` // 问题的key值
		OptionIds   []ButtonClickCallbackRequestBodySelectedItemOptionId `xml:"OptionIds"`   // 对应问题的选项列表
	}
	ButtonClickCallbackRequestBodySelectedItemOptionId string
)
