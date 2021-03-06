package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/anujva/nand2tetris/assembler"
)

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
func NewFileReader(fileName string, scanner ...*bufio.Scanner) *FileReader {
	if scanner == nil {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		scanner = make([]*bufio.Scanner, 1)
		scanner[0] = bufio.NewScanner(file)
	}

	var line *string
	if scanner[0].Scan() {
		line = StrPtr(scanner[0].Text())
	}

	fileReader := &FileReader{
		fileName: fileName,
		scanner:  scanner[0],
		line:     line,
	}

	return fileReader
}

func readLineAndPerformAction(
	itr FileIterable,
	a *assembler.HackAssembler,
) {
	for itr.hasNext() {
		str := itr.next()
		a.AssembleFile(str)
	}
}

func readLineAndPerformActionFirstPass(
	itr FileIterable,
	a *assembler.HackAssembler,
) {
	for itr.hasNext() {
		str := itr.next()
		a.FirstPass(str)
	}
}

func readLineAndPerformActionSecondPass(
	itr FileIterable,
	a *assembler.HackAssembler,
) {
	for itr.hasNext() {
		str := itr.next()
		a.SecondPass(str)
	}
}

func main() {
	//Read in the file
	fileReader := NewFileReader(filepath.Clean(os.Args[1]))
	a := assembler.New()
	readLineAndPerformActionFirstPass(fileReader, a)
	fileReader = NewFileReader(filepath.Clean(os.Args[1]))
	readLineAndPerformActionSecondPass(fileReader, a)
	fileReader = NewFileReader(filepath.Clean(os.Args[1]))
	readLineAndPerformAction(fileReader, a)
}
