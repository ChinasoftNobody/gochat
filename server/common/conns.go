package common

import (
	"fmt"
	"github.com/ChinasoftNobody/gochat/common/message"
	"github.com/wonderivan/logger"
	"net"
)

//存放所有活跃连接,最大值1024
var connPool = make(map[string]net.Conn, 1024)

//允许进入排队的连接数
var connChan = make(chan net.Conn, 1024)

func init() {
	logger.Info("初始化连接池...")
	go func() {
		for {
			conn := <-connChan
			//这里开始对连接进行监听服务
			connPool[conn.LocalAddr().String()] = conn
			go serverForConn(conn)

		}
	}()
	logger.Info("初始化连接池完毕")
}

/**
新的连接尝试放入池中
*/
func NewConn(conn net.Conn) {
	go func() {
		connChan <- conn
	}()
	return
}

/**
对制定key的连接执行相应的操作
*/
func ExecToConn(connKeys []string, handle func(conn net.Conn, arg []interface{}) (err error, response []interface{})) {
	for _, key := range connKeys {
		if data, err1 := connPool[key]; err1 == false {
			fmt.Printf("连接池中未找到[%s]的连接", key)
			continue
		} else {
			go handle(data, nil)
		}

	}
}

/**
对制定key的连接执行相应的操作
*/
func ExecToAllConn(handle func(conn net.Conn, arg []interface{}) (err error, response []interface{})) {
	for _, conn := range connPool {
		go handle(conn, nil)
	}
}

/**
对成功连接至服务器的连接进行监听服务
*/
func serverForConn(conn net.Conn) {
	ExecToAllConn(func(conn net.Conn, arg []interface{}) (err error, response []interface{}) {
		ok := message.SendMessage(conn, "上线了")
		if !ok {
			logger.Error("发送消息失败：" + conn.RemoteAddr().String())
		}
		return
	})
	for {
		meta, err := message.ReceiveMessage(conn)
		if err != nil {
			return
		}
		logger.Info("新的消息:", meta.Content)
	}
}
