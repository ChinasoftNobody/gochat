package dto

import "bytes"

type UserInfo struct {
	Name  string
	Addr  string
	Alias string
}

/**
判断名称是否包含关键字
*/
func (tmp UserInfo) ContainName(keyword string) bool {
	return bytes.Index([]byte(tmp.Name), []byte(keyword)) >= 0
}
