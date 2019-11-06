package main

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/sqweek/dialog"
)

func transUI() {
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())

	w := app.NewWindow("Hello")

	// Input Entry
	inputEntry := widget.NewEntry()
	inputFileSelectButton := widget.NewButton("...", func() {
		// filename, _ := dialog.File().Filter("XML files", "xml").Title("Export to XML").Save()
		filename, _ := dialog.File().Filter("All", "*").Title("Input File").Load()
		fmt.Println(filename)
		inputEntry.Text = filename
	})
	inputPart := widget.NewHBox(inputEntry, inputFileSelectButton)

	// Output Entry
	outputEntry := widget.NewEntry()
	outputEntryFileSelectButton := widget.NewButton("...", func() {
		// filename, _ := dialog.File().Filter("XML files", "xml").Title("Export to XML").Save()
		filename, _ := dialog.File().Filter("All", "*").Title("Input File").Save()
		fmt.Println(filename)
		outputEntry.Text = filename
	})
	outputPart := widget.NewHBox(outputEntry, outputEntryFileSelectButton)

	quitButton := widget.NewButton("Quit", func() {
		app.Quit()
	})

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		inputPart,
		outputPart,
		quitButton,
	))

	w.ShowAndRun()
}
