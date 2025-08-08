package ir

import "cobol/parser"

type Program struct {
	Ops []Op
}

type Op interface{}

type Display struct {
	Value     string
	IsLiteral bool
}

type Move struct {
	From        string
	FromLiteral bool
	To          string
}

type Add struct {
	Value   string
	Literal bool
	To      string
}

type Subtract struct {
	Value   string
	Literal bool
	From    string
}

type Multiply struct {
	Variable string
	Value    string
	Literal  bool
}

type Divide struct {
	Variable string
	Value    string
	Literal  bool
}

type Compute struct {
	Target string
	Expr   string
}

type Initialize struct {
	Variable string
}

type Stop struct{}

func FromAST(p *parser.Program) *Program {
	irp := &Program{}
	for _, stmt := range p.Statements {
		switch s := stmt.(type) {
		case *parser.Display:
			irp.Ops = append(irp.Ops, &Display{Value: s.Value, IsLiteral: s.IsLiteral})
		case *parser.Move:
			irp.Ops = append(irp.Ops, &Move{From: s.From, FromLiteral: s.FromLiteral, To: s.To})
		case *parser.Add:
			irp.Ops = append(irp.Ops, &Add{Value: s.Value, Literal: s.Literal, To: s.To})
		case *parser.Subtract:
			irp.Ops = append(irp.Ops, &Subtract{Value: s.Value, Literal: s.Literal, From: s.From})
		case *parser.Multiply:
			irp.Ops = append(irp.Ops, &Multiply{Variable: s.Variable, Value: s.Value, Literal: s.Literal})
		case *parser.Divide:
			irp.Ops = append(irp.Ops, &Divide{Variable: s.Variable, Value: s.Value, Literal: s.Literal})
		case *parser.Compute:
			irp.Ops = append(irp.Ops, &Compute{Target: s.Target, Expr: s.Expr})
		case *parser.Initialize:
			irp.Ops = append(irp.Ops, &Initialize{Variable: s.Variable})
		case *parser.Stop:
			irp.Ops = append(irp.Ops, &Stop{})
		}
	}
	return irp
}
