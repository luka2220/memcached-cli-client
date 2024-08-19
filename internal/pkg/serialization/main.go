package serialization

import (
	"bytes"
	"errors"
	"fmt"
)

// Creates the byte stream for the command section of the tcp protocol
// Valid commands: set, get, add, replace, append, prepend
func SerializeCommand(cmd string, key string, flags uint16, exptime int, size int) (*bytes.Buffer, error) {

	switch cmd {
	case "set", "get", "add", "replace", "append", "prepend":
		msg := fmt.Sprintf("%s %s %d %d %d\r\n", cmd, key, flags, exptime, size)
		byteStream := bytes.NewBufferString(msg)

		return byteStream, nil
	}

	e := fmt.Sprintf("%s is not a valid command. Enter one of set, add, replace, append, prepend", cmd)
	return nil, errors.New(e)
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
