package tests

import (
	"os"
	"testing"

	"cobol/ir"
	"cobol/parser"
	javat "cobol/translator/java"
	pyt "cobol/translator/python"
)

func TestPythonTranslation(t *testing.T) {
	src, err := os.ReadFile("../examples/hello_world.cob")
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	prog, _ := parser.Parse(string(src))
	irp := ir.FromAST(prog)
	out := pyt.Translate(irp)
	expected, err := os.ReadFile("../examples/hello_world.py")
	if err != nil {
		t.Fatalf("read python example: %v", err)
	}
	if out != string(expected) {
		t.Fatalf("python translation mismatch:\nexpected:\n%q\nactual:\n%q", string(expected), out)
	}
}

func TestJavaTranslation(t *testing.T) {
	src, err := os.ReadFile("../examples/hello_world.cob")
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	prog, _ := parser.Parse(string(src))
	irp := ir.FromAST(prog)
	out := javat.Translate(irp)
	expected, err := os.ReadFile("../examples/hello_world.java")
	if err != nil {
		t.Fatalf("read java example: %v", err)
	}
	if out != string(expected) {
		t.Fatalf("java translation mismatch:\nexpected:\n%q\nactual:\n%q", string(expected), out)
	}
}

func TestPythonTranslationVerbs(t *testing.T) {
	src, err := os.ReadFile("../examples/verbs.cob")
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	prog, _ := parser.Parse(string(src))
	irp := ir.FromAST(prog)
	out := pyt.Translate(irp)
	expected, err := os.ReadFile("../examples/verbs.py")
	if err != nil {
		t.Fatalf("read python example: %v", err)
	}
	if out != string(expected) {
		t.Fatalf("python translation mismatch:\nexpected:\n%q\nactual:\n%q", string(expected), out)
	}
}

func TestJavaTranslationVerbs(t *testing.T) {
	src, err := os.ReadFile("../examples/verbs.cob")
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	prog, _ := parser.Parse(string(src))
	irp := ir.FromAST(prog)
	out := javat.Translate(irp)
	expected, err := os.ReadFile("../examples/verbs.java")
	if err != nil {
		t.Fatalf("read java example: %v", err)
	}
	if out != string(expected) {
		t.Fatalf("java translation mismatch:\nexpected:\n%q\nactual:\n%q", string(expected), out)
	}
}
