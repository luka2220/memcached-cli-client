package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

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
	client, err := initClient(host, port)
	if err != nil {
		fmt.Printf("An error occured initializing the server: %v", err)
		return
	}

	defer client.conn.Close()

	cmd, err := serialization.SerializeCommand("set", key, 0, 0, len(value))
	if err != nil {
		fmt.Printf("An error occured serializing data: %v", err)
		return
	}

	data := serialization.SerializeDataBlock(value)

	_, err = client.conn.Write(cmd.Bytes())
	if err != nil {
		fmt.Printf("An error occured sending data to the server: %v", err)
		return
	}

	_, err = client.conn.Write(data.Bytes())
	if err != nil {
		fmt.Printf("An error occured sending data to the server: %v", err)
		return
	}
}

func SendGetCommand(host string, port int, key string) {
	client, err := initClient(host, port)
	if err != nil {
		fmt.Printf("An error occured connecting to the server: %v", err)
		return
	}

	defer client.conn.Close()

	cmd, err := serialization.SerializeCommand("get", key, 0, 0, 0)
	if err != nil {
		fmt.Printf("An error had occured serializing the data: %v", err)
		return
	}

	_, err = client.conn.Write(cmd.Bytes())
	if err != nil {
		fmt.Printf("An error occured sending data to the client: %v", err)
		return
	}

	reader := bufio.NewReader(client.conn)

	var response string

	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("Error reading from the server: ", err)
			break
		}

		if message == "END\r\n" {
			break
		}

		response += message
	}

	// BUG: Something is going wrong here...
	// I need to find a better way to deserialize and parse the response from the server...
	// Another issue if we try to get a key that does not exist and index out of range error is thrown...
	v := strings.Split(response, "\r\n")
	fmt.Println(v[len(v)-2])
}

func SendAddCommand(host string, port int, key string, value string) {
	client, err := initClient(host, port)
	if err != nil {
		fmt.Println("Error connecting to the client: ", err)
		return
	}

	defer client.conn.Close()

	buffCmd, err := serialization.SerializeCommand("add", key, 0, 0, len(value))
	if err != nil {
		fmt.Println("Error serializing add command: ", err)
		return
	}

	_, err = client.conn.Write(buffCmd.Bytes())
	if err != nil {
		fmt.Println("Error sending data to the server: ", err)
	}

	dataCmd := serialization.SerializeDataBlock(value)
	_, err = client.conn.Write(dataCmd.Bytes())
	if err != nil {
		fmt.Println("Error sending data to the server: ", err)
	}

	reader := bufio.NewReader(client.conn)

	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("Error reading data from server: ", err)
			return
		}

		switch message {
		case "ERROR\r\n":
			fmt.Println("The server responded with an error")
			return
		case "NOT_STORED\r\n":
			fmt.Println("key already exists in the cache")
			return
		case "STORED\r\n":
			return
		default:
			continue
		}
	}
}

func SendReplaceCommand(host string, port int, key string, value string) {
	client, err := initClient(host, port)
	if err != nil {
		fmt.Println("Error starting the TCP client: ", err)
		return
	}

	defer client.conn.Close()

	buffCmd, err := serialization.SerializeCommand("replace", key, 0, 0, len(value))
	if err != nil {
		fmt.Println("Error serializing the command: ", err)
		return
	}

	_, err = client.conn.Write(buffCmd.Bytes())
	if err != nil {
		fmt.Println("Error sending command to server: ", err)
		return
	}

	buffData := serialization.SerializeDataBlock(value)
	_, err = client.conn.Write(buffData.Bytes())
	if err != nil {
		fmt.Println("Error sending data block to server: ", err)
	}

	reader := bufio.NewReader(client.conn)

	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("Error reading response from server: ", err)
			return
		}

		switch message {
		case "ERROR\r\n":
			fmt.Println("Error replacing value: ", err)
			return
		case "NOT_STORED\r\n":
			fmt.Printf("Error replacing %s, does not exist on server\n", key)
			return
		case "STORED\r\n":
			return
		default:
			continue
		}
	}
}

func SendAppendCommand(host string, port int, key string, value string) {
	client, err := initClient(host, port)
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}

	defer client.conn.Close()

	buffCmd, err := serialization.SerializeCommand("append", key, 0, 0, len(value))
	if err != nil {
		fmt.Println("Error serializing command: ", err)
		return
	}

	_, err = client.conn.Write(buffCmd.Bytes())
	if err != nil {
		fmt.Println("Error sending command to server: ", err)
		return
	}

	buffData := serialization.SerializeDataBlock(value)
	_, err = client.conn.Write(buffData.Bytes())

	reader := bufio.NewReader(client.conn)

	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("Error reading from the server: ", err)
			return
		}

		switch message {
		case "ERROR\r\n":
			fmt.Println("Server responded with an error...")
			return
		case "NOT_STORED\r\n":
			fmt.Printf("%s is not stored in the memcached server\n", key)
			return
		case "STORED\r\n":
			return
		default:
			continue
		}
	}
}

func SendPrependCommmand(host string, port int, key string, value string) {
	client, err := initClient(host, port)
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}

	defer client.conn.Close()

	buffCmd, err := serialization.SerializeCommand("prepend", key, 0, 0, len(value))
	if err != nil {
		fmt.Println("Error serializing command: ", err)
		return
	}

	_, err = client.conn.Write(buffCmd.Bytes())
	if err != nil {
		fmt.Println("Error writting command to server: ", err)
		return
	}

	buffData := serialization.SerializeDataBlock(value)
	_, err = client.conn.Write(buffData.Bytes())
	if err != nil {
		fmt.Println("Error writting data to server:", err)
		return
	}

	reader := bufio.NewReader(client.conn)

	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("Error reading data from server: ", err)
			return
		}

		switch message {
		case "ERROR\r\n":
			fmt.Println("Server responded with an error")
			return
		case "NOT_STORED\r\n":
			fmt.Println("key & value not stored, check if the key exists on the server")
			return
		case "STORED\r\n":
			return
		default:
			continue
		}
	}
}

func SendCasCommand(host string, port int, key string, value string, token int) {
	client, err := initClient(host, port)
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}

	defer client.conn.Close()

	bufferCmd := serialization.SerializeCASCommand(key, 0, 0, len(value), token)
	_, err = client.conn.Write(bufferCmd.Bytes())
	if err != nil {
		fmt.Println("Error sending data to server: ", err)
		return
	}

	bufferData := serialization.SerializeDataBlock(value)
	_, err = client.conn.Write(bufferData.Bytes())
	if err != nil {
		fmt.Println("Error sending data to server: ", err)
		return
	}

	reader := bufio.NewReader(client.conn)

	for {
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("Error reading data from server: ", err)
			return
		}

		switch message {
		case "ERROR\r\n":
			fmt.Println("Error from server, could not parse command...")
			return
		case "NOT_FOUND\r\n":
			fmt.Println("Error, key does not exist on server...")
			return
		case "EXISTS\r\n":
			fmt.Println("Error, CAS token provided by the client does not match the current version of the item on the server")
			return
		case "STORED\r\n":
			return
		default:
			continue
		}
	}
}

func SendGetsCommand(host string, port int, key string) {
	client, err := initClient(host, port)
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}

	defer client.conn.Close()
}
