package runtime

import (
	"strconv"

	"github.com/rroe/novus/object"
)

func TypeFloat(args ...object.Object) object.Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments given: got=%d, want=1",
			len(args))
	}
	if (args[0].Type() == object.INTEGER_OBJ || args[0].Type() == object.STRING_OBJ) == false {
		return NewError("Argument to `float` must be INT or STRING, got %s",
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
		return NewError("could not parse '%q' as float", val.Value)
	}
	return &object.Float{Value: value}
}
