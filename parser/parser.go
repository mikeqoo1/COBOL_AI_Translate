package parser

import (
	"bufio"
	"strings"
)

// Program represents a parsed COBOL program.  It contains a list of
// statements in the PROCEDURE DIVISION as well as a simple symbol table of
// working-storage data items.
type Program struct {
	Statements  []Statement
	SymbolTable map[string]*DataItem
}

// Statement is implemented by all COBOL statements that the parser
// recognizes.  The parser is intentionally minimal and only understands a
// handful of arithmetic verbs and DISPLAY/STOP.
type Statement interface{}

// DataItem represents a variable declared in the DATA DIVISION.
type DataItem struct {
	Name string
}

// Display corresponds to the DISPLAY statement.  Either Value (for string
// literals) or Identifier is set.
type Display struct {
	Value      string
	Identifier string
}

// The remaining structs represent arithmetic and move style verbs.  For the
// purposes of this project they simply record the identifiers that appear in
// the statement in the order they were encountered.
type Move struct{ Operands []string }
type Add struct{ Operands []string }
type Subtract struct{ Operands []string }
type Multiply struct{ Operands []string }
type Divide struct{ Operands []string }
type Compute struct{ Operands []string }
type Initialize struct{ Operands []string }

type Stop struct{}

// Parse performs a very small amount of COBOL parsing sufficient for unit
// tests.  It recognises DATA DIVISION/WORKING-STORAGE declarations and a small
// subset of PROCEDURE DIVISION statements containing identifiers.
func Parse(src string) (*Program, error) {
	p := &Program{SymbolTable: make(map[string]*DataItem)}
	scanner := bufio.NewScanner(strings.NewReader(src))
	inWorking := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.TrimSuffix(line, ".")
		lineUpper := strings.ToUpper(line)

		switch {
		case lineUpper == "DATA DIVISION" || lineUpper == "WORKING-STORAGE SECTION":
			inWorking = true
			continue
		case lineUpper == "PROCEDURE DIVISION":
			inWorking = false
			continue
		}

		if inWorking {
			fields := strings.Fields(lineUpper)
			if len(fields) >= 2 {
				name := fields[1]
				p.SymbolTable[name] = &DataItem{Name: name}
			}
			continue
		}

		switch {
		case strings.HasPrefix(lineUpper, "DISPLAY "):
			start := strings.Index(line, "\"")
			end := strings.LastIndex(line, "\"")
			if start >= 0 && end > start {
				value := line[start+1 : end]
				p.Statements = append(p.Statements, &Display{Value: value})
			} else {
				parts := strings.Fields(line)
				if len(parts) > 1 {
					p.Statements = append(p.Statements, &Display{Identifier: parts[1]})
				}
			}
		case strings.HasPrefix(lineUpper, "MOVE "):
			p.Statements = append(p.Statements, &Move{Operands: parseOperands(line)})
		case strings.HasPrefix(lineUpper, "ADD "):
			p.Statements = append(p.Statements, &Add{Operands: parseOperands(line)})
		case strings.HasPrefix(lineUpper, "SUBTRACT "):
			p.Statements = append(p.Statements, &Subtract{Operands: parseOperands(line)})
		case strings.HasPrefix(lineUpper, "MULTIPLY "):
			p.Statements = append(p.Statements, &Multiply{Operands: parseOperands(line)})
		case strings.HasPrefix(lineUpper, "DIVIDE "):
			p.Statements = append(p.Statements, &Divide{Operands: parseOperands(line)})
		case strings.HasPrefix(lineUpper, "COMPUTE "):
			p.Statements = append(p.Statements, &Compute{Operands: parseOperands(line)})
		case strings.HasPrefix(lineUpper, "INITIALIZE "):
			p.Statements = append(p.Statements, &Initialize{Operands: parseOperands(line)})
		case lineUpper == "STOP RUN":
			p.Statements = append(p.Statements, &Stop{})
		}
	}
	return p, scanner.Err()
}

// parseOperands extracts identifiers from a statement line by removing common
// punctuation and keywords.  It is intentionally permissive and is only meant
// for tests.
func parseOperands(line string) []string {
	tokens := strings.Fields(strings.TrimSuffix(line, "."))
	operands := []string{}
	for _, tok := range tokens[1:] { // skip verb
		upper := strings.ToUpper(tok)
		if upper == "TO" || upper == "FROM" || upper == "BY" || upper == "INTO" || upper == "GIVING" || upper == "=" {
			continue
		}
		// split on arithmetic operators
		parts := strings.FieldsFunc(tok, func(r rune) bool {
			return strings.ContainsRune("+-*/", r)
		})
		for _, p := range parts {
			if p != "" {
				operands = append(operands, p)
			}
		}
	}
	return operands
}
