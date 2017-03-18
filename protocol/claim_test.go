package protocol

import (
	"github.com/SmartPool/smartpool-client"
	"math/big"
	"reflect"
	"sort"
	"testing"
)

func newTestShares() Shares {
	return Shares{
		&testShare{c: big.NewInt(10), d: big.NewInt(10), h: 1},
		&testShare{c: big.NewInt(1), d: big.NewInt(4), h: 2},
		&testShare{c: big.NewInt(5), d: big.NewInt(5), h: 3},
		&testShare{c: big.NewInt(8), d: big.NewInt(9), h: 4},
	}
}

func TestSharesSortSharesAscendingCounter(t *testing.T) {
	s := newTestShares()
	sort.Sort(s)
	for i, _ := range s[:] {
		if i > 0 && s[i].Counter().Cmp(s[i-1].Counter()) < 0 {
			t.Fail()
		}
	}
}

func TestClaimNumSharesReturnsNumberOfShare(t *testing.T) {
	s := &Claim{Shares{
		&testShare{big.NewInt(10), big.NewInt(10), 0}},
		nil,
		nil,
		nil,
	}
	if s.NumShares().Int64() != 1 {
		t.Fail()
	}
}

func TestClaimAddShare(t *testing.T) {
	s := NewClaim()
	s.AddShare(&testShare{big.NewInt(10), big.NewInt(10), 0})
	if s.NumShares().Int64() != 1 {
		t.Fail()
	}
}

func TestClaimDifficultyReturnMinDifficultyOfItsShares(t *testing.T) {
	s := &Claim{newTestShares(), nil, nil, nil}
	if s.Difficulty().Cmp(big.NewInt(4)) != 0 {
		t.Fail()
	}
}

func TestClaimMinReturnMinCounterOfItsShares(t *testing.T) {
	s := &Claim{newTestShares(), nil, nil, nil}
	if s.Min().Cmp(big.NewInt(1)) != 0 {
		t.Fail()
	}
}

func TestClaimMinReturnMaxCounterOfItsShares(t *testing.T) {
	s := &Claim{newTestShares(), nil, nil, nil}
	if s.Max().Cmp(big.NewInt(10)) != 0 {
		t.Fail()
	}
}

func TestClaimAugMerkle(t *testing.T) {
	s := &Claim{newTestShares(), nil, nil, nil}
	result := smartpool.SPHash{151, 170, 188, 143, 116, 177, 57, 134, 27, 7, 73, 156, 173, 26, 237, 226}
	if !reflect.DeepEqual(s.AugMerkle(), result) {
		t.Fail()
	}
}
