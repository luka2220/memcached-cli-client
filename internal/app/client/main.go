package client

import (
	"fmt"
	"net"

	"github.com/luka2220/tools/ccmc/internal/pkg/serialization"
)

type tcpClient struct {
	conn    net.Conn
	address string
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
		address,
	}, nil
}

func SendSetCommand(host string, port int, key string, value string) {
	c, err := initClient(host, port)
	if err != nil {
		e := fmt.Sprintf("An error occured initializing the tcp client: %v", err)
		panic(e)
	}

	defer c.conn.Close()

	cmd, err := serialization.SerializeCommand("set", key, 0, 0, len(value))
	if err != nil {
		e := fmt.Sprintf("An error occured serializing data: %v", err)
		panic(e)
	}

	data := serialization.SerializeDataBlock(value)

	_, err = c.conn.Write(cmd.Bytes())
	if err != nil {
		e := fmt.Sprintf("An error occured sending data to the server: %v", err)
		panic(e)
	}

	_, err = c.conn.Write(data.Bytes())
	if err != nil {
		e := fmt.Sprintf("An error occured sending data to the server: %v", err)
		panic(e)
	}

	fmt.Printf("stored key=%s, value=%s at %s", key, value, c.address)
}
