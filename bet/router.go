package bet

import (
	"BlockchainBetSportGolang/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{
	blockchain: &Blockchain{
		Chain:        Blocks{},
		PendingBets:  Bets{},
		NetworkNodes: []string{}}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"GetBlockchain",
		"GET",
		"/blockchain",
		controller.GetBlockchain,
	},
	Route{
		"RegisterAndBroadcastNode",
		"POST",
		"/register-and-broadcast-node",
		controller.RegisterAndBroadcastNode,
	},
	Route{
		"RegisterNode",
		"POST",
		"/register-node",
		controller.RegisterNode,
	},
	Route{
		"RegisterNodesBulk",
		"POST",
		"/register-nodes-bulk",
		controller.RegisterNodesBulk,
	},
	Route{
		"RegisterBet",
		"POST",
		"/bet",
		controller.RegisterBet,
	},
	Route{
		"RegisterAndBroadcastBet",
		"POST",
		"/bet/broadcast",
		controller.RegisterAndBroadcastBet,
	},
	Route{
		"Mine",
		"GET",
		"/mine",
		controller.Mine,
	},
	Route{
		"ReceiveNewBlock",
		"POST",
		"/receive-new-block",
		controller.ReceiveNewBlock,
	},
	Route{
		"Consensus",
		"GET",
		"/consensus",
		controller.Consensus,
	},
	Route{
		"GetBetsForMatch",
		"GET",
		"/match/{matchId}",
		controller.GetBetsForMatch,
	},
	Route{
		"GetBetsForPlayer",
		"GET",
		"/player/{playerName}",
		controller.GetBetsForPlayer,
	},
}

//NewRouter configures a new router to the API
func NewRouter(nodeAddress string) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	controller.currentNodeURL = "http://localhost:" + nodeAddress

	// create Genesis block
	controller.blockchain.CreateNewBlock(100, "0", "0")

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
