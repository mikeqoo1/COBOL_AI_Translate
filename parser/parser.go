package parser

import (
	"bufio"
	"strings"
)

type Program struct {
	Statements []Statement
}

type Statement interface{}

type Display struct {
	Value string
}

type Stop struct{}

func Parse(src string) (*Program, error) {
	p := &Program{}
	scanner := bufio.NewScanner(strings.NewReader(src))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.TrimSuffix(line, ".")
		lineUpper := strings.ToUpper(line)
		switch {
		case strings.HasPrefix(lineUpper, "DISPLAY "):
			start := strings.Index(line, "\"")
			end := strings.LastIndex(line, "\"")
			if start >= 0 && end > start {
				value := line[start+1 : end]
				p.Statements = append(p.Statements, &Display{Value: value})
			}
		case lineUpper == "STOP RUN":
			p.Statements = append(p.Statements, &Stop{})
		}
	}
	return p, scanner.Err()
}
