package serv

import (
	"github.com/ChinasoftNobody/gochat/server/common"
	"github.com/ChinasoftNobody/gochat/server/service"
	"github.com/wonderivan/logger"
	"net"
)

/**
使用tcp创建一个服务器
*/
func StartGoChatServer(url string) {
	listener, err := net.Listen("tcp", url)
	defer listener.Close()
	if err != nil {
		panic("启动服务器失败")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("接收客户端连接失败", err)
			continue
		}
		//启动协程对客户端连接进行服务
		go newClient(conn)
	}
}

/**
新的客户端连接至服务器
*/
func newClient(conn net.Conn) {
	common.NewConn(conn)
	service.NewClientConnect(conn)

}
