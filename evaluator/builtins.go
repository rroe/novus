package evaluator

import (
	"github.com/rroe/novus/object"
	"github.com/rroe/novus/runtime"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: runtime.ListLength,
	},
	"range": &object.Builtin{
		Fn: runtime.Range,
	},
	"print": &object.Builtin{
		Fn: runtime.Print,
	},
	"input": &object.Builtin{
		Fn: runtime.Input,
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
	"str_join": &object.Builtin{
		Fn: runtime.StringJoin,
	},
	"file_write_lines": &object.Builtin{
		Fn: runtime.FileWriteLines,
	},
	"file_read_lines": &object.Builtin{
		Fn: runtime.FileReadLines,
	},
}
