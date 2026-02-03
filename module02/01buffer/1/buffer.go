package main

import "fmt"

func main() {
	b := make([]byte, 5)
	b = append(b, 'h')
	fmt.Printf("%s", b[5:])
}
