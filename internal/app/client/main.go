package client

import (
	"bufio"
	"fmt"
	"github.com/luka2220/tools/ccmc/internal/pkg/serialization"
	"net"
	"strings"
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
	client, err := initClient(host, port)
	if err != nil {
		e := fmt.Sprintf("An error occured initializing the server: %v", err)
		panic(e)
	}

	defer client.conn.Close()

	cmd, err := serialization.SerializeCommand("set", key, 0, 0, len(value))
	if err != nil {
		e := fmt.Sprintf("An error occured serializing data: %v", err)
		panic(e)
	}

	data := serialization.SerializeDataBlock(value)

	_, err = client.conn.Write(cmd.Bytes())
	if err != nil {
		e := fmt.Sprintf("An error occured sending data to the server: %v", err)
		panic(e)
	}

	_, err = client.conn.Write(data.Bytes())
	if err != nil {
		e := fmt.Sprintf("An error occured sending data to the server: %v", err)
		panic(e)
	}
}

func SendGetCommand(host string, port int, key string) {
	client, err := initClient(host, port)
	if err != nil {
		e := fmt.Sprintf("An error occured connecting to the server: %v", err)
		panic(e)
	}

	defer client.conn.Close()

	cmd, err := serialization.SerializeCommand("get", key, 0, 0, 0)
	if err != nil {
		e := fmt.Sprintf("An error had occured serializing the data: %v", err)
		panic(e)
	}

	_, err = client.conn.Write(cmd.Bytes())
	if err != nil {
		e := fmt.Sprintf("An error occured sending data to the client: %v", err)
		panic(e)
	}

	reader := bufio.NewReader(client.conn)

	// Listen for a response from the server
	var response string

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from the server: ", err)
			break
		}

		if message == "END\r\n" {
			break
		}

		response += message
	}

	// Extract the stored value from the server's response
	v := strings.Split(response, "\r\n")
	fmt.Println(v[len(v)-2])
}
