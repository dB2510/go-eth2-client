// Code generated by fastssz. DO NOT EDIT.
// Hash: 11fa0877ac77d8a12ad1e550521cb5ecb386af10f1e7f8183b385bf37f95b972
package capella

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Withdrawal object
func (w *Withdrawal) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(w)
}

// MarshalSSZTo ssz marshals the Withdrawal object to a target array
func (w *Withdrawal) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Index'
	dst = ssz.MarshalUint64(dst, uint64(w.Index))

	// Field (1) 'Address'
	dst = append(dst, w.Address[:]...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, uint64(w.Amount))

	return
}

// UnmarshalSSZ ssz unmarshals the Withdrawal object
func (w *Withdrawal) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 36 {
		return ssz.ErrSize
	}

	// Field (0) 'Index'
	w.Index = WithdrawalIndex(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'Address'
	copy(w.Address[:], buf[8:28])

	// Field (2) 'Amount'
	w.Amount = phase0.Gwei(ssz.UnmarshallUint64(buf[28:36]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Withdrawal object
func (w *Withdrawal) SizeSSZ() (size int) {
	size = 36
	return
}

// HashTreeRoot ssz hashes the Withdrawal object
func (w *Withdrawal) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(w)
}

// HashTreeRootWith ssz hashes the Withdrawal object with a hasher
func (w *Withdrawal) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Index'
	hh.PutUint64(uint64(w.Index))

	// Field (1) 'Address'
	hh.PutBytes(w.Address[:])

	// Field (2) 'Amount'
	hh.PutUint64(uint64(w.Amount))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Withdrawal object
func (w *Withdrawal) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(w)
}
