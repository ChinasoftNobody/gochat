package dto

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
通用消息体，由元消息中content转化而来
*/
type CommonMsg struct {

	//消息类型
	Type       int    `json:"type"`
	Close      bool   `json:"close"`
	Content    string `json:"content"`
	TargetAddr string `json:"target_addr"`
}

type CommonMessage struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(255);comment:'asdasd'"`
}

/**
连接信息
*/
type ClientConnectDto struct {
	gorm.Model
	RemoteAddr     string `gorm:"unique_index"`
	LocalAddr      string
	ConnectTime    time.Time
	DisConnectTime time.Time
}
