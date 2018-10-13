package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/wonderivan/logger"
	"strings"
)

/**
注册消费者
*/
func ConsumerRegister(addrs string, topic string, handler func(msg []byte) bool) bool {
	consumer, err := sarama.NewConsumer(strings.Split(addrs, ","), nil)
	if err != nil {
		logger.Error("Failed to start consumer: %s", err)
	}
	if consumer == nil || handler == nil {
		logger.Error("consumer is empty")
		return false
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		logger.Error("Failed to get the list of partitions: ", err)
		return false
	}
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logger.Info("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				logger.Info("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				flag := handler(msg.Value)
				if flag {
					logger.Info("handle message success")
				} else {
					logger.Error("handle message failed")

				}
			}
		}(pc)
	}
	return true
}
