package appInterface

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"ssaf2/internal/archive/nodes"
	"syscall"
)

// Encode file and store it in tmp filePath.
// Note that data pointers are not set, as they will be changed during archive formation
func TmpEncodeFile(
	fp string,
	batchSize uint64,
	tmpFp string,
	readBuffSize uint64) {

	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(nil)
	}
	defer file.Close()

	nodeDesc := computeNodeDescriptor(fp, readBuffSize, batchSize)
	bytes, err := nodeDesc.Encode()
	if err != nil {
		log.Fatal(bytes)
	}
	file.Write(bytes)

	encNode := computeEncodingNode(fp, readBuffSize)
	bytes, err = encNode.Encode()
	if err != nil {
		log.Fatal(err)
	}
	file.Write(bytes)

}

func computeEncodingNode(fp string, readBuffSize uint64) nodes.EncodingNode {
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, readBuffSize)
	byteCount := make(map[byte]uint64)

	for {
		_, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, val := range buffer {
			byteCount[val] += 1
		}
	}

	values := make([]byte, 256)
	var b uint16 = 0
	for b < 256 {
		values[b] = byte(b)
		b++
	}

	// now sort this byte values based on there popularity in this file
	// there index is there rank in file (ascending)
	// value = byte value, index = rank
	sort.Slice(values, func(i, j int) bool {
		val1 := values[i]
		val2 := values[j]
		return byteCount[val1] < byteCount[val2] || val1 < val2
	})

	// now let's transfer it into ranks array
	// value = rank, index = byte value
	ranks := make([]byte, 256)
	for index, val := range values {
		ranks[val] = byte(index)
	}

	return nodes.EncodingNode{
		Ranks: [256]byte(ranks),
	}

}

func computeFileHash(fp string, readBuffSize uint64) ([]byte, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hasher := sha256.New()

	buff := make([]byte, readBuffSize)
	for {
		_, err := file.Read(buff)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		_, err = hasher.Write(buff)
		if err != nil {
			log.Fatal(err)
		}

	}

	return hasher.Sum(nil), nil
}

func computeNodeDescriptor(fp string, readBuffSize uint64, batchSize uint64) nodes.FileNode {
	info, err := os.Stat(fp)
	if err != nil {
		log.Fatal(err)
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalf("Couldn't get stat for %s", fp)
	}

	hash, err := computeFileHash(fp, readBuffSize)
	if err != nil {
		log.Fatalf("Couldn't compute sha256 for %s", fp)
	}

	batchCount := uint64(info.Size()) / batchSize
	if info.Size()%int64(batchSize) > 0 {
		batchCount += 1
	}

	return nodes.FileNode{
		ModTime:        info.ModTime(),
		FileSize:       uint64(info.Size()),
		CheckSum:       [32]byte(hash),
		ID:             stat.Ino,
		Name:           filepath.Base(fp),
		IsDir:          info.IsDir(),
		DataOffset:     -1,
		DataNodesCount: batchCount,
	}

}
