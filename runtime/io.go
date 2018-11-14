package runtime

import (
	"fmt"

	"github.com/rroe/novus/object"
)

func Print(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return NULL
}
