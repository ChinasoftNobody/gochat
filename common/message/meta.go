/**
定义消息格式及类型
*/
package message

import "time"

/**
元消息
*/
type Meta struct {
	Type     Type        `json:"type"`
	Content  interface{} `json:"content"`
	Source   string      `json:"source"`
	SendTime time.Time   `json:"send_time"`
}
