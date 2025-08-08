package ir

import "cobol/parser"

type Program struct {
	Ops []Op
}

type Op interface{}

type Display struct {
	Value string
}

type Stop struct{}

func FromAST(p *parser.Program) *Program {
	irp := &Program{}
	for _, stmt := range p.Statements {
		switch s := stmt.(type) {
		case *parser.Display:
			irp.Ops = append(irp.Ops, &Display{Value: s.Value})
		case *parser.Stop:
			irp.Ops = append(irp.Ops, &Stop{})
		}
	}
	return irp
}
