package kafka

import (
	"github.com/wonderivan/logger"
)

/**
初始化之后，注册消息的消费者及生产者
*/
func init() {
	addr := "localhost:9092"
	topic := "test"
	ConsumerRegister(addr, topic, func(msg []byte) bool {
		if len(msg)%2 == 0 {
			logger.Info("message: %v,消费成功", string(msg))
			return true
		} else {
			logger.Info("message: %v，消费失败需要重新消费", string(msg))
			return false
		}
	})
	_, err := RegisterProducer(addr, topic)
	if err != nil {
		logger.Error("注册生产者失败")
	}
}
