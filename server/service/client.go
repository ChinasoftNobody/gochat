/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package service

import (
	"fmt"
	"github.com/ChinasoftNobody/gochat/server/db"
	"github.com/ChinasoftNobody/gochat/server/dto"
	"net"
	"time"
)

/**
新连接连接至服务器
记录客户端信息至数据库
*/
func NewClientConnect(conn net.Conn) {
	fmt.Printf("新的客户端连接[%s]\n", conn.RemoteAddr().String())
	clientConnect := dto.ClientConnectDto{RemoteAddr: conn.RemoteAddr().String(), LocalAddr: conn.LocalAddr().String(),
		ConnectTime: time.Now()}
	db.DB.Create(&clientConnect)
	fmt.Println(clientConnect.ID)
}
