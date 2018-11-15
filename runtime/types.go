package runtime

import (
	"strconv"

	"github.com/rroe/novus/object"
)

func TypeFloat(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}
	if (args[0].Type() == object.INTEGER_OBJ || args[0].Type() == object.STRING_OBJ) == false {
		return MakeError("Argument to `float` must be INT or STRING, got %s",
			args[0].Type())
	}

	// Can only be two things at this point - int or string
	if args[0].Type() == object.INTEGER_OBJ {
		val := args[0].(*object.Integer)
		return &object.Float{Value: float64(val.Value)}
	}
	val := args[0].(*object.String)
	value, err := strconv.ParseFloat(val.Value, 64)
	if err != nil {
		return MakeError("could not parse '%q' as float", val.Value)
	}
	return &object.Float{Value: value}
}

func TypeString(args ...object.Object) object.Object {
	if len(args) != 1 {
		return MakeError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}

	return &object.String{Value: args[0].Inspect()}
}
