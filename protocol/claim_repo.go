package protocol

import (
	"github.com/SmartPool/smartpool-client"
)

// ClaimRepo holds many claims but only 1 active clam at a time which is
// storing coming shares.
type ClaimRepo interface {
	// AddShare stores share to current active claim
	AddShare(s smartpool.Share) error
	// GetCurrentClaim returns the active claim, seal it as closed claim and
	// initialize a new active claim to store coming shares. In the mean time
	// the sealed claim should be used by SmartPool to do the protocol.
	// GetCurrentClaim returns nil when there are not enough shares in the claim
	// as compared to the threshold.
	GetCurrentClaim(threshold int) smartpool.Claim
	Persist(storage smartpool.PersistentStorage) error
	NoActiveShares() uint64
}
