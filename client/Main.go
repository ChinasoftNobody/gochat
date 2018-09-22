package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChinasoftNobody/gochat/client/chat"
	"github.com/ChinasoftNobody/gochat/client/common"
	"github.com/ChinasoftNobody/gochat/client/dto"
	"net"
)

/**
客户端
*/
func main() {
	//启动聊天客户端
	conn, _ := chat.StartGoChat("localhost:8000")
	defer conn.Close()
	go sendMessage(conn)
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取数据失败:", err)
			return
		}
		fmt.Printf("接收数据：[%s]\n", string(buf[:n]))
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
			fmt.Println("接收系统输入数据失败", err)
			continue
		}
		tmp := string(ipt[:])
		fmt.Println("输入数据：", tmp)
		//封装msg
		msg := dto.CommonMsg{Type: common.MSG_TYPE_STRING, Content: tmp}
		msgBytes := make([]byte, 10240)
		msgBytes, err = json.Marshal(msg)
		if err != nil {
			fmt.Println("json转化消息失败", err)
			continue
		}
		_, err1 := conn.Write(msgBytes)
		if err1 != nil {
			fmt.Println("发送数据失败", err1)
			continue
		}

	}
}
