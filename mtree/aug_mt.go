package mtree

import (
	"container/list"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Counter interface{}

type AugData struct {
	Min  Counter
	Max  Counter
	Hash smartpool.SPHash
}

func (ad AugData) Copy() NodeData {
	h := smartpool.SPHash{}
	min := big.NewInt(0)
	max := big.NewInt(0)
	copy(h[:], ad.Hash[:])
	return AugData{
		min.Add(min, ad.Min.(*big.Int)),
		max.Add(max, ad.Max.(*big.Int)),
		h,
	}
}

func (ad AugData) CounterBytes() []byte {
	max := msbPadding(ad.Max.(*big.Int).Bytes(), 16)
	min := msbPadding(ad.Min.(*big.Int).Bytes(), 16)
	return append(max, min...)
}

type AugTree struct {
	MerkleTree
}

func _min(a, b Counter) Counter {
	left := a.(*big.Int)
	right := b.(*big.Int)
	if left.Cmp(right) == -1 {
		return left
	} else {
		return right
	}
}

func _max(a, b Counter) Counter {
	left := a.(*big.Int)
	right := b.(*big.Int)
	if left.Cmp(right) == 1 {
		return left
	} else {
		return right
	}
}

func _augModifier(data NodeData) {
	dummy := data.(AugData)
	max := dummy.Max.(*big.Int)
	min := dummy.Min.(*big.Int)
	min.Add(max, big.NewInt(1))
	max.Add(max, big.NewInt(2))
}

func _augElementHash(data ElementData) NodeData {
	s := data.(smartpool.Share)
	// fmt.Printf("Constructing node:\n")
	// fmt.Printf("	Min: %v\n", s.Counter())
	// fmt.Printf("	Max: %v\n", s.Counter())
	// fmt.Printf("	Hash: %s\n", s.Hash().Hex())
	return AugData{
		Min:  s.Counter(),
		Max:  s.Counter(),
		Hash: s.Hash(),
	}
}

func _augHash(a, b NodeData) NodeData {
	left := a.(AugData)
	right := b.(AugData)
	h := smartpool.SPHash{}
	keccak := crypto.Keccak256(
		left.CounterBytes(),
		append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, left.Hash[:]...),
		right.CounterBytes(),
		append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, right.Hash[:]...),
	)
	copy(h[:smartpool.HashLength], keccak[smartpool.HashLength:])
	return AugData{
		Min:  _min(left.Min, right.Min),
		Max:  _max(left.Max, right.Max),
		Hash: h,
	}
}

func NewAugTree() *AugTree {
	mtbuf := list.New()
	return &AugTree{
		MerkleTree{
			mtbuf,
			_augHash,
			_augElementHash,
			_augModifier,
			false,
			map[uint32]bool{},
			[]uint32{},
			0,
			0,
			[]NodeData{},
		},
	}
}

func (amt AugTree) RootHash() smartpool.SPHash {
	return amt.Root().(AugData).Hash
}

func (amt AugTree) RootMin() *big.Int {
	return amt.Root().(AugData).Min.(*big.Int)
}

func (amt AugTree) RootMax() *big.Int {
	return amt.Root().(AugData).Max.(*big.Int)
}

func (amt AugTree) CounterBranchArray() []*big.Int {
	if amt.finalized {
		result := []*big.Int{}
		branches := amt.Branches()
		var node AugData
		for _, k := range amt.Indices() {
			nodes := branches[k].ToNodeArray()[1:]
			for _, n := range nodes {
				node = n.(AugData)
				be := smartpool.BranchElement{}
				copy(be[:], node.CounterBytes())
				result = append(result, be.Big())
			}
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

func (amt AugTree) HashBranchArray() []*big.Int {
	if amt.finalized {
		result := []*big.Int{}
		branches := amt.Branches()
		var node AugData
		for _, k := range amt.Indices() {
			nodes := branches[k].ToNodeArray()[1:]
			for _, n := range nodes {
				node = n.(AugData)
				be := smartpool.BranchElement{}
				copy(be[:], msbPadding(node.Hash[:], 32))
				result = append(result, be.Big())
			}
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}
