package D5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalNumber(t *testing.T) {
	input := Number(1.2)
	e := &Evaluator{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}

func TestEvalString(t *testing.T) {
	input := String("amirreza")
	e := &Evaluator{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
func TestEvalBool(t *testing.T) {
	input := Bool(true)
	e := &Evaluator{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
func TestEvalSimpleTable(t *testing.T) {
	input := Table{
		"name": String("amireza"),
	}
	e := &Evaluator{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
