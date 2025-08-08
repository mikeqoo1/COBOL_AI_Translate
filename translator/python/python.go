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
			b.WriteString("print(\"")
			b.WriteString(o.Value)
			b.WriteString("\")\n")
		}
	}
	return b.String()
}
