package set

import (
	"fmt"
	"testing"
)

func TestSetSerializeCommand(t *testing.T) {
	buffer1cmd, err := SerializeSetCommand("greeting", 0, 0, 13)
	if err != nil {
		e := fmt.Sprintf("An error occured where is shouldn't have: %v", err)
		t.Fatal(e)
	}

	// expected size is 23 bytes => set greeting 0 0 13\r\n
	if buffer1cmd.Len() != 23 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d, expected=%d", buffer1cmd.Len(), 23)
		t.Fatal(e)
	}

	if buffer1cmd.String() != "set greeting 0 0 13\r\n" {
		e := fmt.Sprintf("Incorrect buffer string, got=%s, expected=%s", buffer1cmd.String(), "set greeting 0 0 13\r\n")
		t.Fatal(e)
	}

	buffer1dataBlock, err := SerializeSetDataBlock("Hello, World!")
	if err != nil {
		e := fmt.Sprintf("An error occured where is shouldn't have: %v", err)
		t.Fatal(e)
	}

	// expected size is 17 bytes => Hello, World!\r\n
	if buffer1dataBlock.Len() != 17 {
		e := fmt.Sprintf("Buffer size incorrect, got=%d, expected=%d", buffer1dataBlock.Len(), 17)
		t.Fatal(e)
	}

	if buffer1dataBlock.String() != "Hello, World!" {
		e := fmt.Sprintf("Incorrect buffer string value, got=%s, expected=%s", buffer1dataBlock.String(), "Hello, World!")
		t.Fatal(e)
	}
}
