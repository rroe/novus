package evaluator

import (
	"github.com/rroe/novus/object"
	"github.com/rroe/novus/runtime"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: runtime.ListLength,
	},
	"print": &object.Builtin{
		Fn: runtime.Print,
	},
	"first": &object.Builtin{
		Fn: runtime.ListFirst,
	},
	"last": &object.Builtin{
		Fn: runtime.ListLast,
	},
	"rest": &object.Builtin{
		Fn: runtime.ListRest,
	},
	"push": &object.Builtin{
		Fn: runtime.ListPush,
	},
	"float": &object.Builtin{
		Fn: runtime.TypeFloat,
	},
	"string": &object.Builtin{
		Fn: runtime.TypeString,
	},
	"file_write_lines": &object.Builtin{
		Fn: runtime.FileWriteLines,
	},
	"file_read_lines": &object.Builtin{
		Fn: runtime.FileReadLines,
	},
}
