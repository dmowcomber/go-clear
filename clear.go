package clear

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (
	clear            map[string]func()
	goos             = runtime.GOOS
	linuxEquivalents = []string{"linux", "darwin", "android", "solaris", "openbsd", "freebsd"}
	// used for testing because I can't mock the function exec.Command
	lastCommand *exec.Cmd
)

func init() {
	clear = make(map[string]func())

	for _, linuxEquivalent := range linuxEquivalents {
		clear[linuxEquivalent] = func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			lastCommand = cmd
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		lastCommand = cmd
	}
}

func setGOOS(os string) {
	goos = os
}

func Clear() error {
	value, ok := clear[goos]
	if ok {
		value()
		return nil
	} else {
		return fmt.Errorf("Your os %q is unsupported. Terminal screen cannot be cleared", goos)
	}
}
