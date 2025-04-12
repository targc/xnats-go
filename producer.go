package xnats

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go/jetstream"
)

type Producer struct {
	js jetstream.JetStream
}

func (p *Producer) Produce(ctx context.Context, subject string, message []byte) error {
	_, err := p.js.Publish(ctx, subject, message)

	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) ProduceJson(ctx context.Context, subject string, message interface{}) error {
	j, err := json.Marshal(message)

	if err != nil {
		return err
	}

	err = p.Produce(ctx, subject, j)

	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) ProduceAsync(subject string, message []byte) error {
	_, err := p.js.PublishAsync(subject, message)

	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) ProduceAsyncJson(subject string, message interface{}) error {
	j, err := json.Marshal(message)

	if err != nil {
		return err
	}

	err = p.ProduceAsync(subject, j)

	if err != nil {
		return err
	}

	return nil
}
