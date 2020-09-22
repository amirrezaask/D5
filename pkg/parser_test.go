package pkg

import (
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
