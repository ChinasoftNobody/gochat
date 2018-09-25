package dto

import (
	"bytes"
	"github.com/lxn/walk"
)

type IUserModel interface {
	Data() *UserModel
}

type UserInfo struct {
	Name  string
	Addr  string
	Alias string
	Value string
}

type UserModel struct {
	walk.ListModelBase
	Items []UserInfo
}

func (m *UserModel) Value(index int) interface{} {
	return m.Items[index].Name
}

func (m *UserModel) ItemCount() int {
	return len(m.Items)
}

func (m *UserModel) Data() *UserModel {
	return m
}

/**
判断名称是否包含关键字
*/
func (tmp UserInfo) ContainName(keyword string) bool {
	return bytes.Index([]byte(tmp.Name), []byte(keyword)) >= 0
}
