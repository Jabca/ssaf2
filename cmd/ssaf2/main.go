package main

import (
	"flag"
	"log"
	"path/filepath"
)

func main() {
	batchCount := flag.Int("batchSize", 1024*1024, "size of chunks (int bytes) to which files will be divided for processing")
	recursive := flag.Bool("recursive", false, "Treat input path as dir and recursively process it")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Error: no target path specified")
	}
	if len(args) > 1 {
		log.Fatal("Error: there can only be one target path")
	}

	absPath, err := filepath.Abs(args[0])

}
