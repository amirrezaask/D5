package D5

import (
	"fmt"
)

type Backend interface {
	Eval(src interface{}) error
}

type Block = map[string]interface{}

type Interpreter struct {
	state map[string]interface{}
}

func NewInterpreter() *Interpreter {
	return &Interpreter{state: make(map[string]interface{})}
}

func (e *Interpreter) evalIf(b Block) (interface{}, error) {
	cond := b["condition"]
	then := b["then"]
	_else := b["else"]
	condVal, err := e.Eval(cond)
	if err != nil {
		return nil, err
	}
	if condVal.(bool) {
		return e.Eval(then)
	} else {
		return e.Eval(_else)
	}
}

func (e *Interpreter) evalPut(b Block) (interface{}, error) {
	key, err := e.Eval(b["key"])
	if err != nil {
		return nil, err
	}
	value, err := e.Eval(b["value"])
	if err != nil {
		return nil, err
	}
	to, err := e.Eval(b["to"])
	if err != nil {
		return nil, err
	}

	switch to.(type) {
	case string:
		if to.(string) == "state" {
			e.state[fmt.Sprint(key)] = value
			return nil, nil
		}
	case Block:
		to.(Block)[fmt.Sprint(key)] = value
		return to, nil
	}
	return nil, fmt.Errorf("not supported argument for to: %T", to)
}

func (e *Interpreter) evalGet(b Block) (interface{}, error) {
	from, err := e.Eval(b["from"])
	if err != nil {
		return nil, err
	}
	key, err := e.Eval(b["key"])
	if err != nil {
		return nil, err
	}
	switch from.(type) {
	case string:
		if from.(string) == "state" {
			return b[key.(string)], nil
		}
		return nil, fmt.Errorf("only state is supported for now")
	case Block:
		return from.(Block)[key.(string)], nil
	default:
		return nil, fmt.Errorf("we dont support %T as from argument", from)
	}
}

func (e *Interpreter) Eval(src interface{}) (interface{}, error) {
	switch src.(type) {
	case Block:
		s := src.(Block)
		typ, exists := s["type"]
		if !exists {
			return s, nil
		}
		switch typ {
		case "if":
			return e.evalIf(s)
		case "value":
			return s["value"], nil
		case "get":
			return e.evalGet(s)
		case "put":
			return e.evalPut(s)
		}
	case int, float64, string, bool:
		return src, nil
	}

	return nil, fmt.Errorf("no type matched: %T", src)
}
