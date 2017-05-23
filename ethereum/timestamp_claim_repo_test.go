package ethereum

import (
	"encoding/json"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func newRepo() *TimestampClaimRepo {
	return NewTimestampClaimRepo(
		big.NewInt(100000),
		"0xe034afdcc2ba0441ff215ee9ba0da3e86450108d",
		"0x9af93376af1ddd22fa2e94fd0a030b3dea96bb96",
		&testPersistentStorage{},
	)
}

func newTestShare() *Share {
	jsonStr := `{"header":{"parentHash":"0xf1a495416b302d37af3072be729d971f06f890ab907f9267daa1ff8b3daecb09","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","miner":"0x9af93376af1ddd22fa2e94fd0a030b3dea96bb96","stateRoot":"0x28cc5f7a6852d59ad3c0de622554eada30835a856d08fd6d2f73169a299e3a3f","transactionsRoot":"0x1cdadce60b3b3deeadf556aa7d790e40efdb69f5dd8542d6b73fc92a45806606","receiptsRoot":"0x536f8e55d08a19a09b8454b534cba193ed33158e37e5c07c2235f285c02d9338","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","difficulty":"0x8aae19b2","number":"0xef5c3","gasLimit":"0x47b784","gasUsed":"0xa410","timestamp":"0x59230730","extraData":"0x536d617274506f6f6c2d4e736a645a465776556f6e5530713030303030303030","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000"},"nonce":"0x59cead693238d7b4","mix":"0x31a24b0ccd070af5f901e798d6d760b46ecfd8e1ee04ffe46c12dc41a9a5551b","share_diff":100000,"miner":"0xe034afdcc2ba0441ff215ee9ba0da3e86450108d","state":1}`
	s := gobShare{}
	json.Unmarshal([]byte(jsonStr), &s)
	return &Share{
		s.BlockHeader,
		s.Nonce,
		s.MixDigest,
		s.ShareDifficulty,
		s.MinerAddress,
		s.SolutionState,
		nil,
	}
}

func endlessAddShare(repo *TimestampClaimRepo, stop chan bool) {
	var counter int64 = 0
	for {
		select {
		case <-stop:
			return
		default:
			counter++
			s := newTestShare()
			s.nonce = types.EncodeNonce(uint64(rand.Int63()))
			s.blockHeader.Time.Add(s.blockHeader.Time, big.NewInt(counter))
			repo.AddShare(s)
		}
	}
}

func endlessPersist(repo *TimestampClaimRepo, storage smartpool.PersistentStorage, stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			repo.Persist(storage)
		}
	}
}

func endlessGetCurrentClaim(t *testing.T, repo *TimestampClaimRepo, stop chan bool) {
	var claim smartpool.Claim
	for {
		select {
		case <-stop:
			if claim == nil {
				t.Fail()
			}
			return
		default:
			claim = repo.GetCurrentClaim(2)
			if claim != nil {
				return
			}
		}
	}
}

func TestPersistAndAddShareAndGetCurrentClaimConcurrently(t *testing.T) {
	repo := newRepo()
	storage := &testPersistentStorage{}
	stopAddingShare := make(chan bool, 1)
	stopPersisting := make(chan bool, 1)
	stopGettingClaim := make(chan bool, 1)
	go endlessAddShare(repo, stopAddingShare)
	go endlessPersist(repo, storage, stopPersisting)
	go endlessGetCurrentClaim(t, repo, stopGettingClaim)
	time.Sleep(50 * time.Millisecond)
	stopAddingShare <- true
	stopPersisting <- true
	stopGettingClaim <- true
	time.Sleep(50 * time.Millisecond)
}
