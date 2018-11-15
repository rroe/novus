package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/rroe/novus/evaluator"
	"github.com/rroe/novus/lexer"
	"github.com/rroe/novus/object"
	"github.com/rroe/novus/parser"
	"github.com/rroe/novus/repl"
)

var (
	Version string
)

func main() {
	args := os.Args[1:]

	data, err := Asset("novus_std/stdlib.novus")

	if err != nil {
		panic("Unable to load stdlib: " + err.Error())
	}

	if len(args) == 1 {
		contents, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		runFile(string(data) + string(contents))
	} else {
		fmt.Printf("Novus v%s\n", Version)
		repl.Start(os.Stdin, os.Stdout, string(data))
	}
}

func runFile(content string) {
	env := object.NewEnvironment()
	l := lexer.New(content)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		// Do nothing
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, repl.NOVUS_BANNER)
	io.WriteString(out, "Error encountered on parsing file.\n")
	io.WriteString(out, " - Errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
