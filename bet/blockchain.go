package bet

import "time"

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
func (b *Blockchain) CreateNewBlock(nonce int32, previousBlockHash string, hash string) Block {
	newBlock := Block{
		Bets:      b.PendingBets,
		Timestamp: time.Now(),
		Nonce:     nonce,
		Hash:      hash, PreviousBlockHash: previousBlockHash}

	b.PendingBets = Bets{}
	b.Chain = append(b.Chain, newBlock)
	return newBlock
}

//GetLastBlock ...
