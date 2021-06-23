package D5

import (
	"errors"
	"fmt"
	"reflect"
)

type Node interface {
	Type() string
	Value() interface{}
}
type Number float64

func (n Number) Type() string {
	return "number"
}
func (n Number) Value() interface{} {
	return n
}

type String string

func (n String) Type() string {
	return "string"
}
func (n String) Value() interface{} {
	return n
}

type Map map[string]interface{}

func (n Map) Type() string {
	return "map"
}
func (n Map) Value() interface{} {
	return n
}

type Array []interface{}

func (a Array) Type() string {
	return "array"
}
func (a Array) Value() interface{} {
	return a
}

type Bool struct {
    value bool
}

func (n Bool) Type() string {
	return "bool"
}
func (n Bool) Value() interface{} {
	return n.value
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) parseMaps(src interface{}) (Map, error) {
	switch src.(type) {
	case Map:
		newTable := Map{}
		for k, v := range src.(Map) {
			exprVal, err := p.Parse(v)
			if err != nil {
				return nil, err
			}
			newTable[k] = exprVal
		}
		return newTable, nil
	case map[interface{}]interface{}:
		newTable := Map{}
		for k, v := range src.(map[interface{}]interface{}) {
			exprVal, err := p.Parse(v)
			if err != nil {
				return nil, err
			}
			newTable[fmt.Sprint(k)] = exprVal
		}
		return newTable, nil

	case map[string]interface{}:
		newTable := Map{}
		for k, v := range src.(map[string]interface{}) {
			exprVal, err := p.Parse(v)
			if err != nil {
				return nil, err
			}
			newTable[k] = exprVal
		}
		return newTable, nil
	default:
		return nil, errors.New("Map type did not match")

	}
}

func (p *Parser) parseArray(src interface{}) (Node, error) {
	switch src.(type) {
	case Array:
		newArray := Array{}
		for _, v := range src.(Array) {
			exprVal, err := p.Parse(v)
			if err != nil {
				return nil, err
			}
			newArray = append(newArray, exprVal)
		}
		return newArray, nil
	case []interface{}:
		newArray := Array{}
		for _, v := range src.([]interface{}) {
			exprVal, err := p.Parse(v)
			if err != nil {
				return nil, err
			}
			newArray = append(newArray, exprVal)
		}
		return newArray, nil
	default:
		return nil, errors.New("Array type did not match")
	}

}

func (p *Parser) parseComplexDataStructure(src interface{}) (Node, error) {
	t := reflect.TypeOf(src)
	switch k := t.Kind(); k {
	case reflect.Map:
		return p.parseMaps(src)
	case reflect.Slice:
		return p.parseArray(src)
	}
	return nil, errors.New("No Type matched")
}

func (p *Parser) Parse(src interface{}) (Node, error) {
	switch src.(type) {
	case int:
		return Number(src.(int)), nil
	case float32:
		return Number(src.(float32)), nil
	case float64:
		return Number(src.(float64)), nil
	case string:
		return String(src.(string)), nil
	case bool:
		return Bool{src.(bool)}, nil
	default:
		return p.parseComplexDataStructure(src)
	}
}
