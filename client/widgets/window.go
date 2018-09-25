package widgets

import (
	"github.com/ChinasoftNobody/gochat/client/dto"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"sync"
)

const WindowWidth = 300
const WindowHeight = 600
const WindowTitle = "GoChat"

/**
主window
*/
var chatMainWindow *ChatMainWindow

/**
主窗口互斥锁，避免并发操作导致资源竞争
*/
var mainWindowLock sync.Once

/**
主界面窗口类型
*/
type ChatMainWindow struct {
	MainWindow       *MainWindow
	SearchUserWidget *SearchUserWidget
	UserListWidget   *UserListWidget
}

/**
定义启动方法
*/
func (chatWindow ChatMainWindow) RunChart() {
	chatWindow.MainWindow.Run()
}

/**
启动主客户端入口
*/
func SingleWindow() *ChatMainWindow {
	if chatMainWindow == nil {
		mainWindowLock.Do(initChatMainWindow)
	}
	return chatMainWindow
}

/**
获取ChatWindow窗口，并初始化组件信息
*/
func initChatMainWindow() {

	chatMainWindow = &ChatMainWindow{}
	var searchLineEdit *walk.LineEdit
	var userListBox *walk.ListBox
	var model dto.IUserModel = &dto.UserModel{Items: []dto.UserInfo{}}
	chatMainWindow.UserListWidget = &UserListWidget{
		ListBox: ListBox{
			AssignTo: &userListBox,
			Model:    model,
		},
	}
	chatMainWindow.SearchUserWidget = &SearchUserWidget{
		LineEdit: LineEdit{AssignTo: &searchLineEdit},
	}
	chatMainWindow.SearchUserWidget.RegisterAction()
	chatMainWindow.UserListWidget.RegisterAction()
	chatMainWindow.MainWindow = &MainWindow{
		Title:   WindowTitle,
		Size:    Size{Width: WindowWidth, Height: WindowHeight},
		Layout:  VBox{},
		Enabled: true,
		Children: []Widget{
			*chatMainWindow.SearchUserWidget,
			*chatMainWindow.UserListWidget,
		},
	}

	return
}
