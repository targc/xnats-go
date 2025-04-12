package xnats

import (
	"github.com/nats-io/nats.go/jetstream"
)

type Consumer struct {
	cons jetstream.Consumer
}

func (c *Consumer) Consume(handler jetstream.MessageHandler) (jetstream.ConsumeContext, error) {
	cctx, err := c.cons.Consume(handler)

	if err != nil {
		return nil, err
	}

	return cctx, nil
}
