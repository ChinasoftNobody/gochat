package chat

import (
	"github.com/wonderivan/logger"
	"net"
)

func StartGoChat(serverPath string) (conn net.Conn, ok bool) {
	conn, err := net.Dial("tcp", serverPath)
	if err != nil {
		logger.Error("连接服务器失败", err)
		ok = false
		return
	}
	logger.Info("已连接至服务器")
	ok = true
	return
}
