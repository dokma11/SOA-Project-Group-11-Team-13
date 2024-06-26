package nats

import (
	"github.com/nats-io/nats.go"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
)

type Publisher struct {
	conn    *nats.EncodedConn
	subject string
}

func NewNATSPublisher(host, port, subject string) (saga.Publisher, error) {
	conn, err := getConnection(host, port)
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Publisher{
		conn:    encConn,
		subject: subject,
	}, nil
}

func (p *Publisher) Publish(message interface{}) error {
	err := p.conn.Publish(p.subject, message)
	if err != nil {
		return err
	}
	return nil
}
