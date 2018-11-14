package runtime

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rroe/novus/object"
)

func FileWriteLines(args ...object.Object) object.Object {
	if len(args) != 2 {
		return MakeError("Wrong number of arguments given: got=%d, want=2",
			len(args))
	}

	if args[1].Type() != object.ARRAY_OBJ {
		return MakeError("Second argument for 'file_write_lines' must be type ARRAY, got %s", args[1].Type())
	}

	path := args[0].Inspect()
	vals := args[1].(*object.Array)
	toWrite := []string{}

	for _, e := range vals.Elements {
		toWrite = append(toWrite, e.Inspect())
	}

	err := writeLines(toWrite, path)

	if err != nil {
		return MakeError("Unable to write to file %s", path)
	}

	return &object.Boolean{Value: true}
}

func FileReadLines(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}

	path := args[0].Inspect()

	lines, err := readLines(path)

	if err != nil {
		return MakeError("Could not read from file '%s'", path)
	}

	contents := []object.Object{}

	for _, e := range lines {
		contents = append(contents, &object.String{Value: e})
	}

	return &object.Array{Elements: contents}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
