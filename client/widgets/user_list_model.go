package widgets

import (
	. "github.com/lxn/walk/declarative"
)

type UserListWidget struct {
	SimpleWidget
	ListBox
}

/**
为本模型注册事件处理
*/
func (tmp *UserListWidget) RegisterAction() {

}
