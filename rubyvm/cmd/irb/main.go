package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/multiverse-os/ruby/rubyvm/interpreter/vm"
)

func main() {
	home := os.Getenv("HOME")
	rubyvmHome := filepath.Join(home, ".rubyvm")

	vm := vm.NewVM(rubyvmHome, "(rubyvm irb")
	defer vm.Exit()

	for {
		txt := readInput()
		if txt == "quit\n" {
			break
		}
		result, err := vm.Run(txt)
		if err != nil {
			fmt.Printf(" => %s", err.Error())
			continue
		}

		if result != nil {
			fmt.Printf("=> %s", result.String())
		} else {
			fmt.Printf("=> %#v", result)
		}
		println("")
	}
}

func readInput() string {
	print("> ")
	bio := bufio.NewReader(os.Stdin)
	userInput, err := bio.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return userInput
}
