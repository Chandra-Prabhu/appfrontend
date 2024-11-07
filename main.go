package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var wid1 []*fyne.CanvasObject
var content1 fyne.Container

func main() {
	fmt.Println("HW")
	a := app.New()
	a.Settings().SetTheme(newFysionTheme())
	w := a.NewWindow("HW")
	w.Title()
	w.Resize(fyne.NewSize(300, 50))
	//ab := widget.NewEntry()
	//ab.Resize(fyne.NewSize(150, 50))
	as := make(map[string][]assumptions, 0)
	as["Commercial"] = append(as["Commercial"], newAssumption("Capacity", "MW"))
	as["Commercial"] = append(as["Commercial"], newAssumption("PPA Length", "years"))
	as["Commercial"] = append(as["Commercial"], newAssumption("Construction Period", "years"))
	as["Commercial"] = append(as["Commercial"], newAssumption("tariff", "Rs./KWh"))
	as["Financing"] = append(as["Financing"], newAssumption("Interest rate", "%"))
	as["Financing"] = append(as["Financing"], newAssumption("repayment method", "select"))
	as["Financing"] = append(as["Financing"], newAssumption("Minimum Debt repayment in a year", "%"))
	as["Financing"] = append(as["Financing"], newAssumption("Minimum DSRA", ""))
	tabs := []string{"Commercial", "Financing"}
	wid1 = inputrenderer(as, tabs[0])
	var toptab *fyne.Container
	click := tabs[0]
	tabsbtn := make([]fyne.CanvasObject, len(tabs))
	for i := 0; i < len(tabs); i++ {
		tabsbtn[i] = widget.NewButton(tabs[i], func() {
			if click != tabs[i] {
				wid1 = inputrenderer(as, tabs[i])
				content1 = *display(wid1)
				click = tabs[i]
			}
		})
	}
	toptab = container.NewHBox(tabsbtn...)
	//abText := widget.NewLabel("Capacity in MW")
	//ab.SetPlaceHolder("Enter the capacity")
	//ab1 := container.NewAdaptiveGrid(2, abText, ab)
	content1 = *display(wid1)
	content := container.NewBorder(toptab, nil, nil, nil, &content1)
	w.SetContent(content)
	w.ShowAndRun()
}

type assumptions struct {
	Name   string
	Fvalue float64
	Ovalue int
	Unit   string
	Type   int
	Entry  *widget.Entry
}

func newAssumption(name string, unit string) assumptions {
	assumption := assumptions{}
	assumption.Name = name
	assumption.Unit = unit
	assumption.Fvalue = 0
	assumption.Ovalue = 0
	assumption.Type = 0
	assumption.Entry = widget.NewEntry()
	return assumption
}

func inputmaker(assumption assumptions) fyne.CanvasObject {
	assumption.Entry.PlaceHolder = "Enter the " + assumption.Name
	ac := widget.NewLabel(fmt.Sprint(assumption.Name + " in " + assumption.Unit))
	abds := container.NewAdaptiveGrid(2, ac, assumption.Entry)
	return abds
}

func inputrenderer(as map[string][]assumptions, selectedTab string) []*fyne.CanvasObject {
	var wid1 []*fyne.CanvasObject
	for i := range as[selectedTab] {
		abds := inputmaker(as[selectedTab][i])
		wid1 = append(wid1, &abds)
	}
	return wid1
}

func display(wid1 []*fyne.CanvasObject) *fyne.Container {
	wid2 := make([]fyne.CanvasObject, 0)
	for i := 0; i < len(wid1); i++ {
		wid2 = append(wid2, *wid1[i])
	}
	return container.NewVBox(wid2...)
}
