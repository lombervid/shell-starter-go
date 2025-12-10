package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type execer interface {
	Exec(args ...string)
}

type executable struct {
	name string
	path string
}

func (c executable) Exec(args ...string) {
	output, err := exec.Command(c.name, args...).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Printf("%s", output)
}

var PathDirs = []string{}
var Executables = make(map[string]executable)
var Builtin = map[string]execer{
	"exit": exit{},
	"echo": echo{},
	"type": typeCmd{},
	"pwd":  pwd{},
}

func init() {
	path := os.Getenv("PATH")
	PathDirs = strings.Split(path, ":")
}

func FindExecutable(name string) (executable, bool) {
	if exec, ok := Executables[name]; ok {
		return exec, true
	}

	for _, dir := range PathDirs {
		if isExecutableIn(dir, name) {
			exec := executable{
				name: name,
				path: fmt.Sprintf("%s/%s", dir, name),
			}
			Executables[name] = exec
			return exec, true
		}
	}

	return executable{}, false
}

func isExecutableIn(directory, name string) bool {
	files, err := os.ReadDir(directory)
	if err != nil {
		return false
	}

	for _, file := range files {
		if file.Name() != name {
			continue
		}

		if file.IsDir() {
			continue
		}

		info, err := file.Info()
		if err != nil {
			continue
		}

		// Not executable
		if info.Mode()&0111 == 0 {
			continue
		}

		return true
	}

	return false
}
