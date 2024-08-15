package serialization

import (
	"bytes"
	"fmt"
)

// Creates the byte stream for the command section of the tcp protocol
func SerializeCommand(cmd string, key string, flags uint16, exptime int, size int) (*bytes.Buffer, error) {
	msg := fmt.Sprintf("%s %s %d %d %d\r\n", cmd, key, flags, exptime, size)
	byteStream := bytes.NewBufferString(msg)

	return byteStream, nil
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
