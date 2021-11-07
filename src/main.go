package main

import (
	"github.com/rivo/tview"
	"main.go/swap"
)

func main() {
	userConfig := loadConfig()

	if !userConfig["initialized"].(bool) {
		setTrackedTokens()
		setTrackedPools()
	}

	var tokensToApprove []interface{} = userConfig["needsApprove"].([]interface{})
	for index := 0; index < len(tokensToApprove); index++ {
		swap.Approve(tokensToApprove[index].(string))
	}

	//go strategy.MainTick()
	app := tview.NewApplication()
	modal := tview.NewModal().
		SetText("Do you want to quit the application? \n Can I do new lines?").
		AddButtons([]string{"Enter Swap Interface", "Cancel"}).
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
