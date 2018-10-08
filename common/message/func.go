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
	if len(bytes) > Size {
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
func ReceiveMessage(conn net.Conn) (meta Meta, err error) {
	byteBuff := make([]byte, Size)
	n, err := conn.Read(byteBuff)
	if err != nil {
		logger.Error("接收数据失败,断开客户端连接", err)
		return
	}
	meta = Meta{}
	err = json.Unmarshal(byteBuff[:n], &meta)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}
