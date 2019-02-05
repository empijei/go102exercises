package main

import (
	"io"
	"log"
	"os"
)

func main() {
	var fs []io.Reader
	for _, fname := range os.Args[1:] {
		f, err := os.Open(fname)
		if err != nil {
			log.Fatalf("Could not process file: %v", err)
		}
		defer f.Close()
		fs = append(fs, f)
	}
	_, err := io.Copy(os.Stdout, io.MultiReader(fs...))
	if err != nil {
		log.Fatalf("Error during copy: %v", err)
	}
}
