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

func (*Identifier) Pattern() {}

func (*Identifier) Expression() {}

func (*Identifier) forStatementInit() {}

func (*Identifier) forInStatementLeft() {}

// interface Literal <: Expression
type Literal struct {
	Node
	Value interface{}
}

func (*Literal) Expression() {}

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
	Statement()
}

// interface ExpressionStatement <: Statement
type ExpressionStatement struct {
	Node
	Expression Expression
}

func (exprStmt *ExpressionStatement) Statement() {}

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

func (*BlockStatement) Statement() {}

// interface FunctionBody <: BlockStatement
type FunctionBody struct {
	BlockStatement
}

// interface EmptyStatement <: Statement
type EmptyStatement struct {
	Node
}

func (*EmptyStatement) Statement() {}

// interface DebuggerStatement <: Statement
type DebuggerStatement struct {
	Node
}

func (*DebuggerStatement) Statement() {}

// interface WithStatement <: Statement
type WithStatement struct {
	Node
	Object Expression
	Body Statement
}

func (*WithStatement) Statement() {}

// interface ReturnStatement <: Statement
type ReturnStatement struct {
	Node
	Argument *Expression
}

func (*ReturnStatement) Statement() {}

// interface LabeledStatement <: Statement
type LabeledStatement struct {
	Node
	Label Identifier
	Body Statement
}

func (*LabeledStatement) Statement() {}

// interface BreakStatement <: Statement
type BreakStatement struct {
	Node
	Label *Identifier
}

func (*BreakStatement) Statement() {}

// interface ContinueStatement <: Statement
type ContinueStatement struct {
	Node
	Label *Identifier
}

func (*ContinueStatement) Statement() {}

// interface IfStatement <: Statement
type IfStatement struct {
	Node
	Test Expression
	Consequent Statement
	Alternate *Statement
}

func (*IfStatement) Statement() {}

// interface SwitchStatement <: Statement
type SwitchStatement struct {
	Node
	Discriminant Expression
	Cases []SwitchCase
}

func (*SwitchStatement) Statement() {}

// interface SwitchCase <: Node
type SwitchCase struct {
	Node
	Test *Expression
	Consequent []Statement
}

// interface ThrowStatement <: Statement
type ThrowStatement struct {
	Node
	Argument Expression
}

func (*ThrowStatement) Statement() {}

// interface TryStatement <: Statement
type TryStatement struct {
	Node
	Block BlockStatement
	Handler *CatchClause
	Finalizer *BlockStatement
}

func (*TryStatement) Statement() {}

// interface CatchClause <: Node
type CatchClause struct {
	Node
	Param Pattern
	Body BlockStatement
}

// interface WhileStatement <: Statement
type WhileStatement struct {
	Node
	Test Expression
	Body Statement
}

func (*WhileStatement) Statement() {}

// interface DoWhileStatement <: Statement
type DoWhileStatement struct {
	Node
	Test Expression
	Body Statement
}

func (*DoWhileStatement) Statement() {}

type forStatementInit interface {
	forStatementInit()
}

// interface ForStatement <: Statement
type ForStatement struct {
	Node
	Init *forStatementInit
	Test *Expression
	Update *Expression
	Body Statement
}

func (*ForStatement) Statement() {}

type forInStatementLeft interface {
	forInStatementLeft()
}

// interface ForInStatement <: Statement
type ForInStatement struct {
	Node
	Left forInStatementLeft
	Right Expression
	Body Statement
}

func (*ForInStatement) Statement() {}

type Declaration interface {
	Statement
}

// interface FunctionDeclaration <: Function, Declaration
type FunctionDeclaration struct {
	Function
}

func (*FunctionDeclaration) Statement() {}

// interface VariableDeclaration <: Declaration
type VariableDeclaration struct {
	Node
	Declarations []VariableDeclarator
	Kind string
}

func (*VariableDeclaration) Statement() {}

func (*VariableDeclaration) forStatementInit() {}

func (*VariableDeclaration) forInStatementLeft() {}

// interface VariableDeclarator <: Node
type VariableDeclarator struct {
	Node
	Id Pattern
	Init *Expression
}

type Expression interface {
	forStatementInit
	Expression()
}

type Pattern interface {
	forInStatementLeft
	Pattern()
}