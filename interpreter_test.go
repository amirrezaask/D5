package D5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalNumber(t *testing.T) {
	input := 1.2
	e := &Interpreter{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}

func TestEvalString(t *testing.T) {
	input := "amirreza"
	e := &Interpreter{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
func TestEvalBool(t *testing.T) {
	input := true
	e := &Interpreter{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
func TestEvalSimpleTable(t *testing.T) {
	input := Block{
		"name": "amireza",
	}
	e := &Interpreter{}
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}

func _if(cond interface{}, then interface{}, _else interface{}) Block {
	return Block{"type": "if", "condition": cond, "then": then, "else": _else}
}

func TestEvalIf(t *testing.T) {
	input := _if(true, 1, 2)
	e := NewInterpreter()
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, 1, output)
}

func TestEvalPutInMap(t *testing.T) {
	input := Block{
		"type": "put",
		"to":    Block{},
		"key":   "mamad",
		"value": 1,
	}
	e := NewInterpreter()
	output, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t, Block{"mamad": 1}, output)
}

func TestEvalPutInState(t *testing.T) {
	input := Block{
		"type": "put",
		"to":    "state",
		"key":   "mamad",
		"value": 1,
	}
	e := NewInterpreter()
	_, err := e.Eval(input)
	assert.NoError(t, err)
	assert.Equal(t,1, e.state["mamad"])
}