package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChinasoftNobody/gochat/client/chat"
	"github.com/ChinasoftNobody/gochat/client/common"
	"github.com/ChinasoftNobody/gochat/client/config"
	"github.com/ChinasoftNobody/gochat/client/dto"
	"github.com/ChinasoftNobody/gochat/client/widgets"
	"github.com/wonderivan/logger"
	"log"
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
	go sendMessage(conn)
	go readMessage(conn)
}

/**
接收信息并打印数据
*/
func readMessage(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			logger.Error("读取数据失败:", err)
			return
		}
		logger.Info("接收数据：[%s]\n", string(buf[:n]))
	}
}

/**
从客户端输入器接收数据，并且封装信息发送个服务器
*/
func sendMessage(conn net.Conn) {
	for {
		var ipt = make([]byte, 1024)
		_, err := fmt.Scan(&ipt)
		if err != nil {
			logger.Error("接收系统输入数据失败", err)
			continue
		}
		tmp := string(ipt[:])
		log.Println("输入数据：", tmp)
		//封装msg
		msg := dto.CommonMsg{Type: common.MsgTypeString, Content: tmp}
		msgBytes := make([]byte, 10240)
		msgBytes, err = json.Marshal(msg)
		if err != nil {
			logger.Error("json转化消息失败")
			continue
		}
		_, err1 := conn.Write(msgBytes)
		if err1 != nil {
			logger.Error("发送数据失败")
			continue
		}

	}
}

//初始化配置信息
func init() {
	//初始化日志配置
	config.InitLog()
}
