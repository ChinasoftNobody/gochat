package widgets

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"sync"
)

const WindowWidth = 300
const WindowHeight = 600
const WindowTitle = "GoChat"

/**
查询用户的输入框
*/
var searchLineEdit *walk.LineEdit

/**
用户列表框
*/
var userListBox *walk.ListBox

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
	MainWindow      *MainWindow
	SearchUserModel *SearchUserModel
	UserListModel   *UserListModel
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

	contactWidget := &UserListModel{
		ListBox: ListBox{
			AssignTo: &userListBox,
		},
	}

	contactWidget.RegisterAction()
	lineEdit := &SearchUserModel{}
	lineEdit.AssignTo = &searchLineEdit
	lineEdit.RegisterAction()
	mainWindow := &MainWindow{
		Title:   WindowTitle,
		Size:    Size{Width: WindowWidth, Height: WindowHeight},
		Layout:  VBox{},
		Enabled: true,
		Children: []Widget{
			lineEdit,
			contactWidget,
		},
	}

	chatMainWindow = &ChatMainWindow{
		mainWindow,
		lineEdit,
		contactWidget,
	}
	return
}
