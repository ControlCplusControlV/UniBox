package main

import (
	"github.com/rivo/tview"
	"main.go/swap"
	"main.go/strategy"
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

	go strategy.MainTick()

	app := tview.NewApplication()
	form := tview.NewForm().
		AddButton("Enter Swap Interface", func() {
			swapTerminal(userConfig)
		}).
		AddButton("Enter Analytics Interface", func() {
			drawTerminal(userConfig)
		})
	form.SetBorder(true).SetTitle("UniBox v0.1").SetTitleAlign(tview.AlignCenter)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}

}
