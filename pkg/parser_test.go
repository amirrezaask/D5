package D5

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	p := NewParser()
	input := 12
	expected := NumberExpr(12)
	e, err := p.Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, e)
}

func TestParseString(t *testing.T) {
	p := NewParser()
	input := "D5 is new lisp"
	expected := StringExpr("D5 is new lisp")
	e, err := p.Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, e)
}

func TestParseBool(t *testing.T) {
	p := NewParser()
	input := true
	expected := BoolExpr(true)
	e, err := p.Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, e)
}

func TestParseMap(t *testing.T) {
	p := NewParser()
	input := Table{
		"1":    2,
		"name": "amirreza",
	}
	expected := TableExpr{
		"1":    NumberExpr(2),
		"name": StringExpr("amirreza"),
	}
	e, err := p.Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, e)
}
func readJson(bs []byte) (Table, error) {
	m := Table{}
	err := json.Unmarshal(bs, &m)
	return m, err
}
func TestParseComplexProgram(t *testing.T) {
	p := NewParser()
	input := []byte(`{
		"if": {
			"then": "amirreza",
			"else": "amirreza2"
		}
	}`)
	table, err := readJson(input)
	assert.NoError(t, err)
	e, err := p.Parse(table)
	_ = e
	assert.NoError(t, err)
}
