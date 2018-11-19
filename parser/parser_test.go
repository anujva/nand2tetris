package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestRemoveInsideWhiteSpaces(t *testing.T) {
	str1 := removeWhiteSpaces("abc d")
	if strings.Compare(str1, "abcd") != 0 {
		t.Fail()
	}
	str2 := removeWhiteSpaces("   a   b    c   d    ")
	if strings.Compare(str2, "abcd") != 0 {
		t.Fail()
	}
}

var code = `
// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/06/rect/Rect.asm

// Draws a rectangle at the top-left corner of the screen.
// The rectangle is 16 pixels wide and R0 pixels high.

   @0 //This is another whitespace that will need to be removed
   D=M
   @INFINITE_LOOP
   D;JLE 
   @counter
   M=D
   @SCREEN
   D=A
   @address
   M=D
(LOOP)
   @address
   A=M
   M=-1
   @address
   D=M
   @32
   D=D+A
   @address
   M=D
   @counter
   MD=M-1
   @LOOP
   D;JGT
(INFINITE_LOOP)
   @INFINITE_LOOP
   0;JMP
`

func TestPrintLinesAfterRemovingWhiteSpaces(t *testing.T) {
	t.Skip()
	codeLines := strings.Split(code, "\n")
	fmt.Println("The line with whitespaces removed:")
	for _, codeLine := range codeLines {
		l := removeWhiteSpaces(codeLine)
		if len(l) != 0 {
			fmt.Println(l)
		}
	}
}

func TestParser(t *testing.T) {
	parser := New()
	tokens := parser.Parse("@123")

	for _, token := range tokens {
		fmt.Printf("%+v\n", token)
	}
}

func TestParserCInstruction(t *testing.T) {
	parser := New()
	tokens := parser.Parse("CMP;JMP")

	for _, token := range tokens {
		fmt.Printf("%+v\n", token)
	}
}

func TestSplitCInstruction(t *testing.T) {
	val := splitCInstruction("CMP")
	fmt.Printf("%+v\n", val)
}

func TestSplitCInstructionAnotherWay(t *testing.T) {
	val := splitCInstruction("CMP;JMP")
	fmt.Printf("%+v\n", val)
}

func TestSplitCInstructionOneMoreWay(t *testing.T) {
	val := splitCInstruction("DEST=CMP")
	fmt.Printf("%+v\n", val)
}

func TestSplitCInstructionWithTrailing(t *testing.T) {
	val := splitCInstruction("DEST=CMP;")
	fmt.Printf("%+v\n", val)
}

func TestSplitCInstructionFull(t *testing.T) {
	val := splitCInstruction("DEST=CMP;JMP")
	fmt.Printf("%+v\n", val)
}
