package D5

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	p := NewParser()
	input := 12
	expected := Number(12)
	e, err := p.Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, e)
}

func TestParseString(t *testing.T) {
	p := NewParser()
	input := "D5 is new lisp"
	expected := String("D5 is new lisp")
	e, err := p.Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, e)
}

func TestParseBool(t *testing.T) {
	p := NewParser()
	input := true
	expected := Bool(true)
	e, err := p.Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, e)
}

func TestParseMap(t *testing.T) {
	p := NewParser()
	input := map[string]interface{}{
		"1":    2,
		"name": "amirreza",
	}
	expected := Table{
		"1":    Number(2),
		"name": String("amirreza"),
	}
	e, err := p.Parse(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, e)
}
func readJson(bs []byte) (map[string]interface{}, error) {
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
