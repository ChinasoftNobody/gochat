/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package widgets

/**
定义模型需要实现的接口
*/
type Model interface {
	RegisterAction()
	BindData([]byte)
}

/**
定义基本模型类型
实现基本的事件注册
*/
type SimpleModel struct {
	//数据信息
	Data interface{}
}

/**
这里实现基本的时间注册方法
*/
func (tmp SimpleModel) RegisterAction() {
	//这里不做任何业务处理，只是实现了当前方法
}

/**
基本赋值
*/
func (tmp SimpleModel) BindData(data []byte) {
	tmp.Data = data
}
