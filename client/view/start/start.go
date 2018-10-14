package start

import (
	"fmt"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
)

func MainWindow() {
	//create windows;创建window窗口
	//parm 1: style of window;参数一表示创建窗口的样式
	//SW_TITLEBAR :with titlebar;顶层窗口，有标题栏
	//SW_RESIZEABLE : resizeable; 可调整大小
	//SW_CONTROLS :has max, min button;有最小/最大按钮
	//SW_MAIN : start window, if it close, application will exit; 应用程序主窗口，关闭后其他所有窗口也会关闭
	//SW_ENABLE_DEBUG : debug;可以调试
	//parm 2: the rect of window;参数二表示创建窗口的矩形
	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_RESIZEABLE|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG,
		sciter.NewRect(200, 200, 360, 600))
	if err != nil {
		log.Fatal(err)
	}
	//load file;加载文件
	w.LoadFile("view/start/start.html")
	//set title; 设置标题
	w.SetTitle("GoChat")
	defineFunctions(w)
	load(w)
	//show;显示窗口
	w.Show()
	//run, loop message;运行窗口，进入消息循环
	w.Run()
}
func load(w *window.Window) {
	root, _ := w.GetRootElement()
	searchEle, _ := root.SelectById("iptSearch")
	searchEle.AttachEventHandler(&sciter.EventHandler{})
	ele, _ := root.SelectById("ulCommunications")
	child, _ := ele.NthChild(0)
	ele1, _ := sciter.CreateElement("li", "test")
	class, _ := child.Attr("class")
	ele1.SetAttr("class", class)

	ele.Append(ele1)
	w.DefineFunction("change", func(args ...*sciter.Value) *sciter.Value {
		fmt.Println(args)
		return sciter.NewValue("1")
	})
}

//注册方法事件
func defineFunctions(w *window.Window) {
	//获取聊天列表
	w.DefineFunction("", func(args ...*sciter.Value) *sciter.Value {
		return sciter.NewValue("")
	})
}
