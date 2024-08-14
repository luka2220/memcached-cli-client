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
	t1, err := SerializeSetCommand("greeting", 0, 0, 13)
	if err != nil {
		e := fmt.Sprintf("An error occured calling SerializeSetCommand: %v", err)
		t.Fatal(e)
	}

	// expected command size is 23 bytes => set greeting 0 0 13\r\n
	if t1.Len() != 23 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t1.Len(), 23)
		t.Fatal(e)
	}

	if t1.String() != "set greeting 0 0 13\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t1.String(), "set greeting 0 0 13\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "set secret 0 900 44\r\n"
	// 44 bytes for the string "secret message for memcached server to store"
	// 900 second expiration time
	t2, err := SerializeSetCommand("secret", 0, 900, 44)
	if err != nil {
		e := fmt.Sprintf("An error occured calling SerializeSetCommand: %v", err)
		t.Fatal(e)
	}

	// expected command size is 23 bytes => set secret 0 900 44\r\n
	if t2.Len() != 23 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t2.Len(), 23)
		t.Fatal(e)
	}

	if t2.String() != "set secret 0 900 44\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t2.String(), "set secret 0 900 44\r\n")
		t.Fatal(e)
	}
}

func TestSerializeSetDataBlock(t *testing.T) {
	// NOTE: Test case 1 => "Hello, World!\r\n"
	t1, err := SerializeSetDataBlock("Hello, World!")
	if err != nil {
		e := fmt.Sprintf("An error occured calling SerializeSetDataBlock: %v", err)
		t.Fatal(e)
	}

	// expected size is 17 bytes => Hello, World!\r\n
	if t1.Len() != 17 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d, expected=%d", t1.Len(), 17)
		t.Fatal(e)
	}

	if t1.String() != "Hello, World!\r\n" {
		e := fmt.Sprintf("Incorrect buffer string value, got=%s, expected=%s", t1.String(), "Hello, World!\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "secret message for memcached server to store\r\n"
	t2, err := SerializeSetDataBlock("secret message for memcached server to store")
	if err != nil {
		e := fmt.Sprintf("An error occured calling SerializeSetDataBlock: %v", err)
		t.Fatal(e)
	}

	// expected size is 48 bytes => secret message for memcached server to store\r\n
	if t2.Len() != 48 {
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

	t1, err := DeserializeSetCommand(t1buffer)
	if err != nil {
		e := fmt.Sprintf("An error occured calling DeserializeSetCommand: %v", err)
		t.Fatal(e)
	}

	if t1 != "STORED\r\n" {
		e := fmt.Sprintf("Deserialized buffer is incorrect, got=%s, expected=%s", t1, "STORED\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => NOT_STORED\r\n
	t2buffer := bytes.NewBufferString("NOT_STORED\r\n")

	t2, err := DeserializeSetCommand(t2buffer)
	if err != nil {
		e := fmt.Sprintf("An error occured calling DeserializeSetCommand: %v", err)
		t.Fatal(e)
	}

	if t1 != "NOT_STORED\r\n" {
		e := fmt.Sprintf("Deserialized buffer is incorrect, got=%s, expected=%s", t2, "NOT_STORED\r\n")
		t.Fatal(e)
	}
}
