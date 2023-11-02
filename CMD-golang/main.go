package main

import (
	"fmt"
	"os"
	"strings"

	keyboard "github.com/eiannone/keyboard"
	colors "github.com/iVitaliya/colors-go"
	tview "github.com/rivo/tview"
)

var frames = []string{
	colors.BgBlack(colors.White("1.")) + " " + colors.BgBlack(colors.White("This is test 1.")),
	colors.BgBlack(colors.White("2.")) + " " + colors.BgBlack(colors.White("This is test 2.")),
	colors.BgBlack(colors.White("3.")) + " " + colors.BgBlack(colors.White("This is test 3.")),
}

func animate(ind int) string {
	arr := make([]string, len(frames))

	frames[ind] = colors.BgWhite(colors.Black(frames[ind]))

	return frames[ind]
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
		fmt.Println("\n" + strings.Join(frames, "\n"))
		var i int = 0

		app.QueueUpdateDraw(func() {
			textView.SetText(strings.Join(frames, "\n"))
		})

		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}

			switch key {
			case keyboard.KeyEsc:
				os.Exit(0)
			case keyboard.KeyArrowDown:
				if i != len(frames)-1 {
					i = i + 1
					animate(i)
				}
			case keyboard.KeyArrowUp:
				if i == 0 || i < len(frames)-1 {
					i = i - 1
					animate(i)
				}
			}
		}
	}()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
