package main

import (
	"github.com/rivo/tview"
)

func init() {

}

func main() {

	box := tview.NewBox().
		SetBorder(true).
		SetTitle("AUTOPOMr")

	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}

}
