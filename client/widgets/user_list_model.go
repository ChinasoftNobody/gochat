package widgets

import (
	"encoding/json"
	"github.com/ChinasoftNobody/gochat/client/dto"
	. "github.com/lxn/walk/declarative"
)

type UserListModel struct {
	SimpleModel
	ListBox
	Data *[]dto.UserInfo
}

/**
为本模型注册事件处理
*/
func (tmp *UserListModel) RegisterAction() {
	data, _ := json.Marshal([]dto.UserInfo{
		{"a", "a", "a"},
		{"b", "b", "b"},
	})
	tmp.BindData(data)
}

/**
为列表组建绑定数据
*/
func (tmp *UserListModel) BindData(data []byte) {
	var userList *[]dto.UserInfo
	err := json.Unmarshal(data, &userList)
	if err != nil {
		panic("数据类型转化失败")
	}
	var userNames = make([]string, len(*userList))
	for i := 0; i < len(*userList); i++ {
		userNames[i] = (*userList)[i].Name
	}
	tmp.Data = userList
	//TODO 这里还有一点问题，为什么拿不到对象呢
	//(*tmp.ListBox.AssignTo).SetModel(userNames)
}
