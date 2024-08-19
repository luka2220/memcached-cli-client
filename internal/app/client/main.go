package client

import (
	"fmt"
	"net"
)

type tcpClient struct {
	conn net.Conn
	host string
	port int
}

func initClient(host string, port int) (*tcpClient, error) {
	if host == "" {
		host = "localhost"
	}

	if port == 0 {
		port = 11211
	}

	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &tcpClient{
		conn,
		host,
		port,
	}, nil
}

func SendSetCommand(host string, port int, key string, value string) {
	c, err := initClient(host, port)
	if err != nil {
		e := fmt.Sprintf("An error occured initializing the tcp client: %v", err)
		panic(e)
	}

	defer c.conn.Close()

	c.conn.Write()
}
