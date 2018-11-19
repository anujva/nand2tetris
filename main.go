package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
		tkns := a.Parser.Parse(str)
		finalString := ""
		if len(tkns) == 3 {
			//This is a c instruction
			finalString1, _ := a.Code.TranslateToken(tkns[1])
			finalString2, _ := a.Code.TranslateToken(tkns[0])
			finalString3, _ := a.Code.TranslateToken(tkns[2])
			finalString = finalString1 + finalString2 + finalString3
		} else if len(tkns) == 1 {
			//This is a instruction
			finalString, _ = a.Code.TranslateToken(tkns[0])
		}
		if len(finalString) > 0 {
			fmt.Println(finalString)
		}
	}
}

func main() {
	//Read in the file
	fileReader := NewFileReader("/home/anuj/nand2tetris/projects/06/max/MaxL.asm")
	a := assembler.New()
	readLineAndPerformAction(fileReader, a)
}
