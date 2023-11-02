package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	keyboard "github.com/eiannone/keyboard"
	colors "github.com/iVitaliya/colors-go"
	tview "github.com/rivo/tview"
)

var frames = []string{
	colors.BgWhite(colors.Black("1.")) + " " + colors.BgBlack(colors.White("This is test 1.")),
	colors.BgWhite(colors.Black("2.")) + " " + colors.BgBlack(colors.White("This is test 2.")),
	colors.BgWhite(colors.Black("3.")) + " " + colors.BgBlack(colors.White("This is test 3.")),
}

func animate() {

}

func main() {
	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetText("Animating...").
		SetTextAlign(tview.AlignCenter).
		SetSize(30, 30)

	app.SetRoot(textView, true)

	go func() {
		err := keyboard.Open()
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = keyboard.Close()
		}()

		fmt.Println("Press ESC to exit the menu.")

		app.QueueUpdateDraw(func() {
			textView.SetText(strings.Join(frames, "\n"))
		})

		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}

			switch key {
			case keyboard.KeyEsc:
				os.Exit(0)
			case 
			} 
		}
	}()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
