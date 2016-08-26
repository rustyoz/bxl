package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rustyoz/bxl/bxlbin"
	"github.com/rustyoz/bxl/bxlparser"
	"github.com/rustyoz/gokicadlib"
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
		fmt.Println("Usage: bxl filename.bxl or bxl *.bxl")
		return
	}
	if strings.HasPrefix(os.Args[1], "*.") {
		files, _ := filepath.Glob(os.Args[1])
		fmt.Println(files)
		for _, f := range files {
			fmt.Println(f)
			path := filepath.Join(wd, f)
			fmt.Println(path)
			ProcessFile(path)

		}
	} else {
		ProcessFile(os.Args[1])
	}

}

func ProcessFile(f string) {
	output, err := DecodeFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	parser := bxlparser.NewBxlParser()
	parser.Parse(output)
	for _, p := range parser.Patterns {
		m := p.ToKicad()
		var module_file *os.File
		module_file, err = os.Create(strings.Replace(m.Name, " ", "_", -1) + ".kicad_mod")
		_, err = module_file.WriteString(m.ToSExp())
		module_file.Close()

	}
	var schematiclib_file *os.File
	schematiclib_file, err = os.Create(strings.Replace(parser.Symbol.Name, " ", "_", -1) + ".lib")
	schematiclib := &gokicadlib.SchematicLibrary{}
	schematiclib.AddSymbol(*parser.Symbol.Kicad())
	schematiclib.KicadLib().WriteTo(schematiclib_file)
	schematiclib_file.Close()
}
