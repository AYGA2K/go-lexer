package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/AYGA2K/go-lexer/lexer"
)

func main() {
	code, err := readFileContent()
	if err != nil {
		log.Fatalf("erorr reading the code file: %s", err.Error())
	}

	l := lexer.NewLexer(code)

	for tok := l.NextToken(); tok.Typ != lexer.EOF; tok = l.NextToken() {
		fmt.Printf("%+v \n", tok)
	}
}
func readFileContent() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(cwd, "code")
	file := filepath.Join(dir, "main.go")

	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
