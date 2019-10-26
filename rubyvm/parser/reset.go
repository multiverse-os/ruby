package parser

import "github.com/multiverse-os/ruby/rubyvm/ast"

func Reset() {
	Statements = []ast.Node{}
	DebugStatements = []string{}
}
