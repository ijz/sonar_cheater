package gameplay

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type KeyboardInput struct {
	reader *bufio.Reader
}

func NewKeyboardInput() *KeyboardInput {
	ki := new(KeyboardInput)
	ki.reader = bufio.NewReader(os.Stdin)
	return ki
}

func (ki *KeyboardInput) readLine() (string, error) {
	if s, err := ki.reader.ReadString('\n'); nil == err {
		return strings.TrimSpace(s), nil
	} else {
		return "", err
	}
}

func (ki *KeyboardInput) Prompt(msg string) (string, error) {
	fmt.Printf("%s...", msg)
	return ki.readLine()
}
