package api

import "github.com/consensys/gnark/frontend"

var Api frontend.API

func Add(a, b frontend.Variable) frontend.Variable {
	return Api.Add(a, b)
}

func Sub(a, b frontend.Variable) frontend.Variable {
	return Api.Sub(a, b)
}

func Mul(a, b frontend.Variable) frontend.Variable {
	return Api.Mul(a, b)
}

func Compiler() frontend.Compiler {
	return Api.Compiler()
}

func AssertIsEqual(a, b frontend.Variable) {
	Api.AssertIsEqual(a, b)
}
