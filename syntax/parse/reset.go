package parser

import (
	ast "github.com/multiverse-os/ruby/rubyvm/vm/ast"
)

func Reset() {
	Statements = []ast.Node{}
	DebugStatements = []string{}
}
