package dataAggregator

import (
	"fmt"

	"github.com/senseyeio/roger"
)

func evaluateRCmd(rCmd string) interface{} {
	rClient, err := roger.NewRClient("127.0.0.1", 6311)
	if err != nil {
		fmt.Println("Failed to connect")
		return nil
	}
	sesh, _ := rClient.GetSession()

	result, err := sesh.Eval(rCmd)
	if err != nil {
		fmt.Println(err)
		sesh.Close()
		return nil
	} else {
		sesh.Close()
		return result
	}
}

func GetFactoryStats() interface{} {
	return evaluateRCmd("factory_stats_v2()")
}

func ProtocolHistoricalStats() interface{} {
	return evaluateRCmd("uniswap_stats_hist_v2()")
}

func CurrentTokenStatsAllPairs(token string) interface{} {
	var rCmd string = "token_stats_v2(token_address = \"" + token + "\")"
	return evaluateRCmd(rCmd)

}

func HistoricalTokenStatsAllPairs(token string) interface{} {
	var rCmd string = "token_stats_hist_v2(token_address = \"" + token + "\")"
	return evaluateRCmd(rCmd)

}

func GetAllPairStatsForToken(token string) interface{} {
	var rCmd string = "token_pair_map_v2(token_address = \"" + token + "\")"
	return evaluateRCmd(rCmd)
}

func ViewStatsV2(timeframe uint8, pair string) interface{} {
	// I probably need more documentation for this method
	if timeframe == 0 {
		var rCmd string = "pair_stats_v2(pair_address  = \"" + pair + "\")"
		return evaluateRCmd(rCmd)
	} else if timeframe == 1 {
		var rCmd string = "pair_stats_hist_hourly_v2(pair_address  = \"" + pair + "\")"
		return evaluateRCmd(rCmd)
	} else if timeframe == 2 {
		var rCmd string = "pair_stats_hist_daily_v2(pair_address  = \"" + pair + "\")"
		return evaluateRCmd(rCmd)
	}
	return nil
}

func GetLiquidityPositions(historical_or_not bool, pair string) interface{} {
	if historical_or_not {
		var rCmd string = "pair_liq_positions_hist_v2(pair_address  = \"" + pair + "\")"
		return evaluateRCmd(rCmd)
	} else {
		var rCmd string = "pair_liq_positions_v2(pair_address  = \"" + pair + "\")"
		return evaluateRCmd(rCmd)
	}
}

func GetUseLiquidityPositions(user string) interface{} {
	var rCmd string = "user_lps_v2(user_address   = \"" + user + "\")"
	return evaluateRCmd(rCmd)
}

func GetHistoricalLiquidityPositions(user string) interface{} {
	var rCmd string = "user_hist_lps_v2(user_address   = \"" + user + "\")"
	return evaluateRCmd(rCmd)
}
