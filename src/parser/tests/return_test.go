package tests

import (
	"testing"
	"lexer"
	"parser"
	"ast"
)

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 1234;
	`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParserProgram()

	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParserProgram() returned nil :( ")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Return statement doesn't contain 3 statements. got=%d", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement) // type assertion to make sure we have a return statement

		if !ok {
			t.Errorf("statement not *ast.ReturnStatement. got=%T", statement)
			continue
		}

		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return'. got=%q", returnStatement.TokenLiteral())
		}
	}
}