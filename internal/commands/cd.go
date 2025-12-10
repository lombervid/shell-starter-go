package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type cd struct {
	currentDirectory string
}

func (c *cd) Exec(args ...string) {
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "cd: too many arguments")
		return
	}

	path := args[0]
	ok := c.chdir(path)

	if ok == false {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", path)
	}
}

func (c *cd) chdir(path string) bool {
	newPath, ok := c.resolvePath(path)
	if ok == false {
		return false
	}

	err := os.Chdir(newPath)
	if err != nil {
		return false
	}

	c.currentDirectory = newPath

	return true
}

func (c *cd) resolvePath(path string) (newPath string, ok bool) {
	if path == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", false
		}
		return homeDir, true
	}

	// set the start path to the current directory
	newPath = c.currentDirectory

	if strings.HasPrefix(path, "/") {
		// absolute path
		return path, true
	} else if strings.HasPrefix(path, "~") {
		// start from user directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", false
		}
		path = strings.Replace(path, "~", "", 1)
		newPath = homeDir
	}

	dirs := strings.Split(path, "/")
	for _, dir := range dirs {
		switch dir {
		case "", ".":
			// same directory
		case "..":
			// go bach one directory
			newPath = filepath.Dir(newPath)
		default:
			newPath = fmt.Sprintf("%s/%s", newPath, dir)
		}
	}

	ok = true
	newPath = strings.ReplaceAll(newPath, "//", "/") // clean path
	return
}
