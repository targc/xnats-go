package xnats

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type ConnectOpt struct {
	Host     string
	Port     int
	User     string
	Password string
}

type XNats struct {
	nc *nats.Conn
	js jetstream.JetStream
}

func Connect(opt ConnectOpt) (*XNats, error) {
	nc, err := nats.Connect(
		fmt.Sprintf("nats://%s:%d", opt.Host, opt.Port),
		nats.UserInfo(
			opt.User,
			opt.Password,
		),
	)

	if err != nil {
		return nil, err
	}

	js, err := jetstream.New(nc)

	if err != nil {
		return nil, err
	}

	return &XNats{
		nc,
		js,
	}, nil
}

func (x *XNats) NC() *nats.Conn {
	return x.nc
}

func (x *XNats) JS() jetstream.JetStream {
	return x.js
}

func (h *XNats) Consumer(ctx context.Context, stream, consumerID string) (*Consumer, error) {
	s, err := h.js.Stream(ctx, stream)

	if err != nil {
		return nil, err
	}

	cons, err := s.Consumer(ctx, consumerID)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		cons,
	}, nil
}

func (h *XNats) Producer() *Producer {
	return &Producer{
		h.js,
	}
}

func (h *XNats) Close() {
	h.nc.Close()
}
