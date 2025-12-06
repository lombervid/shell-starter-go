package builtin

type command interface {
	Execute(args string)
}

var Commands = map[string]command{
	"exit": exit{},
	"echo": echo{},
	"type": typeCmd{},
}
