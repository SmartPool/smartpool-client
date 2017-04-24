package smartpool

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

const (
	HashLength          = 16
	WordLength          = 128
	BranchElementLength = 32
)

type (
	Word          [WordLength]byte
	SPHash        [HashLength]byte
	BranchElement [BranchElementLength]byte
)

func BytesToBig(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)

	return n
}

func rev(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func (w Word) ToUint256Array() []*big.Int {
	result := []*big.Int{}
	for i := 0; i < WordLength/32; i++ {
		z := big.NewInt(0)
		// reverse the bytes because contract expects
		// big Int is construct in little endian
		z.SetBytes(rev(w[i*32 : (i+1)*32]))
		result = append(result, z)
	}
	return result
}

func base62DigitsToString(digits []byte) string {
	for i := 0; i < len(digits); i++ {
		if 0 <= digits[i] && digits[i] <= 9 {
			digits[i] += 48
		} else {
			if 10 <= digits[i] && digits[i] <= 9+26 {
				digits[i] += 97 - 10
			} else {
				if 9+26+1 <= digits[i] && digits[i] <= 9+26+26 {
					digits[i] += 65 - 36
				}
			}
		}
	}
	return string(digits)
}

// return 11 chars base 62 representation of a big int
// base chars are 0-9 a-z A-Z
func BigToBase62(num *big.Int) string {
	digits := []byte{}
	n := big.NewInt(0)
	n.Add(n, num)
	zero := big.NewInt(0)
	base := big.NewInt(62)
	for {
		mod := big.NewInt(0)
		n, mod = n.DivMod(n, base, mod)
		mBytes := mod.Bytes()
		if len(mBytes) == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, mod.Bytes()[0])
		}
		if n.Cmp(zero) == 0 {
			break
		}
	}
	l := len(digits)
	for i := 0; i < 11-l; i++ {
		digits = append(digits, 0)
	}
	return base62DigitsToString(digits)
}

func (h SPHash) Str() string   { return string(h[:]) }
func (h SPHash) Bytes() []byte { return h[:] }
func (h SPHash) Big() *big.Int { return BytesToBig(h[:]) }
func (h SPHash) Hex() string   { return hexutil.Encode(h[:]) }

func (h BranchElement) Str() string   { return string(h[:]) }
func (h BranchElement) Bytes() []byte { return h[:] }
func (h BranchElement) Big() *big.Int { return BytesToBig(h[:]) }
func (h BranchElement) Hex() string   { return hexutil.Encode(h[:]) }

func BranchElementFromHash(a, b SPHash) BranchElement {
	result := BranchElement{}
	copy(result[:], append(a[:], b[:]...)[:BranchElementLength])
	return result
}

// Rig represents mining actor that will be directly interactive with smartpool.
type Rig interface {
	ID() string
}

// Work represents SmartPool work that miner needs to solve to have valid
// shares. Work is easier (has smaller difficulty) than actual Ethereum
// work that the miner can get from the network.
type Work interface {
	ID() string
	// AcceptSolution takes solution to construct and return a Share representing
	// the solution that came from miner.
	AcceptSolution(sol Solution) Share
}

// Solution represents a solution for a work
type Solution interface {
	// WorkID returns the ID to identify the work it is trying to solve
	WorkID() string
}

// Share represent a solution of a Work that comes from the miner.
type Share interface {
	// Counter returns the counter to be used in augmented merkle tree of a claim
	// which contains many shares. This counter must be increasing as shares
	// share coming. In other words, later share must have bigger counter.
	Counter() *big.Int
	// Difficulty returns the difficulty of the share that miner has solved.
	ShareDifficulty() *big.Int
	// Hash return the hash of the share to be used as leaf hash of the augmented
	// merkle tree.
	Hash() SPHash
	// FullSolution returns true if the share is solution for its full block pow
	// hash. It returns false otherwise.
	FullSolution() bool
}

// Claim represent a batch of shares which needs to reorganize its shares in
// ascending order of share counter.
type Claim interface {
	// NumShares returns number of shares that the claim is holding
	NumShares() *big.Int
	GetShare(index int) Share
	// Difficulty returns the min difficulty across all of its shares
	Difficulty() *big.Int
	// Min returns the min counter of the augmented merkle root
	Min() *big.Int
	// Max returns the max counter of the augmented merkle root
	Max() *big.Int
	// AugMerkle returns the hash of the augmented merkle root
	AugMerkle() SPHash
	// SetEvidence sets the share index to be used to prove the claim
	SetEvidence(shareIndex *big.Int)
	// CounterBranch returns array of counters in proof branch of the share
	CounterBranch() []*big.Int
	// HashBranch returns array of hashes in proof branch of the share
	HashBranch() []*big.Int
}
