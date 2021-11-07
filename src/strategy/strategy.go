package strategy

import "main.go/dataAggregator"

func MainTick() {
	for {
		//Put your strategy functions in here
		buyLow()
	}
}

//Example Strategy
func buyLow() bool {
	dataAggregator.CurrentTokenStatsAllPairs("0x110a13fc3efe6a245b50102d2d79b3e76125ae83")
	return true

}
