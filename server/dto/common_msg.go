package dto

/**
通用消息体，由元消息中content转化而来
*/
type CommonMsg struct {
	//消息类型
	Type       int    `json:"type"`
	Close      bool   `json:"close"`
	Content    string `json:"content"`
	TargetAddr string `json:"target_addr"`
}
