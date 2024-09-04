package serialization

import (
	"bytes"
	"errors"
	"fmt"
)

// Creates the byte stream for the command section of the tcp protocol
// Valid commands: set, get, add, replace, append, prepend
// FIX: Add a default case to the switch statement and test with an incorrect command passed in
func SerializeCommand(cmd string, key string, flags uint16, exptime int, size int) (*bytes.Buffer, error) {
	switch cmd {
	case "set", "get", "gets", "add", "replace", "append", "prepend":
		msg := fmt.Sprintf("%s %s %d %d %d\r\n", cmd, key, flags, exptime, size)
		byteStream := bytes.NewBufferString(msg)

		return byteStream, nil
	}

	e := fmt.Sprintf("%s is not a valid command. Enter one of set, add, replace, append, prepend", cmd)
	return nil, errors.New(e)
}

// Creates the byte stream for the cas (check and set) operation
func SerializeCASCommand(key string, flags uint16, exptime int, size int, token int) *bytes.Buffer {
	cas := fmt.Sprintf("cas %s %d %d %d %d\r\n", key, flags, exptime, size, token)
	casStream := bytes.NewBufferString(cas)

	return casStream
}

// Creates the byte stream for the delete operation
func SerializeDeleteCommand(key string) *bytes.Buffer {
	del := fmt.Sprintf("delete %s\r\n", key)
	delStream := bytes.NewBufferString(del)

	return delStream
}

// Creates the byte stream for the incr and decr commands
func SerializeIncrDecrCommand(cmd string, key string, value int) (*bytes.Buffer, error) {
	var stream *bytes.Buffer

	switch cmd {
	case "incr":
		incr := fmt.Sprintf("incr %s %d\r\n", key, value)
		stream = bytes.NewBufferString(incr)
		break
	case "decr":
		decr := fmt.Sprintf("decr %s %d\r\n", key, value)
		stream = bytes.NewBufferString(decr)
		break
	default:
		return nil, errors.New("Enter a either the incr or decr command")
	}

	return stream, nil
}

// Creates the byte stream for the datablock section of the tcp protocol
func SerializeDataBlock(dataBlock string) *bytes.Buffer {
	msg := fmt.Sprintf("%s\r\n", dataBlock)
	byteStream := bytes.NewBufferString(msg)

	return byteStream
}

// Creates a string of the byte stream response from the server
func DeserializeCommand(b *bytes.Buffer) string {
	msg := b.String()

	return msg
}
