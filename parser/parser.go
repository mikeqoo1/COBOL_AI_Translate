package parser

import (
	"bufio"
	"strconv"
	"strings"
)

type Program struct {
	Statements []Statement
}

type Statement interface{}

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

func Parse(src string) (*Program, error) {
	p := &Program{}
	scanner := bufio.NewScanner(strings.NewReader(src))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.TrimSuffix(line, ".")
		lineUpper := strings.ToUpper(line)
		switch {
		case strings.HasPrefix(lineUpper, "DISPLAY "):
			rest := strings.TrimSpace(line[8:])
			if strings.HasPrefix(rest, "\"") && strings.HasSuffix(rest, "\"") {
				value := rest[1 : len(rest)-1]
				p.Statements = append(p.Statements, &Display{Value: value, IsLiteral: true})
			} else if rest != "" {
				p.Statements = append(p.Statements, &Display{Value: rest})
			}
		case strings.HasPrefix(lineUpper, "MOVE "):
			parts := strings.Fields(line)
			if len(parts) >= 4 && strings.ToUpper(parts[2]) == "TO" {
				from := parts[1]
				to := parts[3]
				_, err := strconv.Atoi(from)
				p.Statements = append(p.Statements, &Move{From: from, FromLiteral: err == nil, To: to})
			}
		case strings.HasPrefix(lineUpper, "ADD "):
			parts := strings.Fields(line)
			if len(parts) >= 4 && strings.ToUpper(parts[2]) == "TO" {
				val := parts[1]
				to := parts[3]
				_, err := strconv.Atoi(val)
				p.Statements = append(p.Statements, &Add{Value: val, Literal: err == nil, To: to})
			}
		case strings.HasPrefix(lineUpper, "SUBTRACT "):
			parts := strings.Fields(line)
			if len(parts) >= 4 && strings.ToUpper(parts[2]) == "FROM" {
				val := parts[1]
				from := parts[3]
				_, err := strconv.Atoi(val)
				p.Statements = append(p.Statements, &Subtract{Value: val, Literal: err == nil, From: from})
			}
		case strings.HasPrefix(lineUpper, "MULTIPLY "):
			parts := strings.Fields(line)
			if len(parts) >= 4 && strings.ToUpper(parts[2]) == "BY" {
				variable := parts[1]
				val := parts[3]
				_, err := strconv.Atoi(val)
				p.Statements = append(p.Statements, &Multiply{Variable: variable, Value: val, Literal: err == nil})
			}
		case strings.HasPrefix(lineUpper, "DIVIDE "):
			parts := strings.Fields(line)
			if len(parts) >= 4 && strings.ToUpper(parts[2]) == "BY" {
				variable := parts[1]
				val := parts[3]
				_, err := strconv.Atoi(val)
				p.Statements = append(p.Statements, &Divide{Variable: variable, Value: val, Literal: err == nil})
			}
		case strings.HasPrefix(lineUpper, "COMPUTE "):
			rest := strings.TrimSpace(line[8:])
			parts := strings.SplitN(rest, "=", 2)
			if len(parts) == 2 {
				target := strings.TrimSpace(parts[0])
				expr := strings.TrimSpace(parts[1])
				p.Statements = append(p.Statements, &Compute{Target: target, Expr: expr})
			}
		case strings.HasPrefix(lineUpper, "INITIALIZE "):
			variable := strings.TrimSpace(line[10:])
			if variable != "" {
				p.Statements = append(p.Statements, &Initialize{Variable: variable})
			}
		case lineUpper == "STOP RUN":
			p.Statements = append(p.Statements, &Stop{})
		}
	}
	return p, scanner.Err()
}
