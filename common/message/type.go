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

const Size = 1024*1024*5
