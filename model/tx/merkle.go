package tx

import (
	"github.com/karriz-dev/symbol-sdk/errors"

	"golang.org/x/crypto/sha3"
)

func MerkleRootHash(txs []Transaction) (Hash, error) {
	if len(txs) <= 0 {
		return Hash{}, errors.ErrEmptyTransaction
	}

	var hashes []Hash
	hasher := sha3.New256()

	for _, innerTx := range txs {
		innerTxSerializedBytes, err := innerTx.Serialize()
		if err != nil {
			return Hash{}, err
		}

		hasher.Reset()
		hasher.Write(innerTxSerializedBytes)

		hashes = append(hashes, Hash(hasher.Sum(nil)[:]))
	}

	if len(hashes) == 1 {
		return Hash(hashes[0]), nil
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

			hashes[i/2] = Hash(hasher.Sum(nil))
			i += 2
		}

		numRemainingHashes = numRemainingHashes / 2
	}

	return Hash(hashes[0]), nil
}
