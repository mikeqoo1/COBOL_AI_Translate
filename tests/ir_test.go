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
