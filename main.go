package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var debug *bool

func main() {
	file := flag.String("input", "", "file path")
	debug = flag.Bool("dbg", false, "debug mode")
	flag.Parse()
	if strings.TrimSpace(*file) == "" {
		fmt.Println("file argument is empty")
		os.Exit(1)
	}
	if _, e := os.Stat(*file); os.IsNotExist(e) {
		fmt.Printf("file %s does not exist\n", *file)
		os.Exit(1)
	}
	Debug("Running file %s", file)
}

func Debug(f string, a ...interface{}) {
	if *debug {
		fmt.Printf(f+"\n", a)
	}
}
