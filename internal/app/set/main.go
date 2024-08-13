package set

import (
	"bytes"
	"errors"
)

// Creates the byte stream for the command section of the tcp protocol
func SerializeSetCommand(key string, flags uint16, exptime int, bytes int) (*bytes.Buffer, error) {
	return nil, errors.New("Function not yet implemented...")
}

// Creates the byte stream for the datablock section of the tcp protocol
func SerializeSetDataBlock(dataBlock string) (*bytes.Buffer, error) {
	return nil, errors.New("Function not yet implemented...")
}
