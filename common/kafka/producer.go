package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/wonderivan/logger"
	"strings"
)

var producerMap = make(map[string]sarama.AsyncProducer, 2)

/**
根据topic获取生产者
*/
func getProducer(topic string) (producer sarama.AsyncProducer, err error) {
	if producer = producerMap[topic]; producer != nil {
		return
	} else {
		err = errors.New("获取生产者失败")
	}
	return
}

/**
注册默认生产者
*/
func RegisterProducer(addrs string, topic string) (producer sarama.AsyncProducer, err error) {
	if producer = producerMap[topic]; producer != nil {
		return
	}
	producer, err = sarama.NewAsyncProducer(strings.Split(addrs, ","), nil)
	if err != nil {
		logger.Error("获取生产者失败", err)
		return
	}
	go func() {
		for {
			err := <-producer.Errors()
			logger.Error("发送消息出现异常", err)
		}
	}()
	producerMap[topic] = producer
	return
}

/**
发送消息
*/
func SendMessage(topic string, message string) {
	producer, err := getProducer(topic)
	if err != nil {
		logger.Fatal(err)
	}
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)}
}
