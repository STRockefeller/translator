package output

import (
	"fmt"
	"os"
)

func WriteToFile(filepath, content string) error {
	return os.WriteFile(filepath, []byte(content), 0644)
}

func PrintToConsole(content string) {
	fmt.Println(content)
}
