package sarama

import (
	"github.com/IBM/sarama"
	"testing"
)

var addr = []string{"localhost:9094"}

func TestSyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	// 设置等待服务器确认消息成功写入的时间设置生产者返回成功的消息
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(addr, cfg)
	if err != nil {
		panic("Failed to start Sarama producer: producer err" + err.Error())
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("这是一条消息"),
		// 生产者和消费者都可以使用Headers来传递额外的信息
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("key1"),
				Value: []byte("value1"),
			},
		},
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		t.Error(err)
	}
	t.Logf("partition:%d,offset:%d", partition, offset)
}
