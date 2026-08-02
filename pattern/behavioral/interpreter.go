package behavioral

// 解释器模式（Interpreter）：给定一个语言，定义它的文法的一种表示，并定义一个解释器，
// 这个解释器使用该表示来解释语言中的句子。
//
// 典型场景：简单规则引擎、表达式求值。下面实现一个最小布尔“与”表达式解释器。

// Expression 抽象表达式。
type Expression interface {
	Interpret() bool
}

// TerminalExpression 终结符：一个布尔变量。
type TerminalExpression struct {
	value bool
}

func NewTerminal(value bool) TerminalExpression { return TerminalExpression{value: value} }
func (e TerminalExpression) Interpret() bool    { return e.value }

// AndExpression 非终结符：与运算。
type AndExpression struct {
	left, right Expression
}

func NewAnd(left, right Expression) AndExpression { return AndExpression{left: left, right: right} }
func (e AndExpression) Interpret() bool           { return e.left.Interpret() && e.right.Interpret() }

// OrExpression 非终结符：或运算。
type OrExpression struct {
	left, right Expression
}

func NewOr(left, right Expression) OrExpression { return OrExpression{left: left, right: right} }
func (e OrExpression) Interpret() bool          { return e.left.Interpret() || e.right.Interpret() }
