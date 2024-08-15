package set

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSerializeSetCommand(t *testing.T) {
	// NOTE: Test case 1 => "set greeting 0 0 13\r\n"
	// 13 bytes for the string "Hello, World!"
	// no expiration time
	t1 := SerializeSetCommand("greeting", 0, 0, 13)

	if t1.Len() != 21 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t1.Len(), 21)
		t.Fatal(e)
	}

	if t1.String() != "set greeting 0 0 13\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t1.String(), "set greeting 0 0 13\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "set secret 0 900 44\r\n"
	// 44 bytes for the string "secret message for memcached server to store"
	// 900 second expiration time
	t2 := SerializeSetCommand("secret", 0, 900, 44)

	if t2.Len() != 21 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t2.Len(), 21)
		t.Fatal(e)
	}

	if t2.String() != "set secret 0 900 44\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t2.String(), "set secret 0 900 44\r\n")
		t.Fatal(e)
	}
}

func TestSerializeSetDataBlock(t *testing.T) {
	// NOTE: Test case 1 => "Hello, World!\r\n"
	t1 := SerializeSetDataBlock("Hello, World!")

	if t1.Len() != 15 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d, expected=%d", t1.Len(), 15)
		t.Fatal(e)
	}

	if t1.String() != "Hello, World!\r\n" {
		e := fmt.Sprintf("Incorrect buffer string value, got=%s, expected=%s", t1.String(), "Hello, World!\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "secret message for memcached server to store\r\n"
	t2 := SerializeSetDataBlock("secret message for memcached server to store")

	// expected size is 48 bytes => secret message for memcached server to store\r\n
	if t2.Len() != 46 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t2.Len(), 48)
		t.Fatal(e)
	}

	if t2.String() != "secret message for memcached server to store\r\n" {
		e := fmt.Sprintf("Incorrect buffer string value, got=%s, expected=%s", t2.String(), "secret message for memcached server to store\r\n")
		t.Fatal(e)
	}
}

func TestDeserializeSetCommand(t *testing.T) {
	// The tcp server will responsed to a command with one of the following:
	// STORED\r\n
	// NOT_STORED\r\n

	// NOTE: Test case 1 => STORED\r\n
	t1buffer := bytes.NewBufferString("STORED\r\n")
	t1 := DeserializeSetCommand(t1buffer)

	if t1 != "STORED\r\n" {
		e := fmt.Sprintf("Deserialized buffer is incorrect, got=%s, expected=%s", t1, "STORED\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => NOT_STORED\r\n
	t2buffer := bytes.NewBufferString("NOT_STORED\r\n")
	t2 := DeserializeSetCommand(t2buffer)

	if t2 != "NOT_STORED\r\n" {
		e := fmt.Sprintf("Deserialized buffer is incorrect, got=%s, expected=%s", t2, "NOT_STORED\r\n")
		t.Fatal(e)
	}
}
