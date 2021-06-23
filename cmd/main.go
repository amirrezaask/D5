package main

import (
	D5 "D5/pkg"
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
	t := D5.Map{}
	err = json.Unmarshal(bs, &t)
	if err != nil {
		panic(err)
	}
	p := D5.NewParser()
	expr, err := p.Parse(t)
	if err != nil {
		panic(err)
	}
	// spew.Dump(expr)
	evaluator := D5.Evaluator{}
	fmt.Println(evaluator.Eval(expr))
}
