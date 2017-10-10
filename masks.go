package gotcha

type Context int

const (
	NoContext           Context = 0
	Module                      = 1 << 0
	Strict                      = 1 << 1
	AllowIn                     = 1 << 2
	SimpleArrow                 = 1 << 3
	YieldContext                = 1 << 4
	AwaitContext                = 1 << 5
	InParenthesis               = 1 << 6
	InParameter                 = 1 << 7
	ArrowParameterList          = 1 << 8
	StatementContext            = 1 << 9
	SimpleParameterList         = 1 << 10

	AnnexB             = 1 << 11
	Export             = 1 << 12
	OptionalIdentifier = 1 << 13
	Let                = 1 << 14
	Const              = 1 << 15

	Lexical = Let | Const
)

type Flags int

const (
	NoFlags        Flags = 0
	LineTerminator       = 1 << 0
	HasUnicode           = 1 << 1
	InFunctionBody       = 1 << 2
	AllowCall            = 1 << 3

	/* Numeric */
	Noctal   = 1 << 4 // e.g. `0777`
	BigInt   = 1 << 5 // e.g. `100n`
	Float    = 1 << 6 // e.g. `09.01`
	Exponent = 1 << 7 // e.g. `10e2`

	/* Options */
	OptionsRanges    = 1 << 8
	OptionsLoc       = 1 << 9
	OptionsSource    = 1 << 10
	OptionsJSX       = 1 << 11
	OptionsRaw       = 1 << 12
	OptionsNext      = 1 << 13
	OptionsOnComment = 1 << 14
	OptionsOnToken   = 1 << 15
	OptionsV8        = 1 << 16

	// BigInt implementation can't handle either float or exponent acc. TC-39
	FloatOrExponent = Float | Exponent
)

type ParenthesizedState int

const (
	NoParenthesized ParenthesizedState = 0
	EvalOrArg                          = 1 << 0 // If (async) arrow contains eval or agruments
	Yield                              = 1 << 1 // If (async) arrow contains eval or agruments
	Await                              = 1 << 2 // If async arrow contains 'await'
	Parenthesized                      = 1 << 3 // Tracks invalid parenthesized pattern
)

type AsyncState int

const (
	NoAsyncState AsyncState = iota
	FunctionAsyncState
	IdentifierAsyncState
)
