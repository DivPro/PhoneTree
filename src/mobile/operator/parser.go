package operator

import (
	"encoding/csv"
	"bufio"
	"os"
	"log"
	"io"
	"fmt"
)

type Parser struct {
	filePath string
}

func NewParser(filePath string) *Parser {
	return &Parser{filePath:filePath}
}

func (parser *Parser) Parse() {
	fmt.Println(parser.filePath)
	csvFile, _ := os.Open(parser.filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	reader.LazyQuotes = true
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		if len(line) != 6 {
			continue
		}

		fmt.Println(len(line), line)
	}
}