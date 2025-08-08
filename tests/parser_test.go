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

func TestParseVerbs(t *testing.T) {
	src, err := os.ReadFile("../examples/verbs.cob")
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	prog, err := parser.Parse(string(src))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(prog.Statements) != 12 {
		t.Fatalf("expected 12 statements, got %d", len(prog.Statements))
	}
	if d, ok := prog.Statements[0].(*parser.Display); !ok || d.Value != "START" || !d.IsLiteral {
		t.Fatalf("expected DISPLAY literal 'START'")
	}
	if m, ok := prog.Statements[1].(*parser.Move); !ok || m.From != "5" || m.To != "A" {
		t.Fatalf("expected MOVE 5 TO A")
	}
	if a, ok := prog.Statements[2].(*parser.Add); !ok || a.Value != "3" || a.To != "A" {
		t.Fatalf("expected ADD 3 TO A")
	}
	if s, ok := prog.Statements[3].(*parser.Subtract); !ok || s.Value != "2" || s.From != "A" {
		t.Fatalf("expected SUBTRACT 2 FROM A")
	}
	if mu, ok := prog.Statements[4].(*parser.Multiply); !ok || mu.Variable != "A" || mu.Value != "2" {
		t.Fatalf("expected MULTIPLY A BY 2")
	}
	if d, ok := prog.Statements[5].(*parser.Divide); !ok || d.Variable != "A" || d.Value != "3" {
		t.Fatalf("expected DIVIDE A BY 3")
	}
	if c, ok := prog.Statements[6].(*parser.Compute); !ok || c.Target != "B" || c.Expr != "A + 1" {
		t.Fatalf("expected COMPUTE B = A + 1")
	}
	if i, ok := prog.Statements[7].(*parser.Initialize); !ok || i.Variable != "C" {
		t.Fatalf("expected INITIALIZE C")
	}
	if d, ok := prog.Statements[8].(*parser.Display); !ok || d.Value != "A" || d.IsLiteral {
		t.Fatalf("expected DISPLAY A")
	}
	if d, ok := prog.Statements[9].(*parser.Display); !ok || d.Value != "B" || d.IsLiteral {
		t.Fatalf("expected DISPLAY B")
	}
	if d, ok := prog.Statements[10].(*parser.Display); !ok || d.Value != "C" || d.IsLiteral {
		t.Fatalf("expected DISPLAY C")
	}
	if _, ok := prog.Statements[11].(*parser.Stop); !ok {
		t.Fatalf("expected STOP statement")
	}
}
