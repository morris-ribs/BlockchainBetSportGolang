package bet

//Bet represents a bet
type Bet struct {
	PlayerName   string `json:"player_name"`
	MatchID      string `json:"match_id"`
	TeamOneScore int32  `json:"team_one_score"`
	TeamTwoScore int32  `json:"team_two_score"`
}

//Bets is an array of Bet
type Bets []Bet

//Block ...
type Block struct {
	Index             int    `json:"index"`
	Timestamp         int64  `json:"timestamp"`
	Bets              Bets   `json:"bets"`
	Nonce             int    `json:"nonce"`
	Hash              string `json:"hash"`
	PreviousBlockHash string `json:"previous_block_hash"`
}

//Blocks is an array of Block
type Blocks []Block

//Blockchain ...
type Blockchain struct {
	Chain        Blocks   `json:"chain"`
	PendingBets  Bets     `json:"pending_bets"`
	NetworkNodes []string `json:"network_nodes"`
}

//BlockData is used in hash calculations
type BlockData struct {
	Index string
	Bets  Bets
}
