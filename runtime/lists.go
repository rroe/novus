package runtime

import (
	"github.com/rroe/novus/object"
)

func ListLength(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Hash:
		return &object.Integer{Value: int64(len(arg.Pairs))}
	default:
		return MakeError("Argument to `len` not supported, got %s",
			args[0].Type())
	}
}

func ListFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return MakeError("Argument to `first` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func ListLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return MakeError("Argument to `last` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}

func ListRest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return MakeError("Argument to `rest` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func ListPush(args ...object.Object) object.Object {
	if len(args) != 2 {
		return MakeError("Wrong number of arguments given: got=%d, want=2",
			len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return MakeError("Argument to `push` must be ARRAY, got %s",
			args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}
