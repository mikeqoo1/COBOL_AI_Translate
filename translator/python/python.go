package python

import (
	"cobol/ir"
	"strings"
)

func Translate(p *ir.Program) string {
	var b strings.Builder
	for _, op := range p.Ops {
		switch o := op.(type) {
		case *ir.Display:
			if o.IsLiteral {
				b.WriteString("print(\"")
				b.WriteString(o.Value)
				b.WriteString("\")\n")
			} else {
				b.WriteString("print(")
				b.WriteString(o.Value)
				b.WriteString(")\n")
			}
		case *ir.Move:
			b.WriteString(o.To)
			b.WriteString(" = ")
			b.WriteString(o.From)
			b.WriteString("\n")
		case *ir.Add:
			b.WriteString(o.To)
			b.WriteString(" = ")
			b.WriteString(o.To)
			b.WriteString(" + ")
			b.WriteString(o.Value)
			b.WriteString("\n")
		case *ir.Subtract:
			b.WriteString(o.From)
			b.WriteString(" = ")
			b.WriteString(o.From)
			b.WriteString(" - ")
			b.WriteString(o.Value)
			b.WriteString("\n")
		case *ir.Multiply:
			b.WriteString(o.Variable)
			b.WriteString(" = ")
			b.WriteString(o.Variable)
			b.WriteString(" * ")
			b.WriteString(o.Value)
			b.WriteString("\n")
		case *ir.Divide:
			b.WriteString(o.Variable)
			b.WriteString(" = ")
			b.WriteString(o.Variable)
			b.WriteString(" / ")
			b.WriteString(o.Value)
			b.WriteString("\n")
		case *ir.Compute:
			b.WriteString(o.Target)
			b.WriteString(" = ")
			b.WriteString(o.Expr)
			b.WriteString("\n")
		case *ir.Initialize:
			b.WriteString(o.Variable)
			b.WriteString(" = 0\n")
		}
	}
	return b.String()
}
