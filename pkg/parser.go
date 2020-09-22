package pkg

import (
	"errors"
	"reflect"
)

type Expr interface {
	Eval(out interface{}) error
}

type NumberExpr float64
type StringExpr string
type TableExpr map[interface{}]interface{}
type BoolExpr bool

func (e NumberExpr) Eval(out interface{}) error {
	out = e
	return nil
}

func (e StringExpr) Eval(out interface{}) error {
	out = e
	return nil
}

func (e TableExpr) Eval(out interface{}) error {
	out = e
	return nil
}

func (e BoolExpr) Eval(out interface{}) error {
	out = e
	return nil
}

type Parser struct {
	CompTimeCtx map[string]interface{}
}

func NewParser() *Parser {
	return &Parser{}
}

func parseMaps(src interface{}) (Expr, error) {
	switch src.(type) {
	default:
		return nil, errors.New("In tables keys should be interface{}")
	case map[interface{}]interface{}:
		return TableExpr(src.(map[interface{}]interface{})), nil
	}
}

func parseComplexDataStructure(src interface{}) (Expr, error) {
	t := reflect.TypeOf(src)
	switch k := t.Kind(); k {
	case reflect.Map:
		return parseMaps(src)
	}
	return nil, errors.New("No Type matched")
}

func (p *Parser) Parse(src interface{}) (Expr, error) {
	switch src.(type) {
	case int:
		return NumberExpr(src.(int)), nil
	case float32:
		return NumberExpr(src.(float32)), nil
	case float64:
		return NumberExpr(src.(float64)), nil
	case string:
		return StringExpr(src.(string)), nil
	case bool:
		return BoolExpr(src.(bool)), nil
	default:
		return parseComplexDataStructure(src)
	}
}
