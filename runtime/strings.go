package runtime

import (
	"strings"

	"github.com/rroe/novus/object"
)

func StringJoin(args ...object.Object) object.Object {
	if len(args) != 2 {
		return MakeError("Wrong number of arguments given: got=%d, want=2",
			len(args))
	}
	if (args[0].Type() == object.ARRAY_OBJ) == false {
		return MakeError("First argument to `str_join` must be ARRAY got %s",
			args[0].Type())
	}

	strs := args[0].(*object.Array)
	delim := args[1].Inspect()

	numElems := len(strs.Elements)

	var out string

	for k, v := range strs.Elements {
		val := v.Inspect()
		out = out + val
		if k != (numElems - 1) {
			out = out + delim
		}
	}

	return &object.String{Value: out}
}

func StringLetters(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}

	out := []object.Object{}
	str := args[0].Inspect()

	for _, let := range str {
		letter := &object.String{Value: string(let)}
		out = append(out, letter)
	}

	return &object.Array{Elements: out}
}

func StringContains(args ...object.Object) object.Object {
	if len(args) != 2 {
		return MakeError("Wrong number of arguments given: got=%d, want=2",
			len(args))
	}

	str := args[0].Inspect()
	subStr := args[1].Inspect()

	if strings.Contains(str, subStr) {
		return TRUE
	}
	return FALSE
}

func StringSplit(args ...object.Object) object.Object {
	if len(args) != 2 {
		return MakeError("Wrong number of arguments given: got=%d, want=2",
			len(args))
	}

	out := []object.Object{}
	str := args[0].Inspect()
	delim := args[1].Inspect()

	for _, let := range strings.Split(str, delim) {
		chunk := &object.String{Value: string(let)}
		out = append(out, chunk)
	}

	return &object.Array{Elements: out}
}

func StringLower(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}

	out := strings.ToLower(args[0].Inspect())

	return &object.String{Value: out}
}

func StringUpper(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}

	out := strings.ToUpper(args[0].Inspect())

	return &object.String{Value: out}
}
