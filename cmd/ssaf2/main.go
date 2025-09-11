package main

import (
	"flag"
	"log"
	"path/filepath"
	"ssaf2/internal/appInterface"
)

func main() {
	batchSize := flag.Int64("bs", 1024*1024, "size of chunks (int bytes) to which files will be divided for processing")
	recursive := flag.Bool("rec", false, "Treat input path as dir and recursively process it")
	decode := flag.Bool("dc", false, "Decode input path as archive")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Error: no target path specified")
	}
	if len(args) > 1 {
		log.Fatal("Error: there can only be one target path")
	}

	absPath, err := filepath.Abs(args[0])
	if err != nil {
		log.Fatal(err)
	}
	params := appInterface.CliParams{
		BatchSize:  *batchSize,
		Recursive:  *recursive,
		Decode:     *decode,
		TargetPath: absPath,
	}

	return_code := appInterface.ExecuteApp(params)

	if return_code != 0 {
		log.Printf("Execution failed, app returned error code: %d", return_code)
	} else {
		log.Printf("Execution was successful!")
	}

}
