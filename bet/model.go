package bet

import "time"

//Bet represents a bet
type Bet struct {
	PlayerName   string `json:"playerName"`
	MatchID      string `json:"matchId"`
	TeamOneScore int32  `json:"teamOneScore"`
	TeamTwoScore int32  `json:"teamTwoScore"`
}

//Bets is an array of Bet
type Bets []Bet

//Block ...
type Block struct {
	Index             int
	Timestamp         time.Time
	Bets              Bets
	Nonce             int
	Hash              string
	PreviousBlockHash string
}

//Blocks is an array of Block
type Blocks []Block

//Blockchain ...
type Blockchain struct {
	Chain        Blocks
	PendingBets  Bets
	NetworkNodes []string
}
