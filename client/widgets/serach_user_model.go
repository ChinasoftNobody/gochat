package widgets

import (
	"encoding/json"
	"github.com/ChinasoftNobody/gochat/client/dto"
	"github.com/lxn/walk/declarative"
)

type SearchUserModel struct {
	declarative.LineEdit
	SimpleModel
}

/**
实现model接口
*/
func (model *SearchUserModel) RegisterAction() {
	model.OnTextChanged = func() {
		keyword := (*model.AssignTo).Text()
		fileUserList := make([]dto.UserInfo, 0)
		for _, user := range *SingleWindow().UserListModel.Data {
			if user.ContainName(keyword) {
				fileUserList = append(fileUserList, user)
			}
		}
		data, _ := json.Marshal(fileUserList)
		SingleWindow().UserListModel.BindData(data)
	}
}
