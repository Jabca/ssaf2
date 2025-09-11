package appInterface

import (
	"log"
	"os"
)

type CliParams struct {
	BatchSize  int64
	Recursive  bool
	Decode     bool
	TargetPath string
}

func ExecuteApp(params CliParams) int {
	// check whether path exists or not
	info, err := os.Stat(params.TargetPath)
	if os.IsNotExist(err) {
		log.Fatalf("%s doesn't exist", params.TargetPath)
		// check for isDir/recursive mismatch
	} else if !params.Recursive && info.IsDir() {
		log.Fatalf("%s is dir, used without 'recursive' flag", params.TargetPath)
	} else if params.Recursive && !info.IsDir() {
		log.Fatalf("%s is file, while 'recursive' file is used", params.TargetPath)
	}

	return 0
}
