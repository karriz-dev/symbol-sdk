package merkle

import (
	"errors"

	"github.com/karriz-dev/symbol-sdk/model/tx"
	"golang.org/x/crypto/sha3"
)

var (
	ErrEmptyTransaction = errors.New("transaction list size cannot be zero")
)

func MerkleRootHash(txs []tx.Transaction) (tx.Hash, error) {
	if len(txs) <= 0 {
		return tx.Hash{}, ErrEmptyTransaction
	}

	var hashes []tx.Hash
	hasher := sha3.New256()

	for _, tx := range txs {
		hasher.Reset()
		hasher.Write(tx.Serialize())

		hashes = append(hashes, tx.Hash(hasher.Sum(nil)[:]))
	}

	if len(hashes) == 1 {
		return tx.Hash(hashes[0]), nil
	}

	numRemainingHashes := len(hashes)

	for 1 < numRemainingHashes {
		i := 0
		for i < numRemainingHashes {
			hasher.Reset()
			hasher.Write(hashes[i][:])

			if i+1 < numRemainingHashes {
				hasher.Write(hashes[i+1][:])
			} else {
				hasher.Write(hashes[i][:])
				numRemainingHashes += 1
			}

			hashes[i/2] = tx.Hash(hasher.Sum(nil))
			i += 2
		}

		numRemainingHashes = numRemainingHashes / 2
	}

	return tx.Hash(hashes[0]), nil
}
