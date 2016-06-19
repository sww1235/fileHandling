package fileHandling

import (
	"log"
	"os"
)

func CreateFileHandle(path string) *os.File {
	f, err := os.Open("")
	if err != nil {
		log.Panicf("Failed opening file %s", path)
		log.Panicln(err)
	}
	return f
}

func CloseFileHandle(f *os.File) {
	log.Printf("Closing file %s", f.Name())
	f.Close()
}
