package D5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalNumber(t *testing.T) {
	input := Number(1.2)
	e := &Interpreter{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}

func TestEvalString(t *testing.T) {
	input := String("amirreza")
	e := &Interpreter{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
func TestEvalBool(t *testing.T) {
	input := Bool(true)
	e := &Interpreter{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
func TestEvalSimpleTable(t *testing.T) {
	input := Map{
		"name": String("amireza"),
	}
	e := &Interpreter{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
