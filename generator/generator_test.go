package generator

import (
	"fmt"
	"testing"

	token "github.com/anujva/nand2tetris/token"
)

//Will write some tests for the package here

func testCodeGenerator(t *testing.T) {
	cg := New()
	tok := token.Token{
		Type: token.COMP,
		Val:  "D+M",
	}
	str := cg.translateToken(tok)
	fmt.Println("The translated token value: ", str)
}
