package commands

type execer interface {
	Exec(args string)
}

var Builtin = map[string]execer{
	"exit": exit{},
	"echo": echo{},
	"type": typeCmd{},
}
