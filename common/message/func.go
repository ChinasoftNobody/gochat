package message

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"net"
	"time"
)

/**
发送消息
*/
func SendMessage(conn net.Conn, content interface{}) bool {
	if conn == nil {
		logger.Error("conn is nil, can not send message")
		return false
	}
	if content == nil || content == "" {
		logger.Error("message is empty, can not send message")
		return false
	}
	meta := Meta{Type: TypeSendMessage, Content: content, Source: conn.LocalAddr().String(), SendTime: time.Now()}
	bytes, err := json.Marshal(meta)
	if len(bytes) > 1024*1024*5 {
		logger.Error("message too long, can not send message larger than 5M")
	}
	if err != nil {
		logger.Error(err)
		return false
	}

	_, err = conn.Write(bytes)
	if err != nil {
		logger.Error(err)
		return false
	}
	return true
}

/**
接收消息
*/
func ReceiveMessage(bytes []byte) (content interface{}, err error) {
	meta := Meta{}
	err = json.Unmarshal(bytes, &meta)
	if err != nil {
		logger.Error(err)
		return
	}
	content = meta.Content
	return
}
