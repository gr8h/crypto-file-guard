package merkletree

import (
	"crypto/sha256"
	"errors"
)

func HashData(data Data) (Hash, error) {

	if data == nil {
		return nil, errors.New("merkletree: data is empty")
	}

	sum := sha256.Sum256(data)
	return sum[:], nil
}
