package evidence

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/gracenoah/tendermint/crypto/encoding/amino"
	"github.com/gracenoah/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterEvidenceMessages(cdc)
	cryptoAmino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

// For testing purposes only
func RegisterMockEvidences() {
	types.RegisterMockEvidences(cdc)
}
