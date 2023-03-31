// generated from field.csgo

package runtime

import (
	"github.com/argennon-project/csgo/internal/api"
)

import "math/big"

func FieldOrder() *big.Int {
	return api.Compiler().Field()
}
