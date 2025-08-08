package java

import (
	"cobol/ir"
	"strings"
)

func Translate(p *ir.Program) string {
	var b strings.Builder
	b.WriteString("public class Main {\n")
	b.WriteString("    public static void main(String[] args) {\n")
	for _, op := range p.Ops {
		switch o := op.(type) {
		case *ir.Display:
			b.WriteString("        System.out.println(\"")
			b.WriteString(o.Value)
			b.WriteString("\");\n")
		}
	}
	b.WriteString("    }\n}\n")
	return b.String()
}
