package tests

import (
	"os"
	"testing"

	"cobol/parser"
)

func TestParseHelloWorld(t *testing.T) {
	src, err := os.ReadFile("../examples/hello_world.cob")
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	prog, err := parser.Parse(string(src))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(prog.Statements) != 2 {
		t.Fatalf("expected 2 statements, got %d", len(prog.Statements))
	}
	if d, ok := prog.Statements[0].(*parser.Display); !ok || d.Value != "HELLO WORLD" {
		t.Fatalf("expected DISPLAY 'HELLO WORLD'")
	}
	if _, ok := prog.Statements[1].(*parser.Stop); !ok {
		t.Fatalf("expected STOP statement")
	}
}
