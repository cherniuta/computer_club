package main

import (
	"computer_club/analyze"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := strings.Join(os.Args[1:], "")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	analyze.Analyze(file)

}
