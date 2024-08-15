package serialization

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSerializeCommand(t *testing.T) {
	// NOTE: Test case 1 => "set greeting 0 0 13\r\n"
	t1, err := SerializeCommand("set", "greeting", 0, 0, 13)
	if err != nil {
		t.Fatal(err)
	}

	if t1.Len() != 21 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t1.Len(), 21)
		t.Fatal(e)
	}

	if t1.String() != "set greeting 0 0 13\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t1.String(), "set greeting 0 0 13\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "add secret 0 900 44\r\n"
	t2, err := SerializeCommand("add", "secret", 0, 900, 44)
	if err != nil {
		t.Fatal(err)
	}

	if t2.Len() != 21 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t2.Len(), 21)
		t.Fatal(e)
	}

	if t2.String() != "set secret 0 900 44\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t2.String(), "set secret 0 900 44\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 3 => "replace greeting 0 0 10\r\n"
	t3, err := SerializeCommand("replace", "greeting", 0, 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	if t3.Len() != 25 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t3.Len(), 25)
		t.Fatal(e)
	}

	if t3.String() != "replace greeting 0 0 10\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t3.String(), "replace greeting 0 0 10\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 4 => "append secret 0 0 10\r\n"
	t4, err := SerializeCommand("append", "secret", 0, 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	if t4.Len() != 22 {
		e := fmt.Sprintf("Buffer size is incorrect, got=%d bytes, expected=%d bytes", t4.Len(), 22)
		t.Fatal(e)
	}

	if t4.String() != "append secret 0 0 10\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t4.String(), "append secret 0 0 10\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 4 => "prepend data 0 0 7\r\n"
	t5, err := SerializeCommand("append", "data", 0, 0, 7)
	if err != nil {
		t.Fatal(err)
	}

	if t5.Len() != 20 {
		e := fmt.Sprintf("Buffer size is incorrect, got=%d bytes, expected=%d bytes", t5.Len(), 20)
		t.Fatal(e)
	}

	if t5.String() != "prepend data 0 0 7\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t5.String(), "prepend data 0 0 7\r\n")
		t.Fatal(e)
	}
}

func TestSerializeDataBlock(t *testing.T) {
	// NOTE: Test case 1 => "Hello, World!\r\n"
	t1 := SerializeDataBlock("Hello, World!")

	if t1.Len() != 15 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d, expected=%d", t1.Len(), 15)
		t.Fatal(e)
	}

	if t1.String() != "Hello, World!\r\n" {
		e := fmt.Sprintf("Incorrect buffer string value, got=%s, expected=%s", t1.String(), "Hello, World!\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "secret message for memcached server to store\r\n"
	t2 := SerializeDataBlock("secret message for memcached server to store")

	if t2.Len() != 46 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t2.Len(), 48)
		t.Fatal(e)
	}

	if t2.String() != "secret message for memcached server to store\r\n" {
		e := fmt.Sprintf("Incorrect buffer string value, got=%s, expected=%s", t2.String(), "secret message for memcached server to store\r\n")
		t.Fatal(e)
	}
}

func TestDeserializeCommand(t *testing.T) {
	// The tcp server will responsed to a command with one of the following:
	// STORED\r\n
	// NOT_STORED\r\n

	// NOTE: Test case 1 => STORED\r\n
	t1buffer := bytes.NewBufferString("STORED\r\n")
	t1 := DeserializeCommand(t1buffer)

	if t1 != "STORED\r\n" {
		e := fmt.Sprintf("Deserialized buffer is incorrect, got=%s, expected=%s", t1, "STORED\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => NOT_STORED\r\n
	t2buffer := bytes.NewBufferString("NOT_STORED\r\n")
	t2 := DeserializeCommand(t2buffer)

	if t2 != "NOT_STORED\r\n" {
		e := fmt.Sprintf("Deserialized buffer is incorrect, got=%s, expected=%s", t2, "NOT_STORED\r\n")
		t.Fatal(e)
	}
}
