package parser

// TODO: Is this faster than just doing if a < thing && thing > z? Because if not this is just wasting memory for no real reason
const (
	alphaLower             = "abcdefghijklmnopqrstuvwxyz"
	alphaUpper             = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeric                = "0123456789"
	alpha                  = alphaLower + alphaUpper
	alphaUnderscore        = alphaLower + alphaUpper + "_"
	alphaNumeric           = alphaLower + alphaUpper + numeric
	alphaNumericUnderscore = alphaLower + alphaUpper + numeric + "_"
	validMethodNameRunes   = alphaLower + alphaUpper + "_!?"
)

func lexSymbol(l StatefulRubyLexer) stateFn {
	if !l.accept(alpha + "_@\"") {
		if l.accept(":") {
			l.acceptRun(alphaNumericUnderscore)

			for l.currentIndex()+2 < l.lengthOfInput() && l.slice(l.currentIndex(), l.currentIndex()+2) == "::" {
				l.moveCurrentPositionIndex(2)
				l.acceptRun(alphaNumericUnderscore)
			}

			l.emit(tokenTypeNamespaceResolvedModule)
		} else {
			l.emit(tokenTypeColon)
		}

		return lexSomething
	}

	// skip past the initial colon
	l.moveCurrentTokenStartIndex(1)

	// some dynamic symbols can start with " and '
	if l.slice(l.currentIndex()-1, l.currentIndex()) == "\"" {
		var (
			r    rune
			prev rune
		)

		l.ignore() // ignore : and opening quote

		for {
			prev = r
			switch r = l.next(); {
			case r == '#':
				if l.accept("{") {
					// check that we close the #{} template if present
					for innerR := l.next(); innerR != '}'; innerR = l.next() {
						if innerR == eof {
							l.emit(tokenTypeError)
						}
					}
				}
			case r == '"' && prev != '\\':
				l.moveCurrentPositionIndex(-1)
				l.emit(tokenTypeSymbol)
				l.next()
				l.ignore()
				return lexSomething
			case r == eof:
				l.emit(tokenTypeError)
				return lexSomething
			}
		}
	}

	l.accept("@")
	l.acceptRun(alphaNumericUnderscore + "?!")
	l.emit(tokenTypeSymbol)
	return lexSomething
}
