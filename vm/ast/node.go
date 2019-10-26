package ast

type Nodes []Node

// Node represents a node in Intermediate Representation's AST
type ASTNode struct {
	Name  string
	Nodes []Node
}

// NewNode returns a new Node instance (*Node)
func NewASTNode() ASTNode {
	return ASTNode{}
}

// TODO: What the fucking fuck? Are interface not for having multiple different things have a similar interface ? Because if so no sense. ..
type Node interface {
	LineNumber() int
}

func (nodes Nodes) LineNumber() int {
	return nodes[0].LineNumber()
}

type ConstantInt struct {
	Line  int
	Value int64
}

func (n ConstantInt) LineNumber() int {
	return n.Line
}

type ConstantUint struct {
	Line  int
	Value uint64
}

func (n ConstantUint) LineNumber() int {
	return n.Line
}

type ConstantFloat struct {
	Line  int
	Value float64
}

func (n ConstantFloat) LineNumber() int {
	return n.Line
}

type SimpleString struct {
	Line  int
	Value string
}

func (n SimpleString) LineNumber() int {
	return n.Line
}

func (n SimpleString) StringValue() string {
	return n.Value
}

type InterpolatedString struct {
	Line  int
	Value string
}

func (n InterpolatedString) LineNumber() int {
	return n.Line
}

func (n InterpolatedString) StringValue() string {
	return n.Value
}

type String interface {
	Node
	StringValue() string
}

type CharacterLiteral struct {
	Line  int
	Value string
}

func (n CharacterLiteral) LineNumber() int {
	return n.Line
}

type Symbol struct {
	Line int
	Name string
}

func (n Symbol) LineNumber() int {
	return n.Line
}

type BareReference struct {
	Line int
	Name string
}

func (n BareReference) LineNumber() int {
	return n.Line
}

type Constant struct {
	Line int
	Name string
}

func (n Constant) LineNumber() int {
	return n.Line
}

type CallExpression struct {
	Line          int
	Target        Node
	Func          BareReference
	Args          []Node
	OptionalBlock Block
}

func (n CallExpression) LineNumber() int {
	return n.Line
}

type FuncDecl struct {
	Line    int
	Target  Node
	Name    BareReference
	Args    []MethodParam
	Body    []Node
	Rescues []Node
	Ensure  []Node
}

func (n FuncDecl) LineNumber() int {
	return n.Line
}

func (f FuncDecl) MethodName() string {
	return f.Name.Name
}

type ClassDecl struct {
	Line       int
	Name       string
	SuperClass Class
	Namespace  string
	Body       []Node
}

func (n ClassDecl) LineNumber() int {
	return n.Line
}

func (c ClassDecl) FullName() string {
	if c.Namespace != "" {
		return c.Namespace + "::" + c.Name
	} else {
		return c.Name
	}
}

type Class struct {
	Line              int
	Name              string
	Namespace         string
	IsGlobalNamespace bool
}

func (n Class) LineNumber() int {
	return n.Line
}

func (c Class) FullName() string {
	if c.Namespace != "" {
		return c.Namespace + "::" + c.Name
	} else {
		return c.Name
	}
}

type ModuleDecl struct {
	Line      int
	Name      string
	Namespace string
	Body      []Node
}

func (n ModuleDecl) LineNumber() int {
	return n.Line
}

func (m ModuleDecl) FullName() string {
	if m.Namespace != "" {
		return m.Namespace + "::" + m.Name
	} else {
		return m.Name
	}
}

type Assignment struct {
	Line int
	LHS  Node
	RHS  Node
}

func (n Assignment) LineNumber() int {
	return n.Line
}

type Boolean struct {
	Line  int
	Value bool
}

func (n Boolean) LineNumber() int {
	return n.Line
}

type Negation struct {
	Line   int
	Target Node
}

func (n Negation) LineNumber() int {
	return n.Line
}

type Complement struct {
	Line   int
	Target Node
}

func (n Complement) LineNumber() int {
	return n.Line
}

type Positive struct {
	Line   int
	Target Node
}

func (n Positive) LineNumber() int {
	return n.Line
}

type Negative struct {
	Line   int
	Target Node
}

func (n Negative) LineNumber() int {
	return n.Line
}

type Addition struct {
	Line int
	LHS  Node
	RHS  Node
}

func (n Addition) LineNumber() int {
	return n.Line
}

type Subtraction struct {
	Line int
	LHS  Node
	RHS  Node
}

func (n Subtraction) LineNumber() int {
	return n.Line
}

type Multiplication struct {
	Line int
	LHS  Node
	RHS  Node
}

func (n Multiplication) LineNumber() int {
	return n.Line
}

type Array struct {
	Line  int
	Nodes []Node
}

func (n Array) LineNumber() int {
	return n.Line
}

type Hash struct {
	Line  int
	Pairs []HashKeyValuePair
}

func (n Hash) LineNumber() int {
	return n.Line
}

type HashKeyValuePair struct {
	Line  int
	Key   Node
	Value Node
}

func (n HashKeyValuePair) LineNumber() int {
	return n.Line
}

type GlobalVariable struct {
	Line int
	Name string
}

func (n GlobalVariable) LineNumber() int {
	return n.Line
}

type InstanceVariable struct {
	Line int
	Name string
}

func (n InstanceVariable) LineNumber() int {
	return n.Line
}

type ClassVariable struct {
	Line int
	Name string
}

func (n ClassVariable) LineNumber() int {
	return n.Line
}

type FileNameConstReference struct {
	Line int
}

func (n FileNameConstReference) LineNumber() int {
	return n.Line
}

type LineNumberConstReference struct {
	Line int
}

func (n LineNumberConstReference) LineNumber() int {
	return n.Line
}

type Block struct {
	Line int
	Args []MethodParam
	Body []Node
}

func (n Block) LineNumber() int {
	return n.Line
}

func (b *Block) Provided() bool {
	return b.Body != nil
}

type IfBlock struct {
	Line      int
	Condition Node
	Body      []Node
	Else      []Node
}

func (n IfBlock) LineNumber() int {
	return n.Line
}

type Subshell struct {
	Line    int
	Command string
}

func (n Subshell) LineNumber() int {
	return n.Line
}

type Group struct {
	Line int
	Body []Node
}

func (n Group) LineNumber() int {
	return n.Line
}

type Begin struct {
	Line   int
	Body   []Node
	Rescue []Node
	Else   []Node
	Ensure []Node
}

func (n Begin) LineNumber() int {
	return n.Line
}

type Rescue struct {
	Line      int
	Body      []Node
	Exception RescueException
}

func (n Rescue) LineNumber() int {
	return n.Line
}

type RescueException struct {
	Line    int
	Var     Node
	Classes []Class
}

func (n RescueException) LineNumber() int {
	return n.Line
}

type MethodParam struct {
	Name         string
	DefaultValue Node
	IsSplat      bool
	IsProc       bool
}

type Ternary struct {
	Line      int
	Condition Node
	True      Node
	False     Node
}

func (n Ternary) LineNumber() int {
	return n.Line
}

type Yield struct {
	Line  int
	Value Node
}

func (n Yield) LineNumber() int {
	return n.Line
}

type Return struct {
	Line  int
	Value Node
}

func (n Return) LineNumber() int {
	return n.Line
}

type Next struct {
	Line int
}

func (n Next) LineNumber() int {
	return n.Line
}

type Redo struct {
	Line int
}

func (n Redo) LineNumber() int {
	return n.Line
}

type Break struct {
	Line int
}

func (n Break) LineNumber() int {
	return n.Line
}

type Retry struct {
	Line int
}

func (n Retry) LineNumber() int {
	return n.Line
}

type Loop struct {
	Line      int
	Condition Node
	Body      []Node
}

func (n Loop) LineNumber() int {
	return n.Line
}

type WeakLogicalAnd struct {
	Line int
	LHS  Node
	RHS  Node
}

func (n WeakLogicalAnd) LineNumber() int {
	return n.Line
}

type WeakLogicalOr struct {
	Line int
	LHS  Node
	RHS  Node
}

func (n WeakLogicalOr) LineNumber() int {
	return n.Line
}

type Lambda struct {
	Line int
	Body Block
}

func (n Lambda) LineNumber() int {
	return n.Line
}

type SwitchStatement struct {
	Line      int
	Condition Node
	Cases     []SwitchCase
	Else      []Node
}

func (n SwitchStatement) LineNumber() int {
	return n.Line
}

type SwitchCase struct {
	Line       int
	Conditions []Node
	Body       []Node
}

func (n SwitchCase) LineNumber() int {
	return n.Line
}

type ConditionalAssignment struct {
	Line int
	LHS  Node
	RHS  Node
}

func (n ConditionalAssignment) LineNumber() int {
	return n.Line
}

type ConditionalTruthyAssignment struct {
	Line int
	LHS  Node
	RHS  Node
}

func (n ConditionalTruthyAssignment) LineNumber() int {
	return n.Line
}

type Range struct {
	Line             int
	Start            Node
	End              Node
	ExcludeLastValue bool
}

func (n Range) LineNumber() int {
	return n.Line
}

type StarSplat struct {
	Line  int
	Value Node
}

func (n StarSplat) LineNumber() int {
	return n.Line
}

type RescueModifier struct {
	Line      int
	Statement Node
	Rescue    Node
}

func (n RescueModifier) LineNumber() int {
	return n.Line
}

type Regex struct {
	Line  int
	Value string
}

func (n Regex) LineNumber() int {
	return n.Line
}

type Eigenclass struct {
	Line   int
	Target Node
	Body   []Node
}

func (n Eigenclass) LineNumber() int {
	return n.Line
}

type Alias struct {
	Line int
	To   Symbol
	From Symbol
}

func (n Alias) LineNumber() int {
	return n.Line
}

type Nil struct {
	Line int
}

func (n Nil) LineNumber() int {
	return n.Line
}

type Self struct {
	Line int
}

func (n Self) LineNumber() int {
	return n.Line
}

type Defined struct {
	Line int
	Node Node
}

func (d Defined) LineNumber() int {
	return d.Line
}

type SuperclassMethodImplCall struct {
	Line int
	Args []Node
}

func (s SuperclassMethodImplCall) LineNumber() int {
	return s.Line
}
