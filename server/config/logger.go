/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package config

import (
	"github.com/wonderivan/logger"
)

//初始化日志配置
func InitLog() {
	logger.SetLogger(`{
  "TimeFormat":"2006-01-02 15:04:05",
  "Console": {
    "level": "TRAC",
    "color": true
  },
  "File": {
    "filename": "serv.log",
    "level": "TRAC",
    "daily": true,
    "maxlines": 1000000,
    "maxsize": 1,
    "maxdays": -1,
    "append": true,
    "permit": "0660"
  }
}`)
	logger.Debug("日志初始化")
}
