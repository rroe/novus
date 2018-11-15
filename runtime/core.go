package runtime

import (
	"github.com/rroe/novus/object"
)

func Range(args ...object.Object) object.Object {

	if len(args) != 2 {
		return MakeError("Wrong number of arguments given: got=%d, want=2",
			len(args))
	}

	if args[0].Type() != object.INTEGER_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return MakeError("Arguments to `range` must both be INTEGER, got %s and %s",
			args[0].Type(), args[1].Type())
	}

	bot := args[0].(*object.Integer)
	top := args[1].(*object.Integer)

	outArr := object.Array{Elements: []object.Object{}}
	if bot.Value > top.Value {
		return &outArr
	}

	for i := bot.Value; i < top.Value; i++ {
		outArr.Elements = append(outArr.Elements, &object.Integer{Value: i})
	}

	return &outArr
}
