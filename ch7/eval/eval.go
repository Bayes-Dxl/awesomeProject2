//package eval
//
//import (
//	"fmt"
//	"math"
//)
//
//type Var string
//type literal float64
//
//type unary struct {
//	op rune // '+','-'中的一个
//	x  Expr
//}
//
//type binary struct {
//	op   rune
//	x, y Expr
//}
//
//type call struct {
//	fn   string
//	args []Expr
//}
//
//type Env map[Var]float64
//
//type Expr interface {
//	Eval(env Env) float64
//}
//
//func (v Var) Eval(env Env) float64 {
//	return env[v]
//}
//
//func (l literal) Eval(_ Env) float64 {
//	return float64(l)
//}
//
//func (u unary) Eval(env Env) float64 {
//	switch u.op {
//	case '+':
//		return +u.x.Eval(env)
//	case '-':
//		return -u.x.Eval(env)
//	}
//	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
//}
//
//func (b binary) Eval(env Env) float64 {
//	switch b.op {
//	case '+':
//		return b.x.Eval(env) + b.y.Eval(env)
//	case '-':
//		return b.x.Eval(env) - b.y.Eval(env)
//	case '*':
//		return b.x.Eval(env) * b.y.Eval(env)
//	case '/':
//		return b.x.Eval(env) / b.y.Eval(env)
//	}
//	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
//}
//
//func (c call) Eval(env Env) float64 {
//	switch c.fn {
//	case "pow":
//		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
//	case "sin":
//		return math.Sin(c.args[0].Eval(env))
//	case "sqrt":
//		return math.Sqrt(c.args[0].Eval(env))
//	}
//	panic(fmt.Sprintf("unsupported binary operator: %q", c.fn))
//}
//
////func TestEval(t *testing.T) {
////	tests := []struct {
////		expr string
////		env  Env
////		want string
////	}{
////		{"sqrt(A/pi)", Env{"A": 87628, "pi": math.Pi}, "167"},
////		{"pow(x,3)+pow(y,3)", Env{"x": 12, "y": 1}, "1729"},
////		{"5 / 9 * (F - 32 )", Env{"F": -40}, "-40"},
////		{"5 / 9 * (F - 32 )", Env{"F": 32}, "0"},
////		{"5 / 9 * (F - 32 )", Env{"F": 212}, "100"},
////	}
////
////	var preExpr string
////	for _, test := range tests {
////		if test.expr != preExpr {
////			fmt.Printf("\n%s\n", test.expr)
////			preExpr = test.expr
////		}
////		expr, err := Parse(test.expr)
////		if err != nil {
////			t.Error(err) //解析出错
////			continue
////		}
////		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
////		fmt.Printf("\t%v => %s\n ", test.env, got)
////		if got != test.want {
////			t.Errorf("%s.Eval() in %v = %q, want %q\n",
////				test.expr, test.env, got, test.want)
////		}
////	}
////
////}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//!-Eval2
