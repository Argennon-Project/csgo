// generated from binary.csgo

package convert

import (
	"github.com/argennon-project/csgo/internal/api"
	"github.com/consensys/gnark/frontend"
)

import "github.com/consensys/gnark/std/math/bits"

// AssertBitLen ensures that the binary representation of x has less than bitLen bits. It assumes that x is an unsigned
// number between 0 and P - 1, where P is the order of the underlying field.
func AssertBitLen(bitLen int, x frontend.Variable) {
	bits.ToBinary(api.Api, x, bits.WithNbDigits(bitLen))
}
