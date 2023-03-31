// generated from field.csgo

package runtime

import (
	"github.com/argennon-project/csgo/transpiled/internal/api"
)

import "math/big"

func FieldOrder() *big.Int {
	return api.Compiler().Field()
}
