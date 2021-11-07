package main

<<<<<<< HEAD
func main() {
	//Build your strategies in the strategy.go file and pass the functions into MainTick()
	// go strategy.MainTick()
=======
import (
	"github.com/rivo/tview"
	"main.go/strategy"
	"main.go/swap"
)

func main() {
	userConfig := loadConfig()

	if !userConfig["initialized"].(bool) {
		setTrackedTokens()
		setTrackedPools()
	}

	var ANSCIIUniCorn string = "   UniBox         \n      \\                \n       _\\,,            \n      \"-=\\~     _  \n         \\~___( ~\n        _|/---\\_   \n        \\        \\     \n Credit (ejm97)"
	var tokensToApprove []interface{} = userConfig["needsApprove"].([]interface{})
	for index := 0; index < len(tokensToApprove); index++ {
		swap.Approve(tokensToApprove[index].(string))
	}

	go strategy.MainTick()
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
>>>>>>> 997fe36d7a7f2c2d6bacd051fbf48e63c69129a2

}
