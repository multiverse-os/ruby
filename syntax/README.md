# Ruby Syntax Library
Building a fully compatibile Go interpreter, and virtual machine that can read
byte code. That will support using Go inline in the same way you can use C
inline with CRuby. GoRuby will likely also migrate several of the core libraries
and key tools into Go for greater preformance, since it appears a substantial
portion of the Ruby codebase is written in Ruby without much consideration for
what pieces could be left in the lower level language for preformance reasons.  

For example, we currently have `irb` implemented in Go instead of Ruby and we
get faster response time which is nice when experimenting. 

A core part of this is working with Ruby syntax. And we want to make the
different components of the GoRuby implementation to be useful as individual
submodules that can be used as libraries for a variety of purposes. 

### Core Components

  * **Syntax** subcomponent will be a stand alone library with the keywords, 
  global variables, method names, tokens, and other components of the language 
  divided up and implemented in such a way it will be easy to upgrade the 
  syntax when new versions are released. We will always strive to follow CRuby 
  and maintain full compatibility. **We think the best way may be to find the 
  best places to parse and build our Go files automatically.** 
  *Additionally we would like to be able to easily extend GoRuby from the Go 
  level. While you can always extend Ruby from your Ruby application, doing 
  so in the low level language would have significant preformance increase 
  over doing it inline. Making this trivial could enable developers to 
  increase the speed of their applications by moving some of their code to 
  Go.*

  * **Interpreter/Evaluator** subcomponent with the goal of providing direct 
  evaluation of Ruby source without converting toa bytecode. 

  * **Transpiler** subcomponent with the goal of providing transpiling to Go,
  and then using Go to transpile to Javascript, Python and other high level
  languages that have complete support in Go.

  * **Parser** subcomponent that can be used to parse Ruby and will be used with
  the lexagraphical scanner to prepare the Ruby for processing.

  * **Lexagraphical Scanner** uses the tokens declared in the `syntax`
  subcomponent to load Ruby and prepare for processing.

  * **REPL** A read–eval–print loop (REPL), also termed an interactive toplevel 
  or language shell, is a simple, interactive computer programming environment 
  that takes single user inputs (i.e., single expressions), evaluates (executes) 
  them, and returns the result to the user; a program written in a REPL 
  environment is executed piecewise. The term is usually used to refer to 
  programming interfaces similar to the classic Lisp machine interactive 
  environment. Common examples include command line shells and similar 
  environments for programming languages, and the technique is very 
  characteristic of scripting languages.[1]

