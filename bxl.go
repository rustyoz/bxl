package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rustyoz/bxl/bxlbin"
	"github.com/rustyoz/bxl/bxlparser"
)

func main() {

	infile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error opening file: ", os.Args[1], err)
	}

	os.Create(os.Args[1] + ".txt")
	outfile, err := os.OpenFile(os.Args[1]+".txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening file: ", os.Args[1]+".txt", err)
	}

	decoder := bxlbin.NewDecoder()
	output, err := decoder.Decode(infile)

	var characters int
	characters, err = outfile.WriteString(output)
	fmt.Println("Characters: ", characters)
	if err != nil {
		fmt.Println(err)
	}
	outfile.Close()

	parser := bxlparser.NewBxlParser()
	parser.Parse(output)

}
