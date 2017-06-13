package ethash

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var (
	Instance   = New(defaultDir(), 3, 3, defaultDir(), 0, 3)
	DefaultDir = defaultDir()
)

func defaultDir() string {
	home := os.Getenv("HOME")
	if user, err := user.Current(); err == nil {
		home = user.HomeDir
	}
	if runtime.GOOS == "windows" {
		return filepath.Join(home, "AppData", "Ethash")
	}
	return filepath.Join(home, ".ethash")
}

func DAGSize(blockNum uint64) uint64 {
	return datasetSizes[blockNum/epochLength]
}

func (ethash *Ethash) GetVerificationIndices(blockNumber uint64, hash common.Hash, nonce uint64) []uint32 {
	// Recompute the digest and PoW value and verify against the header
	cache := ethash.cache(blockNumber)

	size := datasetSize(blockNumber)
	return hashimotoLightIndices(size, cache, hash.Bytes(), nonce)
}

func hashimotoLightIndices(size uint64, cache []uint32, hash []byte, nonce uint64) []uint32 {
	keccak512 := makeHasher(sha3.NewKeccak512())

	lookup := func(index uint32) []uint32 {
		rawData := generateDatasetItem(cache, index, keccak512)

		data := make([]uint32, len(rawData)/4)
		for i := 0; i < len(data); i++ {
			data[i] = binary.LittleEndian.Uint32(rawData[i*4:])
		}
		return data
	}
	return hashimotoIndices(hash, nonce, size, lookup)
}

func hashimotoIndices(hash []byte, nonce uint64, size uint64, lookup func(index uint32) []uint32) []uint32 {
	result := []uint32{}
	// Calculate the number of thoretical rows (we use one buffer nonetheless)
	rows := uint32(size / mixBytes)

	// Combine header+nonce into a 64 byte seed
	seed := make([]byte, 40)
	copy(seed, hash)
	binary.LittleEndian.PutUint64(seed[32:], nonce)

	seed = crypto.Keccak512(seed)
	seedHead := binary.LittleEndian.Uint32(seed)

	// Start the mix with replicated seed
	mix := make([]uint32, mixBytes/4)
	for i := 0; i < len(mix); i++ {
		mix[i] = binary.LittleEndian.Uint32(seed[i%16*4:])
	}
	// Mix in random dataset nodes
	temp := make([]uint32, len(mix))

	for i := 0; i < loopAccesses; i++ {
		parent := fnv(uint32(i)^seedHead, mix[i%len(mix)]) % rows
		result = append(result, parent)
		for j := uint32(0); j < mixBytes/hashBytes; j++ {
			copy(temp[j*hashWords:], lookup(2*parent+j))
		}
		fnvHash(mix, temp)
	}
	return result
}

func MakeDAG(block uint64, dir string) {
	MakeDataset(block-30000, dir)
}

func PathToDAG(epoch uint64, dir string) string {
	seed := ethash.SeedHash(epoch*epochLength + 1)
	var endian string
	if !isLittleEndian() {
		endian = ".be"
	}
	return filepath.Join(dir, fmt.Sprintf("full-R%d-%x%s", 23, seed[:8], endian))
}
