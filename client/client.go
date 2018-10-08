package main

import (
	"github.com/ChinasoftNobody/gochat/client/chat"
	"github.com/ChinasoftNobody/gochat/client/config"
	"github.com/ChinasoftNobody/gochat/client/widgets"
	"github.com/ChinasoftNobody/gochat/common/message"
	"github.com/wonderivan/logger"
	"net"
)

/**
客户端
*/
func main() {
	//连接至服务器
	connectToServer()
	//启动界面客户端
	chatWindow := widgets.SingleWindow()
	chatWindow.RunChart()

}

/**
连接至服务器
*/
func connectToServer() {
	conn, ok := chat.StartGoChat("localhost:8000")
	//defer conn.Close()
	if !ok {
		return
	}
	go sendMessageTest(conn)
	go readMessage(conn)
}

/**
接收信息并打印数据
*/
func readMessage(conn net.Conn) {
	for {
		meta, err := message.ReceiveMessage(conn)
		if err != nil {
			return
		}
		logger.Info("新的消息:", meta.Content)
	}
}

/**
从客户端输入器接收数据，并且封装信息发送个服务器
*/
func sendMessageTest(conn net.Conn) {
	message.SendMessage(conn, "Hello server")
}

//初始化配置信息
func init() {
	//初始化日志配置
	config.InitLog()
}
