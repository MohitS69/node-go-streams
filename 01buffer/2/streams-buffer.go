package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("text.txt")
	lines := strings.Split(string(data), "\n")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
