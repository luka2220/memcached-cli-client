package set

import (
	"bytes"
	"fmt"
)

// Creates the byte stream for the command section of the tcp protocol
func SerializeSetCommand(cmd string, key string, flags uint16, exptime int, size int) *bytes.Buffer {
	msg := fmt.Sprintf("%s %s %d %d %d\r\n", cmd, key, flags, exptime, size)
	byteStream := bytes.NewBufferString(msg)

	return byteStream
}

// Creates the byte stream for the datablock section of the tcp protocol
func SerializeSetDataBlock(dataBlock string) *bytes.Buffer {
	msg := fmt.Sprintf("%s\r\n", dataBlock)
	byteStream := bytes.NewBufferString(msg)

	return byteStream
}

// Creates a string of the byte stream response from the server
func DeserializeSetCommand(b *bytes.Buffer) string {
	msg := b.String()

	return msg
}
