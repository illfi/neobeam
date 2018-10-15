package main

import (
	"flag"
	"fmt"
	"ill.fi/neobeam/interp"
	"io/ioutil"
	"os"
	"strings"
)

var debug *bool

func main() {
	file := flag.String("input", "", "file path")
	debug = flag.Bool("dbg", false, "debug mode")
	display := flag.Bool("display", false, "display interpreted file")
	flag.Parse()
	if strings.TrimSpace(*file) == "" {
		fmt.Println("file argument is empty")
		os.Exit(1)
	}
	if _, e := os.Stat(*file); os.IsNotExist(e) {
		fmt.Printf("file %s does not exist\n", *file)
		os.Exit(1)
	}
	if *display {
		Debug("Creating World...")
		dat, _ := ioutil.ReadFile(*file)
		world := interp.CreateWorld(string(dat))
		fmt.Println(world.Display())
	} else {
		// run
	}
}

func Debug(f string, a ...interface{}) {
	if *debug {
		fmt.Printf(f+"\n", a)
	}
}
