package runtime

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rroe/novus/object"
)

func Print(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return NULL
}

func Input(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}

	prompt := args[0].Inspect()
	inputStream := os.Stdin
	scanner := bufio.NewScanner(inputStream)

	fmt.Print(prompt)
	scanned := scanner.Scan()
	if !scanned {
		return &object.String{Value: ""}
	}

	return &object.String{Value: scanner.Text()}
}
