package client

import (
	"context"
	"net"

	"github.com/segmentio/kafka-go"
)

type Iproducer interface {
	Write(context.Context, []byte) error
}

type producer struct {
	writer *kafka.Writer
}

//init.... initializes local variables
func init() {
	// bServer := fmt.Sprint(viper.Get("KAFKASERVER"))
	// zookeeper = kafka.TCP(bServer)
	// topic = fmt.Sprint(viper.Get("TOPICWRITE"))
}

func NewProducer(zAddr net.Addr, topic string) Iproducer {
	return &producer{
		writer: &kafka.Writer{
			Addr:         zAddr,
			Topic:        topic,
			RequiredAcks: kafka.RequireAll,
		},
	}
}

func (p *producer) Write(ctx context.Context, msg []byte) error {
	defer p.writer.Close()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return p.writer.WriteMessages(ctx, kafka.Message{Value: msg})
}
