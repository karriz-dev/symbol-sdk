package tx

import (
	"errors"

	"github.com/karriz-dev/symbol-sdk/common"
	"golang.org/x/crypto/sha3"
)

var (
	ErrEmptyTransaction = errors.New("transaction list size cannot be zero")
)

func MerkleRootHash(transactions []ITransaction) (common.Hash, error) {
	if len(transactions) <= 0 {
		return common.Hash{}, ErrEmptyTransaction
	}

	var hashes []common.Hash
	hasher := sha3.New256()

	for _, tx := range transactions {
		hasher.Reset()

		serializeBytes, err := tx.Serialize()
		if err != nil {
			return common.Hash{}, err
		}

		hasher.Write(serializeBytes)
		hashes = append(hashes, common.Hash(hasher.Sum(nil)[:]))
	}

	if len(hashes) == 1 {
		return common.Hash(hashes[0]), nil
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

			hashes[i/2] = common.Hash(hasher.Sum(nil))
			i += 2
		}

		numRemainingHashes = numRemainingHashes / 2
	}

	return common.Hash(hashes[0]), nil
}
