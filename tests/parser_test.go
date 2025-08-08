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

// TestParseOperations verifies parsing of a small subset of COBOL syntax such as
// DATA DIVISION declarations and arithmetic verbs with identifiers.
func TestParseOperations(t *testing.T) {
	src := `DATA DIVISION.
WORKING-STORAGE SECTION.
01 A.
01 B.
01 C.
PROCEDURE DIVISION.
DISPLAY "START".
MOVE A TO B.
ADD A TO B.
SUBTRACT A FROM B.
MULTIPLY A BY B.
DIVIDE A INTO B.
COMPUTE A = B + C.
INITIALIZE A B.
DISPLAY A.
STOP RUN.`

	prog, err := parser.Parse(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	if len(prog.SymbolTable) != 3 {
		t.Fatalf("expected 3 data items, got %d", len(prog.SymbolTable))
	}

	if len(prog.Statements) != 10 {
		t.Fatalf("expected 10 statements, got %d", len(prog.Statements))
	}

	if d, ok := prog.Statements[0].(*parser.Display); !ok || d.Value != "START" {
		t.Fatalf("expected first DISPLAY 'START'")
	}

	if mv, ok := prog.Statements[1].(*parser.Move); !ok || len(mv.Operands) != 2 {
		t.Fatalf("expected MOVE with two operands")
	}

	if ad, ok := prog.Statements[2].(*parser.Add); !ok || len(ad.Operands) != 2 {
		t.Fatalf("expected ADD with two operands")
	}

	if sb, ok := prog.Statements[3].(*parser.Subtract); !ok || len(sb.Operands) != 2 {
		t.Fatalf("expected SUBTRACT with two operands")
	}

	if ml, ok := prog.Statements[4].(*parser.Multiply); !ok || len(ml.Operands) != 2 {
		t.Fatalf("expected MULTIPLY with two operands")
	}

	if dv, ok := prog.Statements[5].(*parser.Divide); !ok || len(dv.Operands) != 2 {
		t.Fatalf("expected DIVIDE with two operands")
	}

	if cp, ok := prog.Statements[6].(*parser.Compute); !ok || len(cp.Operands) != 3 {
		t.Fatalf("expected COMPUTE with three operands")
	}

	if init, ok := prog.Statements[7].(*parser.Initialize); !ok || len(init.Operands) != 2 {
		t.Fatalf("expected INITIALIZE with two operands")
	}

	if d, ok := prog.Statements[8].(*parser.Display); !ok || d.Identifier != "A" {
		t.Fatalf("expected DISPLAY of identifier A")
	}

	if _, ok := prog.Statements[9].(*parser.Stop); !ok {
		t.Fatalf("expected STOP statement")
	}
}
