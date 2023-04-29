package nats

import "github.com/nats-io/stan.go"

func Connect(clientId string) (stan.Conn, error) {
	return stan.Connect("test-cluster", clientId, stan.NatsURL("0.0.0.0:4222"))
}