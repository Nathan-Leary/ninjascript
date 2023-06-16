package ninjascript

import (
	math_big "math/big"
)

func init() {
	if _, ok := Api["math/big"]; !ok {
		Api["math/big"] = map[string]interface{}{}
	}
	// Api["math/big"]["NewT"] = math_big.NewT
	Api["math/big"]["MaxExp"] = math_big.MaxExp
	Api["math/big"]["MinExp"] = math_big.MinExp
	Api["math/big"]["MaxPrec"] = math_big.MaxPrec
	Api["math/big"]["Jacobi"] = math_big.Jacobi
	Api["math/big"]["Below"] = math_big.Below
	Api["math/big"]["Exact"] = math_big.Exact
	Api["math/big"]["Above"] = math_big.Above
	Api["math/big"]["ErrNaN"] = math_big.ErrNaN{}
	Api["math/big"]["Float"] = math_big.Float{}
	Api["math/big"]["NewFloat"] = math_big.NewFloat
	Api["math/big"]["ParseFloat"] = math_big.ParseFloat
	Api["math/big"]["Int"] = math_big.Int{}
	Api["math/big"]["NewInt"] = math_big.NewInt
	Api["math/big"]["Rat"] = math_big.Rat{}
	Api["math/big"]["NewRat"] = math_big.NewRat
	Api["math/big"]["ToNearestEven"] = math_big.ToNearestEven
	Api["math/big"]["ToNearestAway"] = math_big.ToNearestAway
	Api["math/big"]["ToZero"] = math_big.ToZero
	Api["math/big"]["AwayFromZero"] = math_big.AwayFromZero
	Api["math/big"]["ToNegativeInf"] = math_big.ToNegativeInf
	Api["math/big"]["ToPositiveInf"] = math_big.ToPositiveInf
	Api["math/big"]["Jacobi"] = math_big.Jacobi
	Api["math/big"]["NewFloat"] = math_big.NewFloat
	Api["math/big"]["ParseFloat"] = math_big.ParseFloat
	Api["math/big"]["NewInt"] = math_big.NewInt
	Api["math/big"]["NewRat"] = math_big.NewRat

}
