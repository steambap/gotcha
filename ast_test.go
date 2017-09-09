package gotcha

import (
	"testing"
)

func TestInterfaceCompile(t *testing.T) {
	var _ Expression = &Identifier{}
	var _ Pattern = &Identifier{}
	var _ Statement = &ExpressionStatement{}
	var _ Statement = &BlockStatement{}
	var _ Statement = &EmptyStatement{}
	var _ Statement = &DebuggerStatement{}
	var _ Statement = &WithStatement{}
	var _ Statement = &ReturnStatement{}
	var _ Statement = &LabeledStatement{}
	var _ Statement = &BreakStatement{}
	var _ Statement = &ContinueStatement{}
	var _ Statement = &IfStatement{}
	var _ Statement = &SwitchStatement{}
	var _ Statement = &ThrowStatement{}
	var _ Statement = &TryStatement{}
	var _ Statement = &WhileStatement{}
	var _ Statement = &DoWhileStatement{}
	var _ Statement = &ForStatement{}
	var _ Declaration = &FunctionDeclaration{}
	var _ Declaration = &VariableDeclaration{}
}
