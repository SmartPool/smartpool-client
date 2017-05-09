package ethereum

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func BuildExtraData(address common.Address, diff *big.Int) string {
	// id = address % (26+26+10)**11
	base := big.NewInt(0)
	base.Exp(big.NewInt(62), big.NewInt(11), nil)
	id := big.NewInt(0)
	id.Mod(address.Big(), base)
	return fmt.Sprintf("SmartPool-%s%s", smartpool.BigToBase62(id), smartpool.BigToBase62(diff))
}
