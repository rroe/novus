package evaluator

import (
	"github.com/rroe/novus/object"
	"github.com/rroe/novus/runtime"
)

var builtins = map[string]*object.Builtin{
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
	"len": &object.Builtin{
		Fn: runtime.ListLength,
	},
	"range": &object.Builtin{
		Fn: runtime.Range,
	},

	"float": &object.Builtin{
		Fn: runtime.TypeFloat,
	},
	"string": &object.Builtin{
		Fn: runtime.TypeString,
	},
	"typeof": &object.Builtin{
		Fn: runtime.TypeOf,
	},

	"str_join": &object.Builtin{
		Fn: runtime.StringJoin,
	},
	"str_letters": &object.Builtin{
		Fn: runtime.StringLetters,
	},
	"str_contains": &object.Builtin{
		Fn: runtime.StringContains,
	},
	"str_split": &object.Builtin{
		Fn: runtime.StringSplit,
	},
	"str_lower": &object.Builtin{
		Fn: runtime.StringLower,
	},
	"str_upper": &object.Builtin{
		Fn: runtime.StringUpper,
	},

	"file_write_lines": &object.Builtin{
		Fn: runtime.FileWriteLines,
	},
	"file_read_lines": &object.Builtin{
		Fn: runtime.FileReadLines,
	},
}
