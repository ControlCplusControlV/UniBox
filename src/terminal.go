package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/guptarohit/asciigraph"
	"github.com/rivo/tview"
	"main.go/dataAggregator"
	"main.go/price"
	"main.go/swap"
)

func Checkboxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func loadConfig() map[string]interface{} {

	var configMap = make(map[string]interface{})

	var avaliableTokensMap = make(map[string]interface{})

	var trackedTokensMap = make(map[string]interface{})

	//hard code available tokens
	avaliableTokensMap["USDC"] = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	//usdt ropsten address
	avaliableTokensMap["USDT"] = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	avaliableTokensMap["DAI"] = "0x6b175474e89094c44da98b954eedeac495271d0f"
	avaliableTokensMap["UNI"] = "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
	avaliableTokensMap["WBTC"] = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
	avaliableTokensMap["SHIB"] = "0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce"
	avaliableTokensMap["LINK"] = "0x514910771af9ca656af840dff83e8264ecf986ca"
	//weth ropsten address
	avaliableTokensMap["WETH"] = "0xc778417e063141139fce010982780140aa0cd5ab"

	//hard code trackded tokens
	trackedTokensMap["DAI"] = false
	trackedTokensMap["LINK"] = false
	trackedTokensMap["SHIB"] = true
	trackedTokensMap["UNI"] = true
	trackedTokensMap["USDC"] = true
	trackedTokensMap["USDT"] = false
	trackedTokensMap["WBTC"] = false

	//hard code config map
	configMap["availableTokens"] = avaliableTokensMap
	configMap["trackedTokens"] = trackedTokensMap
	configMap["initalized"] = true
	configMap["availableTokens"] = []string{"USDC", "USDT", "DAI", "UNI", "WBTC", "SHIB", "LINK", "WETH"}

	return configMap

}

func setTrackedTokens() {

	var avaliableTokens = []string{
		"USDC",
		"USDT",
		"DAI",
		"UNI",
		"WBTC",
		"SHIB",
		"LINK",
	}

	prompt := &survey.MultiSelect{
		Message: "Which tokens do you want to track?",
		Options: avaliableTokens,
	}

	answers := []string{}

	survey.AskOne(prompt, &answers)

	var trackedTokens = make(map[string]bool)

	for index := 0; index < len(avaliableTokens); index++ {
		trackedTokens[avaliableTokens[index]] = false
	}

	for index := 0; index < len(answers); index++ {
		trackedTokens[answers[index]] = true
	}

	jsonString, err := json.Marshal(trackedTokens)

	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("config.json", jsonString, 0644)

	if err != nil {
		fmt.Println(err)
	}

}

func setTrackedPools() {

	var avaliablePools = []string{
		"USDC / UNI",
		"USDT / WBTC",
		"DAI / USDT",
		"UNI / WETJ",
		"WBTC / DAI",
		"SHIB / USDC",
	}

	prompt := &survey.MultiSelect{
		Message: "Which pools do you want to track?",
		Options: avaliablePools,
	}

	answers := []string{}

	survey.AskOne(prompt, &answers)

	var trackedPools = make(map[string]bool)

	for index := 0; index < len(avaliablePools); index++ {
		trackedPools[avaliablePools[index]] = false
	}

	for index := 0; index < len(answers); index++ {
		trackedPools[answers[index]] = true
	}

	jsonString, err := json.Marshal(trackedPools)

	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("config.json", jsonString, 0644)

	if err != nil {
		fmt.Println(err)
	}

}

func stringToFloat(stringArray []string) []float64 {
	var returnArray []float64

	for index := 0; index < len(stringArray); index++ {
		result, err := strconv.ParseFloat(stringArray[index], 64)

		if err != nil {
			fmt.Println("There is an error converting string to float.")
		}

		returnArray = append(returnArray, result)
	}

	return returnArray
}

func graphVolume() string {
	data := dataAggregator.ProtocolHistoricalStats().(map[string]interface{})["dailyVolumeETH"].([]string)

	dataFloat := stringToFloat(data)

	graph := asciigraph.Plot(dataFloat, asciigraph.Precision(10), asciigraph.Width(30), asciigraph.Height(10))

	graph = "Protocol Volume by Day in ETH \n" + graph

	return graph
}

func graphTVL() string {
	data := dataAggregator.ProtocolHistoricalStats().(map[string]interface{})["totalLiquidityETH"].([]string)

	dataFloat := stringToFloat(data)

	graph := asciigraph.Plot(dataFloat, asciigraph.Precision(10), asciigraph.Width(30), asciigraph.Height(10))

	graph = "Protocol TVL by Day in ETH \n" + graph

	return graph
}

func graphThePool() string {
	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 6}
	graph := asciigraph.Plot(data, asciigraph.Precision(10), asciigraph.Width(20), asciigraph.Height(10))

	graph = "A Cool Title \n" + graph

	return graph
}

type Token struct {
	TokenValue string
	name       string
}

func getTrackedTokens(arrayOfTokens []Token) string {
	var outputString string = "Tracked Tokens"
	for index := 0; index < len(arrayOfTokens); index++ {
		var currentToken Token = arrayOfTokens[index]
		thisTokenString := currentToken.name + " " + "+" + currentToken.TokenValue
		outputString = outputString + "\n" + thisTokenString

	}

	return outputString
}

func drawTerminal(config map[string]interface{}) {
	// Main function to draw out terminal windows
	type settings struct {
		tokens       []string
		trackedPools []string
	}

	mainTerminal := tview.NewApplication()

	var trackedTokensList []Token

	tokenList := config["trackedTokens"].(map[string]interface{})

	var avaliableTokens = config["avaliableTokens"].([]interface{})

	var ethereumPrice = price.GetEthereumPriceInUSD()

	for index := 0; index < len(avaliableTokens); index++ {
		if tokenList[avaliableTokens[index].(string)].(bool) {

			var tokenAddress string = config[avaliableTokens[index].(string)].(string)

			var token interface{} = dataAggregator.GetAllPairStatsForToken(tokenAddress)

			var priceInETH string = token.(map[string]interface{})["derivedETH"].(string)

			priceInETHFloat, err := strconv.ParseFloat(priceInETH, 64)
			if err != nil {
				fmt.Println(err)
			}
			var tokenPriceFloat = ethereumPrice * priceInETHFloat
			s := fmt.Sprintf("%f.2", tokenPriceFloat)

			trackedTokensList = append(trackedTokensList, Token{s, avaliableTokens[index].(string)})
		}
	}

	// Next draw out a graph for everything
	trackedTokens := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			mainTerminal.Draw()
		})
	fmt.Fprintf(trackedTokens, "%s ", (getTrackedTokens(trackedTokensList)))
	graphOfTVL := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			mainTerminal.Draw()
		})
	fmt.Fprintf(graphOfTVL, "%s ", graphTVL())
	graphOfVolume := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			mainTerminal.Draw()
		})
	fmt.Fprintf(graphOfVolume, "%s ", graphVolume())

	trackedStat1 := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(false).
		SetChangedFunc(func() {
			mainTerminal.Draw()
		})
	fmt.Fprintf(trackedStat1, "%s ", graphThePool())

	trackedStat2 := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(false).
		SetChangedFunc(func() {
			mainTerminal.Draw()
		})
	fmt.Fprintf(trackedStat2, "%s ", graphThePool())

	trackedStat3 := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(false).
		SetChangedFunc(func() {
			mainTerminal.Draw()
		})
	fmt.Fprintf(trackedStat3, "%s ", graphThePool())

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(15, 0, 0).
		SetBorders(true)

	// Layout for screens wider than 100 cells.
	grid.AddItem(trackedTokens, 0, 0, 4, 1, 1, 100, false).
		//AddItem(ProtocolTVL, 0, 1, 2, 1, 0, 100, false).
		AddItem(graphOfTVL, 0, 1, 2, 1, 0, 100, false).
		AddItem(graphOfVolume, 2, 1, 2, 1, 0, 100, false).
		//User tracked stats
		AddItem(trackedStat1, 0, 2, 2, 1, 20, 100, false).
		AddItem(trackedStat2, 2, 2, 2, 1, 20, 100, false)

	//(p , row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool)
	mainTerminal.SetRoot(grid, true).SetFocus(grid).Run()
}

func swapTerminal(config map[string]interface{}) {
	var avaliableTokens = config["avaliableTokens"].([]interface{})

	avaliableTokensString := make([]string, len(avaliableTokens))
	for i, v := range avaliableTokens {
		avaliableTokensString[i] = fmt.Sprint(v)
	}

	var inputToken string

	var outputToken string

	var inputAmount int64

	app := tview.NewApplication()
	form := tview.NewForm().
		AddDropDown("Input Token", avaliableTokensString, 0, func(option string, optionIndex int) {
			inputToken = option
		}).
		AddDropDown("Output Token", avaliableTokensString, 0, func(option string, optionIndex int) {
			outputToken = option
		}).
		AddInputField("Input Token Amount", "", 20, nil, func(textToCheck string) {
			inputTokens, err := strconv.ParseInt(textToCheck, 10, 0)
			if err != nil {
				panic(err) // Error handling done right
			}
			inputAmount = inputTokens
		}).
		AddButton("Execute Trade", func() {
			swapTokens(inputToken, outputToken, inputAmount)
		})
	form.SetBorder(true).SetTitle("Swap Tokens (Tab to Move On)").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}

func swapTokens(tokenInAddress string, tokenOutAddress string, amountIn int64) {
	swap.SwapExactTokensForTokens(uint(amountIn), 0, []common.Address{common.HexToAddress(tokenInAddress), common.HexToAddress(tokenOutAddress)}, common.HexToAddress(swap.WalletAddress), 10000000000)

}
