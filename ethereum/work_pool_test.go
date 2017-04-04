package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

func newTestWork() *Work {
	json := "{\"parentHash\":\"0x1fa6e98e1ad7dc1755115714498158b5e550c9d562c3babe1275c04d8700cb45\",\"sha3Uncles\":\"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347\",\"miner\":\"0xeb69b29551f5830581a29858d1aca0e39ec14d57\",\"stateRoot\":\"0x199d756fda16fcc46fe4d9dd8b16c7eec8f30f1dd0ecd4779f183579d02596a3\",\"transactionsRoot\":\"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421\",\"receiptsRoot\":\"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421\",\"logsBloom\":\"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\",\"difficulty\":\"0x38ab2a\",\"number\":\"0x102\",\"gasLimit\":\"0x47e890\",\"gasUsed\":\"0x0\",\"timestamp\":\"0x58c1be42\",\"extraData\":\"0x536d617274506f6f6c2d534b787a4a3859666c38735530713030303030303030\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"nonce\":\"0x0000000000000000\"}"
	header := &types.Header{}
	header.UnmarshalJSON([]byte(json))
	return NewWork(
		header,
		"0x958c07589d805f63393997ea0baca55a3101baa2c6616e3a5c8d51baf590d6ec",
		"0x0000000000000000000000000000000000000000000000000000000000000000",
		big.NewInt(100000),
		"0x958c07589d805f63393997ea0baca55a3101baa2c6616e3a5c8d51baf590d6ec",
	)
}

func newTestSolution() *Solution {
	nonce, _ := hexutil.DecodeUint64("0xacc04b3a085f4609")
	return &Solution{
		Nonce:     types.EncodeNonce(nonce),
		Hash:      common.HexToHash("0x958c07589d805f63393997ea0baca55a3101baa2c6616e3a5c8d51baf590d6ec"),
		MixDigest: common.HexToHash("0x39fb1956275f9308fb7bec594108520e9e67d213cd1711874cf02d8450ea171e"),
	}
}

func TestWorkPoolAcceptSolution(t *testing.T) {
	w := newTestWork()
	s := newTestSolution()
	wp := WorkPool{
		w.ID(): w,
	}
	if wp.AcceptSolution(s) == nil {
		t.Fail()
	}
}

func TestWorkPoolDoesntAcceptSolution(t *testing.T) {
	wp := WorkPool{}
	s := newTestSolution()
	if wp.AcceptSolution(s) != nil {
		t.Fail()
	}
}

func TestWorkPoolAddWorkByItsID(t *testing.T) {
	wp := WorkPool{}
	w := newTestWork()
	wp.AddWork(w)
	if wp[w.ID()] == nil {
		t.Fail()
	}
}
