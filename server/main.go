//这只是一个文档信息库
package main

import (
	"github.com/ChinasoftNobody/gochat/common/kafka"
	"github.com/ChinasoftNobody/gochat/server/config"
	"github.com/ChinasoftNobody/gochat/server/serv"
)

/**
服务端，实现用户登陆，用户聊天，用户下线
*/
func main() {
	//启动服务器监听，监听固定端口8000
	serv.StartGoChatServer(":8000")

}

func init() {
	config.InitLog()
	kafka.Test()

}
