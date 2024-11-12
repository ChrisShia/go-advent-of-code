package utils

import (
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func PrintFile(path string) {
	dat, err := os.ReadFile(path)
	check(err)
	fmt.Print(string(dat))
}

func ReadFile(path string) []byte {
	dat, err := os.ReadFile(path)
	check(err)
	return dat
}

func ReadBytesFromFile(path string, noOfBytes int) {
	f, err := OpenFile(path)
	bytes := make([]byte, noOfBytes)
	n1, err := f.Read(bytes)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(bytes[:n1]))
}

func OpenFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	check(err)
	return f, err
}

func OpenFileLogFatal(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func CloseFile(file *os.File) {
	func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
}
