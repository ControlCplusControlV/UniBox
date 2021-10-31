package main

import (
	"fmt"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/guptarohit/asciigraph"
	"github.com/rivo/tview"
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

func promptUser() {
	answers := Checkboxes(
		"Which tokens do you want to track?",
		[]string{
			"USDC",
			"USDT",
			"DAI",
			"UNI",
			"WBTC",
			"SHIB",
			"LINK",
		},
	)
	s := strings.Join(answers, ", ")
	fmt.Println("Oh, I see! You like", s)
}

func graphVolume() string {
	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 6}
	graph := asciigraph.Plot(data)

	graph = "A Cool Title \n" + graph

	return graph
}

func graphTVL() string {
	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 6}
	graph := asciigraph.Plot(data)

	graph = "A Cool Title \n" + graph

	return graph
}

func graphThePool() string {
	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 6}
	graph := asciigraph.Plot(data)

	graph = "A Cool Title \n" + graph

	return graph
}

type Token struct {
	upOrDown      bool
	percentChange string
	name          string
}

func getTrackedTokens(arrayOfTokens []Token) string {
	var outputString string = "Tracked Tokens"
	for index := 0; index < len(arrayOfTokens); index++ {
		var currentToken Token = arrayOfTokens[index]
		if currentToken.upOrDown {
			thisTokenString := currentToken.name + " " + "+" + currentToken.percentChange
			outputString = outputString + "\n" + thisTokenString

		} else {
			thisTokenString := currentToken.name + " " + "-" + currentToken.percentChange
			outputString = outputString + "\n" + thisTokenString

		}
	}

	return outputString
}

func main() {
	drawTerminal()
}
func drawTerminal() {
	// Main function to draw out terminal windows
	type settings struct {
		tokens       []string
		trackedPools []string
	}

	mainTerminal := tview.NewApplication()

	var trackedTokensList []Token = []Token{{true, "15", "USDC"}, {false, "10", "BTC"}, {true, "150", "SHIBA"}}
	//trackedTokens := newPrimitive("Tracked Tokens")
	//ProtocolTVL := newPrimitive("Uniswap TVL")

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
		SetColumns(30, 0, 30).
		SetBorders(true)

	// Layout for screens wider than 100 cells.
	grid.AddItem(trackedTokens, 0, 0, 4, 1, 1, 100, false).
		//AddItem(ProtocolTVL, 0, 1, 2, 1, 0, 100, false).
		AddItem(graphOfTVL, 0, 1, 2, 1, 0, 100, false).
		AddItem(graphOfVolume, 2, 1, 2, 1, 0, 100, false).
		//User tracked stats
		AddItem(trackedStat1, 0, 2, 2, 1, 0, 100, false).
		AddItem(trackedStat2, 2, 2, 2, 1, 0, 100, false)
	mainTerminal.SetRoot(grid, true).SetFocus(grid).Run()
}
