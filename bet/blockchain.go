package bet

import (
	"crypto/sha256"
	"strconv"
	"time"
)

//RegisterBet registers a bet in our blockchain
func (b *Blockchain) RegisterBet(bet Bet) bool {
	b.PendingBets = append(b.PendingBets, bet)
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//RegisterNode registers a node in our blockchain
func (b *Blockchain) RegisterNode(node string) bool {
	if !contains(b.NetworkNodes, node) {
		b.NetworkNodes = append(b.NetworkNodes, node)
	}
	return true
}

//CreateNewBlock ...
func (b *Blockchain) CreateNewBlock(nonce int, previousBlockHash string, hash string) Block {
	newBlock := Block{
		Index:     len(b.Chain) + 1,
		Bets:      b.PendingBets,
		Timestamp: time.Now(),
		Nonce:     nonce,
		Hash:      hash, PreviousBlockHash: previousBlockHash}

	b.PendingBets = Bets{}
	b.Chain = append(b.Chain, newBlock)
	return newBlock
}

//GetLastBlock ...
func (b *Blockchain) GetLastBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

//HashBlock ...
func (b *Blockchain) HashBlock(previousBlockHash string, currentBlockData string, nonce int) string {
	h := sha256.New()
	strToHash := previousBlockHash + currentBlockData + strconv.Itoa(nonce)
	h.Write([]byte(strToHash))
	return string(h.Sum(nil)[:])
}

//ProofOfWork ...
func (b *Blockchain) ProofOfWork(previousBlockHash string, currentBlockData string) int {
	nonce := 0
	hash := b.HashBlock(previousBlockHash, currentBlockData, nonce)
	inputFmt := hash[0:4]
	for inputFmt != "0000" {
		nonce = nonce + 1
		hash = b.HashBlock(previousBlockHash, currentBlockData, nonce)
	}
	return nonce
}
