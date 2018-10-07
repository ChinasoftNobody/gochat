package message

/**
消息类型
*/
type Type int

const (
	TypeSendMessage = iota
	TypeSendImage
	TypeSendFile
	TypeServerBroadcast
)
