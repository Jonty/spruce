package main

import (
    "math/rand"
    "time"
)

type RandStringOperator struct{}

func (RandStringOperator) Setup() error {
        return nil
}

func (RandStringOperator) Phase() OperatorPhase {
        return EvalPhase
}

func (RandStringOperator) Dependencies(_ *Evaluator, _ []*Expr, _ []*Cursor) []*Cursor {
        return []*Cursor{}
}

func (s RandStringOperator) Run(ev *Evaluator, args []*Expr) (*Response, error) {
    Amount := args[0].Literal.(int64)

	return &Response{
		Type:  Replace,
		Value: RandString(Amount),
	}, nil
}

// Taken from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandString(n int64) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func init() {
    rand.Seed(time.Now().UnixNano())
	RegisterOp("randstring", RandStringOperator{})
}
