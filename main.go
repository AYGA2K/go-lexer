package main

import "fmt"

type itemType int
type item struct {
	typ   itemType
	value string
}

const (
	itemError itemType = iota
	itemDot
	itemEOF
)

func main() {
	fmt.Println("YAAAAAAAAAAAAAAAAY")
}
