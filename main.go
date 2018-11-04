package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// HackAssembler behaviors
type HackAssembler interface {
	generateSymbols()
}

//Iterable defines an interface that can be iterated
type Iterable interface {
	hasNext() bool
	next() string
}

//FileIterable Iterable that is file
type FileIterable = Iterable

//FileReader is an implementation of a FileIterable
type FileReader struct {
	fileName string         //name of the file being read
	scanner  *bufio.Scanner //scanner to read the file line by line
	line     *string
}

func (f *FileReader) hasNext() bool {
	if f.line != nil {
		return true
	}
	return false
}

//StrPtr returns a pointer to a string
func StrPtr(s string) *string {
	return &s
}

func (f *FileReader) next() string {
	str := f.line
	f.line = nil
	if f.scanner.Scan() {
		f.line = StrPtr(f.scanner.Text())
	}
	if err := f.scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
		f.line = nil
	}
	return *str
}

//NewFileReader returns a pointer to a fileReader object
func NewFileReader(scanner *bufio.Scanner, fileName string) *FileReader {
	if scanner == nil {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	}

	var line *string
	if scanner.Scan() {
		line = StrPtr(scanner.Text())
	}

	fileReader := &FileReader{
		fileName: fileName,
		scanner:  scanner,
		line:     line,
	}

	return fileReader
}

func readLineAndPerformAction(itr FileIterable) {
	for itr.hasNext() {
		str := itr.next()
		fmt.Println(str)
	}
}

func main() {
	//Read in the file
	fileReader := NewFileReader(nil, "add/Add.asm")
	readLineAndPerformAction(fileReader)
}
