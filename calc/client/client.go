package client

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
)

const (
	MAX_MESSAGE_SIZE = 4096
)

type Client struct {
	Uuid       string
	Connection net.Conn
	logger     *log.Logger
	Messages   chan string
}

func New(logger *log.Logger) *Client {
	return &Client{Uuid: uuid.New().String(), logger: logger, Messages: make(chan string)}
}

func (client *Client) Connect(host string, port int32) error {
	address := fmt.Sprintf("%s:%d", host, port)
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	client.Connection = connection
	return nil
}

func (client *Client) Receive() {
	for {
		message := make([]byte, MAX_MESSAGE_SIZE)
		length, err := client.Connection.Read(message)
		if err != nil {
			client.Connection.Close()
			break
		}
		if length > 0 {
			client.log("Received %d bytes", length)
			client.Messages <- string(message)
		}
	}
}

func (client *Client) Send(message string) {
	client.Connection.Write([]byte(strings.TrimRight(message, "\n")))
}

func (client *Client) log(format string, v ...interface{}) {
	if client.logger != nil {
		client.logger.Printf(format, v...)
	}
}
