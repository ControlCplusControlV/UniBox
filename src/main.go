package main

import (
	"github.com/rivo/tview"
)

func main() {
	userConfig := loadConfig()

	if !userConfig["initialized"].(bool) {
		setTrackedTokens()
		setTrackedPools()
	}

	var ANSCIIUniCorn string = "   UniBox         \n      \\                \n       _\\,,            \n      \"-=\\~     _  \n         \\~___( ~\n        _|/---\\_   \n        \\        \\     \n Credit (ejm97)"

	// Uncomment to enable automatic strategies
	//go strategy.MainTick()
	app := tview.NewApplication()
	modal := tview.NewModal().
		SetText(ANSCIIUniCorn).
		AddButtons([]string{"Enter Swap Interface", "Enter Analytics Interface"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Enter Swap Interface" {
				swapTerminal(userConfig)
			} else if buttonLabel == "Enter Analytics Interface" {
				drawTerminal(userConfig)
			}
		})
	modal.SetBorder(true).SetTitle("UniBox v0.1").SetTitleAlign(tview.AlignCenter)

	if err := app.SetRoot(modal, false).SetFocus(modal).Run(); err != nil {
		panic(err)
	}
}
