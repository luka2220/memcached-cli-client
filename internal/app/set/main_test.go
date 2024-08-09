package set

import (
	"testing"
)

func TestSetSerializeCommand(t *testing.T) {
	t.Fatal("No tests written for serialize set command")

	// Function should serialize some type of data (byte stream) to be sent to the memcached server
	SerializeSetCommand()
}

func TestDeserializeSetCommand(t *testing.T) {
	t.Fatal("No tests written for the deserialize set command")

	// Function should deserialize some data (byte stream) from the memcached server response
	DeserializeSetCommand()
}
