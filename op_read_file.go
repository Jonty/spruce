package main

import (
    "io/ioutil"
    "strings"
)

type ReadFileOperator struct{}

func (ReadFileOperator) Setup() error {
        return nil
}

func (ReadFileOperator) Phase() OperatorPhase {
        return EvalPhase
}

func (ReadFileOperator) Dependencies(_ *Evaluator, _ []*Expr, _ []*Cursor) []*Cursor {
        return []*Cursor{}
}

func (s ReadFileOperator) Run(ev *Evaluator, args []*Expr) (*Response, error) {
    Filename := args[0].Literal.(string)
    FileContents, _ := ioutil.ReadFile(Filename)

	return &Response{
		Type:  Replace,
		Value: strings.TrimSpace(string(FileContents)),
	}, nil
}

func init() {
	RegisterOp("read_file", ReadFileOperator{})
}
