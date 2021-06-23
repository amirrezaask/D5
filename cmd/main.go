package main

import (
	D5 "github.com/amirrezaask/D5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("need file name")
	}
	filename := os.Args[1]
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
    b := D5.Block{}
	err = json.Unmarshal(bs, &b)
	if err != nil {
		panic(err)
	}
	// spew.Dump(expr)
	evaluator := D5.Interpreter{}
	fmt.Println(evaluator.Eval(b))
}
