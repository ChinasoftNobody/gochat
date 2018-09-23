package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ChinasoftNobody/gochat/server/dto"
	"net"
	"time"
)

//存放所有活跃连接,最大值1024
var connPool = make(map[string]net.Conn, 1024)

//允许进入排队的连接数
var connChan = make(chan net.Conn, 1024)

func init() {
	fmt.Println("初始化连接池...")
	go func() {
		for {
			conn := <-connChan
			if len(connPool) == 1024 {
				_, err := conn.Write([]byte("连接数达到上限，继续等待中..."))
				if err != nil {
					fmt.Println("写入消息失败，断开连接", err)
					continue
				}
				connChan <- conn
			}
			_, err := conn.Write([]byte("成功连接至服务器"))
			if err != nil {
				fmt.Println("写入消息失败", err)
				continue
			}
			//这里开始对连接进行监听服务
			connPool[conn.LocalAddr().String()] = conn
			go serverForConn(conn)

		}
	}()
	fmt.Println("初始化连接池完毕")
}

/**
新的连接尝试放入池中
*/
func NewConn(conn net.Conn) (err error) {
	_, err = conn.Write([]byte("正在排队进入服务器..."))
	if err != nil {
		fmt.Println("写入消息失败", err)
		conn.Close()
		return
	}
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
		conn.Write([]byte(conn.RemoteAddr().String() + " 上线了"))
		return
	})
	for {
		byteBuff := make([]byte, 1024)
		n, err := conn.Read(byteBuff)
		if err != nil {
			fmt.Println("接收数据失败,断开客户端连接", err)
			return
		}
		err = handleMsg(&dto.SourceMsg{Addr: conn.RemoteAddr().String(), Content: string(byteBuff[:n]), Timestamp: time.Now().String()}, conn)
		if err != nil {
			return
		}
	}
}

/**
处理消息体
*/
func handleMsg(msg *dto.SourceMsg, conn net.Conn) (error error) {
	fmt.Printf("A message from %s at %s \n[%s]\n", msg.Addr, msg.Timestamp, msg.Content)
	var commonMsg dto.CommonMsg
	err := json.Unmarshal([]byte(msg.Content), &commonMsg)
	if err != nil {
		fmt.Println("不合法的消息格式，非法入侵服务，断开连接...")
		conn.Close()
		error = errors.New("不合法的消息格式，非法入侵服务，断开连接")
		return
	}
	switch commonMsg.Type {
	case MsgTypeString:
		_, err = conn.Write([]byte(commonMsg.Content))
		if err != nil {
			fmt.Println("发送消息失败：", err)
			return
		}
	}
	if commonMsg.Close {
		conn.Close()
	}
	return
}
