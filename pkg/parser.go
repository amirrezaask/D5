package pkg

type Expr interface {
	Eval(out interface{}) error
}

type NumberExpr float64
type StringExpr string
type TableExpr map[interface{}]interface{}
type BoolExpr bool

// TOOD: couroutine expr
