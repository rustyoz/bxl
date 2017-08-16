package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/rustyoz/bxl/bxlbin"
	"github.com/rustyoz/bxl/bxlparser"
	"github.com/rustyoz/gokicadlib"
)

func DecodeFile(path string) (string, error) {
	infile, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file: ", path, err)
	}

	decoder := bxlbin.NewDecoder()
	output, decodeerr := decoder.Decode(infile)

	return output, decodeerr
}

func ProcessFiles(filespath chan string, saveraw bool, symbchan chan *gokicadlib.Symbol, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		f, ok := <-filespath
		if !ok {
			return
		}

		ProcessFile(f, saveraw, symbchan)

	}
}

func ProcessFile(f string, saveraw bool, symbchan chan *gokicadlib.Symbol) {
	output, err := DecodeFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	if saveraw {
		var outfile *os.File
		outfile, err = os.Create(f + ".txt")
		_, err = outfile.WriteString(output)

		if err != nil {
			log.Println("Error writing file: ", f+".txt", err)
		}
		outfile.Close()
	}
	parser := bxlparser.NewBxlParser()
	parser.Parse(output)
	for _, p := range parser.Patterns {
		m, err := p.ToKicad()
		if err != nil {
			log.Fatal(err)
		}
		var module_file *os.File
		module_file, err = os.Create(strings.Replace(m.Name, " /", "_", -1) + ".kicad_mod")
		_, err = module_file.WriteString(m.ToSExp())
		module_file.Close()

	}

	if len(parser.Symbols) > 0 {
		fmt.Println(f, "Symbols:", len(parser.Symbols))

		for _, s := range parser.Symbols {
			sym, err := s.Kicad()
			if err != nil {
				fmt.Print(err)
			}
			symbchan <- sym
		}
	}

}
