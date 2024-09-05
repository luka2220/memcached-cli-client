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

	// NOTE: Test case 2 => "get store 0 0 0\r\n"
	t2, err := SerializeCommand("get", "store", 0, 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	if t2.Len() != 17 {
		e := fmt.Sprintf("Incorrect buffer size, got=%d, expected=%d", t2.Len(), 17)
		t.Fatal(e)
	}

	if t2.String() != "get store 0 0 0\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t2.String(), "get store 0 0 0\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 3 => "add secret 0 900 44\r\n"
	t3, err := SerializeCommand("add", "secret", 0, 900, 44)
	if err != nil {
		t.Fatal(err)
	}

	if t3.Len() != 21 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t3.Len(), 21)
		t.Fatal(e)
	}

	if t3.String() != "add secret 0 900 44\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t3.String(), "add secret 0 900 44\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 4 => "replace greeting 0 0 10\r\n"
	t4, err := SerializeCommand("replace", "greeting", 0, 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	if t4.Len() != 25 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d bytes, expected=%d bytes", t4.Len(), 25)
		t.Fatal(e)
	}

	if t4.String() != "replace greeting 0 0 10\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t4.String(), "replace greeting 0 0 10\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 5 => "append secret 0 0 10\r\n"
	t5, err := SerializeCommand("append", "secret", 0, 0, 10)
	if err != nil {
		t.Fatal(err)
	}

	if t5.Len() != 22 {
		e := fmt.Sprintf("Buffer size is incorrect, got=%d bytes, expected=%d bytes", t5.Len(), 22)
		t.Fatal(e)
	}

	if t5.String() != "append secret 0 0 10\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t5.String(), "append secret 0 0 10\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 6 => "prepend data 0 0 7\r\n"
	t6, err := SerializeCommand("prepend", "data", 0, 0, 7)
	if err != nil {
		t.Fatal(err)
	}

	if t6.Len() != 20 {
		e := fmt.Sprintf("Buffer size is incorrect, got=%d bytes, expected=%d bytes", t6.Len(), 20)
		t.Fatal(e)
	}

	if t6.String() != "prepend data 0 0 7\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t6.String(), "prepend data 0 0 7\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 7 => "gets store42069 1 900 0\r\n"
	t7, err := SerializeCommand("gets", "store42069", 1, 900, 0)
	if err != nil {
		t.Fatal(err)
	}

	if t7.Len() != 25 {
		e := fmt.Sprintf("Buffer size is incorrect, got=%d bytes, expected=%d bytes", t7.Len(), 25)
		t.Fatal(e)
	}

	if t7.String() != "gets store42069 1 900 0\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t7.String(), "gets store42069 1 900 0\r\n")
		t.Fatal(e)
	}
}

func TestSerializeCASCommand(t *testing.T) {
	// NOTE: Test case 1 => "cas key 0 0 5 1\r\n"
	t1 := SerializeCASCommand("key", 0, 0, 5, 1)

	if t1.Len() != 17 {
		e := fmt.Sprintf("Incorrect buffer length, got=%d, expected=%d", t1.Len(), 17)
		t.Fatal(e)
	}

	if t1.String() != "cas key 0 0 5 1\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t1.String(), "cas key 0 0 5 1\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "cas storage 1 900 25 7\r\n"
	t2 := SerializeCASCommand("storage", 1, 900, 25, 7)

	if t2.Len() != 24 {
		e := fmt.Sprintf("Incorrect buffer length, got=%d, expected=%d", t2.Len(), 24)
		t.Fatal(e)
	}

	if t2.String() != "cas storage 1 900 25 7\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t2.String(), "cas storage 1 900 25 7\r\n")
		t.Fatal(e)
	}
}

func TestSerializeDeleteCommand(t *testing.T) {
	// NOTE: Test case 1 => "delete key\r\n"
	t1 := SerializeDeleteCommand("key")

	if t1.Len() != 12 {
		e := fmt.Sprintf("Incorrect buffer size, got=%d, expected=%d", t1.Len(), 12)
		t.Fatal(e)
	}

	if t1.String() != "delete key\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t1.String(), "delete key\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "delete storage54321\r\n"
	t2 := SerializeDeleteCommand("storage54321")

	if t2.Len() != 21 {
		e := fmt.Sprintf("Incorrect buffer size, got=%d, expected=%d", t2.Len(), 21)
		t.Fatal(e)
	}

	if t2.String() != "delete storage54321\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t2.String(), "delete storage54321\r\n")
		t.Fatal(e)
	}
}

func TestSerializeIncrDecrCommand(t *testing.T) {
	// NOTE: Test case 1 => "incr web_key 2\r\n"
	t1, err := SerializeIncrDecrCommand("incr", "web_key", 2)
	if err != nil {
		t.Fatal(err)
	}

	if t1.String() != "incr web_key 2\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t1.String(), "incr web_key 2\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 2 => "decr unknown 1\r\n"
	t2, err := SerializeIncrDecrCommand("decr", "unknown", 1)
	if err != nil {
		t.Fatal(err)
	}

	if t2.String() != "decr unknown 1\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", t2.String(), "decr unknown 1\r\n")
		t.Fatal(e)
	}

	// NOTE: Test case 3 => "add web_key 2\r\n"
	_, err = SerializeIncrDecrCommand("add", "web_key", 2)
	if err == nil {
		e := fmt.Sprintf("An error should be thrown here for invalid command... eith incr or decr")
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
