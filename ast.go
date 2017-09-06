package gotcha

type Position struct {
	Line   int
	Column int
}

func (position *Position) Offset(n int) (newPos Position) {
	newPos.Line = position.Line
	newPos.Column = position.Column + n

	return newPos
}

type SourceLocation struct {
	Source *string
	Start Position
	Stop Position
}

type Node struct {
	Type string
	Loc *SourceLocation
}

// interface Identifier <: Expression, Pattern
type Identifier struct {
	Node
	name string
}

func (ident *Identifier) Pattern() Node {
	return ident.Node
}

func (ident *Identifier) Expression() Node {
	return ident.Node
}

// interface Literal <: Expression
type Literal struct {
	Node
	Value interface{}
}

func (literal *Literal) Expression() Node {
	return literal.Node
}

type Regex struct {
	Pattern string
	Flags string
}

// interface RegExpLiteral <: Literal
type RegExpLiteral struct {
	Literal
	regex Regex
}

// interface Program <: Node
type Program struct {
	Node
	Body []Statement
}

// interface Function <: Node
type Function struct {
	Node
	Id *Identifier
	Params []Pattern
	Body []FunctionBody
}

// interface Statement <: Node
type Statement interface {
	Statement() Node
}

// interface ExpressionStatement <: Statement
type ExpressionStatement struct {
	Node
}

func (exprStmt *ExpressionStatement) Statement() Node {
	return exprStmt.Node
}

// interface Directive <: ExpressionStatement
type Directive struct {
	ExpressionStatement
	Directive string
}

// interface BlockStatement <: Statement
type BlockStatement struct {
	Node
	Body []Statement
}

func (blockStmt *BlockStatement) Statement() Node {
	return blockStmt.Node
}

// interface FunctionBody <: BlockStatement
type FunctionBody struct {
	BlockStatement
}

// interface EmptyStatement <: Statement
type EmptyStatement struct {
	Node
}

func (emptyStmt *EmptyStatement) Statement() Node {
	return emptyStmt.Node
}

// interface DebuggerStatement <: Statement
type DebuggerStatement struct {
	Node
}

func (debuggerStmt *DebuggerStatement) Statement() Node {
	return debuggerStmt.Node
}

type Expression interface {
	Expression() Node
}

type Pattern interface {
	Pattern() Node
}