package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func getConnection(host, port string) (*nats.Conn, error) {
	url := fmt.Sprintf("nats://%s:%s", host, port)
	connection, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
