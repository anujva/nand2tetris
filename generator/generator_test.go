package generator

import (
	"fmt"
	"testing"

	token "github.com/anujva/nand2tetris/token"
)

//Will write some tests for the package here

func TestCodeGenerator(t *testing.T) {
	cg := New()
	tok := token.Token{
		Type: token.COMP,
		Val:  "D+M",
	}
	str, _ := cg.translateToken(tok)
	fmt.Println("The translated token value: ", str)
}

func TestCodeGeneratorDest(t *testing.T) {
	cg := New()
	tok := token.Token{
		Type: token.DEST,
		Val:  "A",
	}
	str, _ := cg.translateToken(tok)
	fmt.Println("The DEST translated token: ", str)
}

func TestCodeGeneratorJMP(t *testing.T) {
	cg := New()
	tok := token.Token{
		Type: token.JUMP,
		Val:  "null",
	}
	str, _ := cg.translateToken(tok)
	fmt.Println("The JMP translated token: ", str)
}

func TestCodeGeneratorJMPSecond(t *testing.T) {
	cg := New()
	tok := token.Token{
		Type: token.JUMP,
		Val:  "JGT",
	}
	str, _ := cg.translateToken(tok)
	fmt.Println("The JMP translated token: ", str)
}

func TestCodeGeneratorAddress(t *testing.T) {
	cg := New()
	tok := token.Token{
		Type: token.ADDRESS,
		Val:  "15",
	}
	str, _ := cg.translateToken(tok)
	fmt.Println("The Address tranlated token: ", str)
}
