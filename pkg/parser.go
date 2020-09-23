package D5

import (
	"errors"
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

type Table map[string]interface{}

func (n Table) Type() string {
	return "table"
}
func (n Table) Value() interface{} {
	return n
}

type Bool bool

func (n Bool) Type() string {
	return "bool"
}
func (n Bool) Value() interface{} {
	return n
}

type Parser struct {
	CompTimeCtx map[string]interface{}
}

func NewParser() *Parser {
	return &Parser{}
}
func (p *Parser) parseMaps(src interface{}) (Table, error) {
	switch src.(type) {
	case map[string]interface{}:
		newTable := Table{}
		for k, v := range src.(map[string]interface{}) {
			exprVal, err := p.Parse(v)
			if err != nil {
				return nil, err
			}
			newTable[k] = exprVal
		}
		return newTable, nil
	default:
		return nil, errors.New("In tables keys should be string")

	}
}

func (p *Parser) parseComplexDataStructure(src interface{}) (Table, error) {
	t := reflect.TypeOf(src)
	switch k := t.Kind(); k {
	case reflect.Map:
		return p.parseMaps(src)
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
		return Bool(src.(bool)), nil
	default:
		return p.parseComplexDataStructure(src)
	}
}
