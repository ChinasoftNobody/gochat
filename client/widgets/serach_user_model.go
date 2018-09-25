package widgets

import (
	"fmt"
	"github.com/ChinasoftNobody/gochat/client/dto"
	"github.com/lxn/walk/declarative"
)

type SearchUserWidget struct {
	declarative.LineEdit
	SimpleWidget
}

/**
这里实现基本的事件注册方法
*/
func (tmp *SearchUserWidget) RegisterAction() {
	tmp.OnTextChanged = func() {
		keyword := (*tmp.AssignTo).Text()
		fmt.Println(keyword)
		var model = (*SingleWindow().UserListWidget.AssignTo).Model()
		value, ok := model.(dto.IUserModel)
		if ok {
			value.Data().Items = append(value.Data().Items, dto.UserInfo{Name: keyword})
			(*SingleWindow().UserListWidget.AssignTo).SetModel(value)
		}

	}
}
