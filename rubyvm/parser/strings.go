package parser

func lexSingleQuoteString(l StatefulRubyLexer) stateFn {
	var (
		r    rune
		prev rune
	)

	l.ignore() // ignore single quote

	for {
		prev = r
		switch r = l.next(); {
		case r == '\'' && prev != '\\':
			l.backup()
			l.emit(tokenTypeString)
			l.accept("'")
			l.ignore()
			return lexSomething
		case r == eof:
			l.emit(tokenTypeError)
			return lexSomething
		}
	}

	return lexSomething
}

func lexDoubleQuoteString(l StatefulRubyLexer) stateFn {
	var (
		r    rune
		prev rune
	)

	l.moveCurrentTokenStartIndex(1)

	for {
		prev = r
		switch r = l.next(); {
		case r == '#':
			if l.accept("{") {
				lexUntilClosingMatchingBraces('{', '}')(l)
			}
		case r == '"' && prev != '\\':
			l.moveCurrentPositionIndex(-1)
			l.emit(tokenTypeDoubleQuoteString)
			l.next()
			l.ignore()
			return lexSomething
		case r == eof:
			l.emit(tokenTypeError)
			return lexSomething
		}
	}

	return lexSomething
}

func lexUntilClosingMatchingBraces(openingBrace, closingBrace rune) func(StatefulRubyLexer) {
	return func(l StatefulRubyLexer) {
		for {
			switch r := l.next(); {
			case r == openingBrace:
				lexUntilClosingMatchingBraces(openingBrace, closingBrace)(l)
			case r == closingBrace:
				return
			case r == eof:
				l.emit(tokenTypeError)
				return
			}
		}
	}
}
