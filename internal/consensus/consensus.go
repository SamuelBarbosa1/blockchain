package consensus

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func ProofOfWork(blockHash string, difficulty int) string {
	nonce := 0
	var hash string
	for {
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s%d", blockHash, nonce))))
		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			break
		}
		nonce++
	}
	return hash
}
