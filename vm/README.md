[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse OS: Go Ruby `vm` implementation
**URL** [multiverse-os.org](https://multiverse-os.org)

The Multiverse OS ruby project provides various solutions for working with Ruby
from Go. These are stop gap solutions until we can complete our implementation
of the Ruby VM in Go. 

Our goal is the ability to parse/lex, for interpretation, and the ability to
compile down into Ruby bytecode and run on the VM. Ideally we want the ability
to transpile into Go as well which would give a variety of possible solutions. 

The VM will be used in the `multiplex shell (multish)`, and provide a scripting
language for Multiverse OS configurations. 


### Progress
The following will track our progress and help guide development:

## Supported features
The `ruby` command:
- [ ] parse program files
- [ ] program file arguments
- [ ] Flags
  - [ ] `-0[octal]`       specify record separator (\0, if no argument)
  - [ ] `-a`              autosplit mode with -n or -p (splits $_ into $F)
  - [ ] `-c`              check syntax only
  - [ ] `-Cdirectory`     cd to directory before executing your script
  - [ ] `-d`              set debugging flags (set $DEBUG to true)
  - [ ] `-e 'command'`    one line of script. Several -e's allowed. Omit [programfile]
  - [ ] `-Eex[:in]`       specify the default external and internal character encodings
  - [ ] `-Fpattern`       split() pattern for autosplit (-a)
  - [ ] `-i[extension]`   edit ARGV files in place (make backup if extension supplied)
  - [ ] `-Idirectory`     specify $LOAD_PATH directory (may be used more than once)
  - [ ] `-l`              enable line ending processing
  - [ ] `-n`              assume 'while gets(); ... end' loop around your script
  - [ ] `-p`              assume loop like -n but print line also like sed
  - [ ] `-rlibrary`       require the library before executing your script
  - [ ] `-s`              enable some switch parsing for switches after script name
  - [ ] `-S`              look for the script using PATH environment variable
  - [ ] `-T[level=1]`     turn on tainting checks
  - [ ] `-v`              print version number, then turn on verbose mode
  - [ ] `-w`              turn warnings on for your script
  - [ ] `-W[level=2]`     set warning level; 0=silence, 1=medium, 2=verbose
  - [ ] `-x[directory]`   strip off text before #!ruby line and perhaps cd to directory
  - [ ] `-h`              show this message, --help for more info

The `irb` command:
- [ ] parse program files
- [ ] program file arguments
- [ ] Flags
  - [ ] `-f`		    Suppress read of ~/.irbrc
  - [ ] `-m`		    Bc mode (load mathn, fraction or matrix are available)
  - [ ] `-d`                Set $DEBUG to true (same as `ruby -d')
  - [ ] `-r load-module`    Same as `ruby -r'
  - [ ] `-I path`           Specify $LOAD_PATH directory
  - [ ] `-U`                Same as `ruby -U`
  - [ ] `-E enc`            Same as `ruby -E`
  - [ ] `-w`                Same as `ruby -w`
  - [ ] `-W[level=2]`       Same as `ruby -W`
  - [ ] `--context-mode n`  Set n[0-3] to method to create Binding Object,
                    when new workspace was created
  - [ ] `--echo`            Show result(default)
  - [ ] `--noecho`          Don't show result
  - [ ] `--inspect`	    Use `inspect' for output (default except for bc mode)
  - [ ] `--noinspect`	    Don't use inspect for output
  - [ ] `--readline`        Use Readline extension module
  - [ ] `--noreadline`	    Don't use Readline extension module
  - [ ] `--prompt prompt-mode`/`--prompt-mode prompt-mode`
		    Switch prompt mode. Pre-defined prompt modes are
		    `default', `simple', `xmp' and `inf-ruby'
  - [ ] `--inf-ruby-mode`   Use prompt appropriate for inf-ruby-mode on emacs.
		    Suppresses --readline.
  - [ ] `--sample-book-mode`/`--simple-prompt`
                    Simple prompt mode
  - [ ] `--noprompt`        No prompt mode
  - [ ] `--single-irb`      Share self with sub-irb.
  - [ ] `--tracer`          Display trace for each execution of commands.
  - [ ] `--back-trace-limit n`
		    Display backtrace top n and tail n. The default
		    value is 16.
  - [ ] `--irb_debug n`	    Set internal debug level to n (not for popular use)
  - [ ] `--verbose`         Show details
  - [ ] `--noverbose`       Don't show details
  - [ ] `-v`, `--version`	  Print the version of irb
  - [ ] `-h`, `--help`      Print help
  - [ ] `--`                Separate options of irb from the list of command-line args


### Supported language feature
- [ ] everything is an object
	- [ ] allow method calls on everything
	- [ ] operators are method calls
- [ ] full UTF8 support
	- [ ] Unicode identifier
	- [ ] Unicode symbols
- [ ] functions
	- [ ] with parens
	- [ ] without parens
	- [ ] return keyword
	- [ ] default values for parameters
	- [ ] keyword arguments
	- [ ] block arguments
	- [ ] hash as last argument without braces
- [ ] function calls
	- [ ] with parens
	- [ ] without parens	
	- [ ] with block arguments
- [ ] conditionals
	- [ ] if
	- [ ] if/else
	- [ ] if/elif/else
	- [ ] tenary `? : `
	- [ ] unless
	- [ ] unless/else
	- [ ] case
	- [ ] `||`
	- [ ] `&&`
- [ ] control flow
	- [ ] for loop
	- [ ] while loop
	- [ ] until loop
	- [ ] break
	- [ ] next
	- [ ] redo
	- [ ] flip flop
- [ ] numbers
	- [ ] integers
		- [ ] integer arithmetics
		- [ ] integers `1234`
		- [ ] integers with underscores `1_234`
		- [ ] decimal numbers `0d170`, `0D170`
		- [ ] octal numbers `0252`, `0o252`, `0O252`
		- [ ] hexadecimal numbers `0xaa`, `0xAa`, `0xAA`, `0Xaa`, `0XAa`, `0XaA`
		- [ ] binary numbers `0b10101010`, `0B10101010`
	- [ ] floats
		- [ ] float arithmetics
		- [ ] `12.34`
		- [ ] `1234e-2`
		- [ ] `1.234E1`
		- [ ] floats with underscores `2.2_22`
- [ ] booleans
- [ ] strings
	- [ ] double quoted
	- [ ] single quoted
	- [ ] character literals (`?\n`, `?a`,...)
	- [ ] `%q{}`
	- [ ] `%Q{}`
	- [ ] heredoc
		- [ ] without indentation (`<<EOF`)
		- [ ] indented (`<<-EOF`)
		- [ ] “squiggly” heredoc `<<~`
		- [ ] quoted heredoc
			- [ ] single quotes `<<-'HEREDOC'`
 			- [ ] double quotes `<<-"HEREDOC"`
 			- [ ] backticks <<-\`HEREDOC\`"
	- [ ] escaped characters
		- [ ] `\a` bell, ASCII 07h (BEL)
		- [ ] 	`\b` backspace, ASCII 08h (BS)
		- [ ] 	`\t` horizontal tab, ASCII 09h (TAB)
		- [ ] 	`\n` newline (line feed), ASCII 0Ah (LF)
		- [ ] 	`\v` vertical tab, ASCII 0Bh (VT)
		- [ ] 	`\f` form feed, ASCII 0Ch (FF)
		- [ ] 	`\r` carriage return, ASCII 0Dh (CR)
		- [ ] 	`\e` escape, ASCII 1Bh (ESC)
		- [ ] 	`\s` space, ASCII 20h (SPC)
		- [ ] 	`\\` backslash, \
		- [ ] 	`\nnn` octal bit pattern, where nnn is 1-3 octal digits ([0-7])
		- [ ] 	`\xnn` hexadecimal bit pattern, where nn is 1-2 hexadecimal digits ([0-9a-fA-F])
		- [ ] `\unnnn` Unicode character, where nnnn is exactly 4 hexadecimal digits ([0-9a-fA-F])
		- [ ] `\u{nnnn ...}` Unicode character(s), where each nnnn is 1-6 hexadecimal digits ([0-9a-fA-F])
		- [ ] `\cx` or `\C-x` control character, where x is an ASCII printable character
		- [ ] `\M-x` meta character, where x is an ASCII printable character
		- [ ] `\M-\C-x` meta control character, where x is an ASCII printable character
		- [ ] `\M-\cx` same as above
		- [ ] `\c\M-x` same as above
		- [ ] `\c?` or `\C-?` delete, ASCII 7Fh (DEL)
	- [ ] interpolation `#{}`
	- [ ] automatic concatenation
- [ ] arrays
	- [ ] array literal `[1,2]`
	- [ ] array indexing `arr[2]`
	- [ ] splat
	- [ ] array decomposition
	- [ ] implicit array assignment
	- [ ] array of strings `%w{}`
	- [ ] array of symbols `%i{}`
- [ ] nil
- [ ] hashes
	- [ ] literal with `=>` notation
	- [ ] literal with `key:` notation
	- [ ] indexing `hash[:foo]`
	- [ ] every Ruby Object can be a hash key
- [ ] symbols
	- [ ] `:symbol`
	- [ ] `:"symbol"`
	- [ ] `:"symbol"` with interpolation
	- [ ] `:'symbol'`
	- [ ] `%s{symbol}`
	- [ ] singleton symbols
- [ ] regexp
	- [ ] `/regex/`
	- [ ] `%r{regex}`
- [ ] ranges
	- [ ] `..` inclusive
	- [ ] `...` exclusive
- [ ] procs `->`
- [ ] variables
	- [ ] variable assignments
	- [ ] globals
- [ ] operators
	- [ ] `+`
	- [ ] `-`
	- [ ] `/`
	- [ ] `*`
	- [ ] `!`
	- [ ] `<`
	- [ ] `>`
	- [ ] `**` (pow)
	- [ ] `%` (modulus)
	- [ ] `&` (AND)
	- [ ] `^` (XOR)
	- [ ] `>>` (right shift)
	- [ ] `<<` (left shift, append)
	- [ ] `==` (equal)
	- [ ] `!=` (not equal)
	- [ ] `===` (case equality)
	- [ ] `=~` (pattern match)
	- [ ] `!~` (does not match)
	- [ ] `<=>` (comparison or spaceship operator)
	- [ ] `<=` (less or equal)
	- [ ] `>=` (greater or equal)
	- [ ] assignment operators
		- [ ] `+=`
		- [ ] `-=`
		- [ ] `/=`
		- [ ] `*=`
		- [ ] `%=`
		- [ ] `**=`
		- [ ] `&=`
		- [ ] `|=`
		- [ ] `^=`
		- [ ] `<<=`
		- [ ] `>>=`
		- [ ] `||=`
		- [ ] `&&=`
- [ ] function blocks (procs)
- [ ] error handling
	- [ ] begin/rescue
	- [ ] ensure
	- [ ] retry
- [ ] constants
- [ ] scope operator `::`
- [ ] classes
	- [ ] class objects
	- [ ] class Class
	- [ ] instance variables
	- [ ] class variables
	- [ ] class methods
	- [ ] instance methods
	- [ ] method overrides
	- [ ] private
	- [ ] protected
	- [ ] public
	- [ ] inheritance
	- [ ] constructors
	- [ ] new
	- [ ] `self`
	- [ ] singleton classes (also known as the metaclass or eigenclass) `class << self`
	- [ ] assigment methods
	- [ ] self defined classes
	- [ ] self defined classes with inheritance
- [ ] modules
- [ ] object main
- [ ] comments '#'
