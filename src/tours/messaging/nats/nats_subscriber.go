package nats

import (
	"github.com/nats-io/nats.go"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
)

type Subscriber struct {
	conn       *nats.EncodedConn
	subject    string
	queueGroup string
}

func NewNATSSubscriber(host, port, subject, queueGroup string) (saga.Subscriber, error) {
	conn, err := getConnection(host, port)
	if err != nil {
		return nil, err
	}
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		conn:       encConn,
		subject:    subject,
		queueGroup: queueGroup,
	}, nil
}

func (s *Subscriber) Subscribe(handler interface{}) error {
	_, err := s.conn.QueueSubscribe(s.subject, s.queueGroup, handler)
	if err != nil {
		return err
	}
	return nil
}
