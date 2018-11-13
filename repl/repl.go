package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rroe/novus/evaluator"
	"github.com/rroe/novus/lexer"
	"github.com/rroe/novus/object"
	"github.com/rroe/novus/parser"
)

const PROMPT = "$> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const NOVUS_BANNER = ` _   _                      
| \ | |                     
|  \| | _____   ___   _ ___ 
| . ' |/ _ \ \ / / | | / __|
| |\  | (_) \ V /| |_| \__ \
|_| \_|\___/ \_/  \__,_|___/
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, NOVUS_BANNER)
	io.WriteString(out, "Error encountered on parsing.\n")
	io.WriteString(out, " - Errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
