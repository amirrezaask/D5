package D5

import (
	"errors"
	"fmt"
)

type Backend interface {
	Eval(n Node, out interface{}) error
}

type Evaluator struct {
	state map[string]interface{}
}

func (e *Evaluator) evalString(s String) interface{} {
	return s
}

func (e *Evaluator) evalIf(b Map) (interface{}, error) {
	cond := b["condition"]
	then := b["then"]
	_else := b["else"]
	condVal, err := e.Eval(cond.(Node))
	if err != nil {
		return nil, err
	}
	if condVal.(bool) {
		return e.Eval(then.(Node))
	} else {
		return e.Eval(_else.(Node))
	}
}

var awesomePrinter = func(obj interface{}) {
	fmt.Printf("%+v", obj)
}

func (e *Evaluator) evalTable(m Map) (interface{}, error) {
	typ, exists := m["type"]
	if !exists {
		return nil, fmt.Errorf("we need a type key")
	}
    fmt.Println(typ)
	switch typ {
	case String("if"):
		return e.evalIf(m)
    case String("value"):
        return m["value"], nil
	}
	return m, nil
}
func (e *Evaluator) Eval(n Node) (interface{}, error) {
	switch n.Type() {
	case "number", "bool":
		return n.Value(), nil
	case "map":
		return e.evalTable(n.Value().(Map))
	case "string":
		return e.evalString(n.Value().(String)), nil
	}
	return nil, errors.New("no type matched")
}
