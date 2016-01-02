package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rustyoz/bxl/bxlbin"
	"github.com/rustyoz/bxl/bxlparser"
)

func DecodeFile(path string) (string, error) {
	infile, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file: ", path, err)
	}
	os.Create(path + ".txt")

	decoder := bxlbin.NewDecoder()
	output, decodeerr := decoder.Decode(infile)
	var outfile *os.File
	outfile, err = os.Create(path + ".txt")
	_, err = outfile.WriteString(output)

	if err != nil {
		log.Println("Error writing file: ", path+".txt", err)
	}
	fmt.Println("Output characters: ", len(output))
	outfile.Close()
	return output, decodeerr
}

func main() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	if len(os.Args) < 2 {
		fmt.Println("Usage: bxl filename.bxl or bxl *.bxl \n Currently only output raw ascii as text files")
		return
	}
	if strings.HasPrefix(os.Args[1], "*.") {
		files, _ := filepath.Glob(os.Args[1])
		for _, f := range files {
			fmt.Println(f)
			path := filepath.Join(wd, f)
			output, err := DecodeFile(path)
			if err != nil {
				parser := bxlparser.NewBxlParser()
				parser.Parse(output)
			}
		}

	} else {
		output, err := DecodeFile(os.Args[1])
		if err != nil {
			parser := bxlparser.NewBxlParser()
			parser.Parse(output)
		}
	}

}
