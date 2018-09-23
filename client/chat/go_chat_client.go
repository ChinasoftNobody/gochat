package chat

import (
	"fmt"
	"net"
)

func StartGoChat(serverPath string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", serverPath)
	if err != nil {
		fmt.Println("连接服务器失败", err)
	}
	fmt.Println("已连接至服务器")
	return
}
