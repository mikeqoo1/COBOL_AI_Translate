package java

import (
	"cobol/ir"
	"strings"
)

func Translate(p *ir.Program) string {
	var b strings.Builder
	b.WriteString("public class Main {\n")
	b.WriteString("    public static void main(String[] args) {\n")
	declared := map[string]bool{}
	for _, op := range p.Ops {
		switch o := op.(type) {
		case *ir.Display:
			if o.IsLiteral {
				b.WriteString("        System.out.println(\"")
				b.WriteString(o.Value)
				b.WriteString("\");\n")
			} else {
				b.WriteString("        System.out.println(")
				b.WriteString(o.Value)
				b.WriteString(");\n")
			}
		case *ir.Move:
			if !declared[o.To] {
				b.WriteString("        int ")
				b.WriteString(o.To)
				declared[o.To] = true
			} else {
				b.WriteString("        ")
				b.WriteString(o.To)
			}
			b.WriteString(" = ")
			b.WriteString(o.From)
			b.WriteString(";\n")
		case *ir.Add:
			if !declared[o.To] {
				b.WriteString("        int ")
				b.WriteString(o.To)
				declared[o.To] = true
			} else {
				b.WriteString("        ")
				b.WriteString(o.To)
			}
			b.WriteString(" = ")
			b.WriteString(o.To)
			b.WriteString(" + ")
			b.WriteString(o.Value)
			b.WriteString(";\n")
		case *ir.Subtract:
			if !declared[o.From] {
				b.WriteString("        int ")
				b.WriteString(o.From)
				declared[o.From] = true
			} else {
				b.WriteString("        ")
				b.WriteString(o.From)
			}
			b.WriteString(" = ")
			b.WriteString(o.From)
			b.WriteString(" - ")
			b.WriteString(o.Value)
			b.WriteString(";\n")
		case *ir.Multiply:
			if !declared[o.Variable] {
				b.WriteString("        int ")
				b.WriteString(o.Variable)
				declared[o.Variable] = true
			} else {
				b.WriteString("        ")
				b.WriteString(o.Variable)
			}
			b.WriteString(" = ")
			b.WriteString(o.Variable)
			b.WriteString(" * ")
			b.WriteString(o.Value)
			b.WriteString(";\n")
		case *ir.Divide:
			if !declared[o.Variable] {
				b.WriteString("        int ")
				b.WriteString(o.Variable)
				declared[o.Variable] = true
			} else {
				b.WriteString("        ")
				b.WriteString(o.Variable)
			}
			b.WriteString(" = ")
			b.WriteString(o.Variable)
			b.WriteString(" / ")
			b.WriteString(o.Value)
			b.WriteString(";\n")
		case *ir.Compute:
			if !declared[o.Target] {
				b.WriteString("        int ")
				b.WriteString(o.Target)
				declared[o.Target] = true
			} else {
				b.WriteString("        ")
				b.WriteString(o.Target)
			}
			b.WriteString(" = ")
			b.WriteString(o.Expr)
			b.WriteString(";\n")
		case *ir.Initialize:
			if !declared[o.Variable] {
				b.WriteString("        int ")
				b.WriteString(o.Variable)
				declared[o.Variable] = true
			} else {
				b.WriteString("        ")
				b.WriteString(o.Variable)
			}
			b.WriteString(" = 0;\n")
		}
	}
	b.WriteString("    }\n}\n")
	return b.String()
}
