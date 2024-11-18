package main

import (
	"fyne.io/fyne/v2/container"
)

/*
func makeGUI(tabsbuild fyne.CanvasObject, submits fyne.CanvasObject) fyne.CanvasObject {
	middle := container.NewVBox(&sctitle, tabsbuild, &content1, submits)
	left := widget.NewLabel("Left")
	right := &casewindow
	objs := []fyne.CanvasObject{middle, left, right}
	return container.New(newAppLayout(middle, left, right), objs...)
}
*/

func makeGUI() {
	finaldisplay = *container.New(newAppLayout(&sctitle, &tabsbuild, &content1, &submitbtns, &casewindow), &sctitle, &tabsbuild, &content1, &submitbtns, &casewindow)
	finaldisplay.Refresh()
}
