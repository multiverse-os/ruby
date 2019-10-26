package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	parser "github.com/multiverse-os/ruby/syntax/parser"
	vm "github.com/multiverse-os/ruby/vm"
)

var verboseFlag = flag.Bool("verbose", false, "enables verbose mode")

func init() {
	flag.BoolVar(verboseFlag, "v", false, "enables verbose mode")
}

// TODO: The goal should be to match the exact behavior the c version provides
// so if no argument is provided it should not error, it should hold open and
// and let ruby to be typed directly. Additionally we should provide all the
// flags like -e for example
func main() {
	flag.Parse()

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("it should hold open and allow user to type in ruby, but currently requires a filename argument, try again with a filename")
		os.Exit(1)
	}

	// for now, assumes this is only being invoked with a filename to interpret
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	home := os.Getenv("HOME")
	rubyvmHome := filepath.Join(home, ".rubyvm")

	rubyVM := vm.NewVM(rubyvmHome, flag.Args()[0])
	defer rubyVM.Exit()

	_, err = rubyVM.Run(string(bytes))

	switch err.(type) {
	case *vm.ParseError:
		offendingFilename := err.(*vm.ParseError).Filename
		println(fmt.Sprintf("Error parsing ruby script %s", offendingFilename))
		println("last 20 statements from the parser:")
		println("")

		debugStatements := []string{}
		for _, d := range parser.DebugStatements {
			debugStatements = append(debugStatements, d)
		}

		threshold := 61
		debugCount := len(debugStatements)
		if debugCount <= threshold {
			for _, stmt := range debugStatements {
				fmt.Printf("\t%s\n", stmt)
			}
		} else {
			for _, stmt := range debugStatements[debugCount-threshold:] {
				fmt.Printf("\t%s\n", stmt)
			}
		}

		os.Exit(1)
	case nil:
	case error:
		panic(err.Error())
	default:
		panic(fmt.Sprintf("%#v", err))
	}
}
