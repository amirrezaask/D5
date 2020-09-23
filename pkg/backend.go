package D5

import "errors"

type Backend interface {
	Eval(n Node, out interface{}) error
}

type Evaluator struct {
	state map[string]interface{}
}

func (e *Evaluator) evalString(s String) interface{} {
	return s
}
func (e *Evaluator) evalTable(t Table) (interface{}, error) {
	_, exists := t["type"]
	if exists {
		//should be evaluated
	}
	return t, nil
}
func (e *Evaluator) Eval(n Node) (interface{}, error) {
	switch n.Type() {
	case "number", "bool":
		return n.Value(), nil
	case "table":
		return e.evalTable(n.Value().(Table))
	case "string":
		return e.evalString(n.Value().(String)), nil
	}
	return nil, errors.New("no type matched")
}
