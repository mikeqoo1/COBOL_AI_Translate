package tests

import (
	"os"
	"testing"

	"cobol/ir"
	"cobol/parser"
)

func TestIRConversion(t *testing.T) {
	src, err := os.ReadFile("../examples/hello_world.cob")
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	prog, err := parser.Parse(string(src))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	irp := ir.FromAST(prog)
	if len(irp.Ops) != 2 {
		t.Fatalf("expected 2 ops, got %d", len(irp.Ops))
	}
	if d, ok := irp.Ops[0].(*ir.Display); !ok || d.Value != "HELLO WORLD" {
		t.Fatalf("expected DISPLAY 'HELLO WORLD'")
	}
	if _, ok := irp.Ops[1].(*ir.Stop); !ok {
		t.Fatalf("expected STOP op")
	}
}

func TestIRVerbs(t *testing.T) {
	src, err := os.ReadFile("../examples/verbs.cob")
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	prog, err := parser.Parse(string(src))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	irp := ir.FromAST(prog)
	if len(irp.Ops) != 12 {
		t.Fatalf("expected 12 ops, got %d", len(irp.Ops))
	}
	if d, ok := irp.Ops[0].(*ir.Display); !ok || d.Value != "START" || !d.IsLiteral {
		t.Fatalf("expected DISPLAY literal 'START'")
	}
	if m, ok := irp.Ops[1].(*ir.Move); !ok || m.From != "5" || m.To != "A" {
		t.Fatalf("expected MOVE 5 TO A")
	}
	if a, ok := irp.Ops[2].(*ir.Add); !ok || a.Value != "3" || a.To != "A" {
		t.Fatalf("expected ADD 3 TO A")
	}
	if s, ok := irp.Ops[3].(*ir.Subtract); !ok || s.Value != "2" || s.From != "A" {
		t.Fatalf("expected SUBTRACT 2 FROM A")
	}
	if mu, ok := irp.Ops[4].(*ir.Multiply); !ok || mu.Variable != "A" || mu.Value != "2" {
		t.Fatalf("expected MULTIPLY A BY 2")
	}
	if d, ok := irp.Ops[5].(*ir.Divide); !ok || d.Variable != "A" || d.Value != "3" {
		t.Fatalf("expected DIVIDE A BY 3")
	}
	if c, ok := irp.Ops[6].(*ir.Compute); !ok || c.Target != "B" || c.Expr != "A + 1" {
		t.Fatalf("expected COMPUTE B = A + 1")
	}
	if i, ok := irp.Ops[7].(*ir.Initialize); !ok || i.Variable != "C" {
		t.Fatalf("expected INITIALIZE C")
	}
	if d, ok := irp.Ops[8].(*ir.Display); !ok || d.Value != "A" || d.IsLiteral {
		t.Fatalf("expected DISPLAY A")
	}
	if d, ok := irp.Ops[9].(*ir.Display); !ok || d.Value != "B" || d.IsLiteral {
		t.Fatalf("expected DISPLAY B")
	}
	if d, ok := irp.Ops[10].(*ir.Display); !ok || d.Value != "C" || d.IsLiteral {
		t.Fatalf("expected DISPLAY C")
	}
	if _, ok := irp.Ops[11].(*ir.Stop); !ok {
		t.Fatalf("expected STOP op")
	}
}
