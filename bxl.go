package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sync"

	"github.com/rustyoz/gokicadlib"
)

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	wd, _ := os.Getwd()
	fmt.Println(wd)
	schemlibname := flag.String("lib", "symbols", "schematic symbol library name")
	rawbxl := flag.Bool("rawbxl", false, "output raw ascii of bxl file")
	flag.Parse()
	if *cpuprofile != "" {
		fmt.Println("Starting CPU Profile")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if len(flag.Args()) < 1 {
		flag.PrintDefaults()
		return
	}
	files, _ := filepath.Glob(flag.Arg(0))
	var slib gokicadlib.SchematicLibrary
	symbolchan := make(chan *gokicadlib.Symbol)
	filechan := make(chan string)
	wg := sync.WaitGroup{}
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go ProcessFiles(filechan, *rawbxl, symbolchan, &wg)
	}
	go func() {
		for _, f := range files {
			path := filepath.Join(wd, f)
			filechan <- path
		}
		close(filechan)
		wg.Wait()
		close(symbolchan)
	}()
	for {
		s, ok := <-symbolchan
		if ok {
			slib.AddSymbol(*s)

		} else {
			break
		}
	}

	slibfile, err := os.Create(*schemlibname + ".lib")
	if err != nil {
		log.Fatal(err)
	}
	slibfile.WriteString(slib.KicadLib().String())
	slibfile.Close()

}
