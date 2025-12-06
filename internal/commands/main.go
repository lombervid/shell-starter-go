package commands

type command interface {
	Execute(args string)
}

var Builtin = map[string]command{
	"exit": exit{},
	"echo": echo{},
	"type": typeCmd{},
}
