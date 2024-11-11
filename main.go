package main

import (
	"fmt"
	"strconv"

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
	w.SetContent(inputwindow())
	w.ShowAndRun()
}

func inputwindow() *fyne.Container {
	as := make(map[string][]assumptions, 0)
	tabs := []string{"Commercial", "Financing", "Projects", "Tax"}
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Capacity", "MW"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("PPA Length", "years"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Construction Period", "months"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Tariff", "Rs./KWh"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Tariff Escalation", "% p.a"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Interest rate", "%"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Debt as % of Capex", "%"))
	as["Financing"] = append(as["Financing"], newAssumptionS("Repayment method", []string{"Equal", "Sculpted"}))
	as["Financing"] = append(as["Financing"], newAssumptionE("Minimum Debt repayment p.a", "% p.a"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Minimum DSCR", ""))
	as["Financing"] = append(as["Financing"], newAssumptionE("DSRA", "months"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Payables", "days"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Receivables", "days"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Debt Tenure", "months"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Capacity", "MW"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Unit Capex", "Rs.Cr./MW"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Unit Opex", "Rs.Cr./MW/yr"))
	as["Projects"] = append(as["Projects"], newAssumptionE("CUF", "%"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Opex escalation", "% p.a"))
	as[tabs[3]] = append(as[tabs[3]], newAssumptionE("Corporate tax", "%"))
	as[tabs[3]] = append(as[tabs[3]], newAssumptionE("O&M GST", "%"))
	as[tabs[3]] = append(as[tabs[3]], newAssumptionS("Depreciation method", []string{"SLM", "Diminshing Balance"}))
	as[tabs[3]] = append(as[tabs[3]], newAssumptionE("Book Depreciation rate", "%"))
	as[tabs[3]] = append(as[tabs[3]], newAssumptionE("Tax Depreciation rate", "%"))
	as[tabs[3]] = append(as[tabs[3]], newAssumptionE("Non Depreciable Value", "%"))
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
	toptab = container.NewGridWithRows(1, tabsbtn...)
	content1 = *display(wid1)
	submit := widget.NewButton("Submit", func() {
		inputgrab(as)
	})
	return container.NewBorder(toptab, submit, nil, nil, &content1)
}

type entryassumptions struct {
	Name   string
	Fvalue float64
	Ovalue int
	Unit   string
	Entry  *widget.Entry
}

type selectassumptions struct {
	Name   string
	Unit   string
	Select *widget.SelectEntry
	Option string
}

type assumptions interface {
	inputmaker() fyne.CanvasObject
	inputsave()
}

func inputgrab(as map[string][]assumptions) {
	for _, tabs := range as {
		for _, assumption := range tabs {
			assumption.inputsave()
		}
	}
}
func (assumption entryassumptions) inputsave() {
	if (assumption.Unit == "months") || (assumption.Unit == "years") {
		assumption.Ovalue, _ = strconv.Atoi(assumption.Entry.Text)
	} else {
		assumption.Fvalue, _ = strconv.ParseFloat(assumption.Entry.Text, 64)
	}
	if string(assumption.Unit[0]) == "%" {
		a := assumption.Entry.Text
		a = a[:len(a)-1]
		assumption.Fvalue, _ = strconv.ParseFloat(a, 64)
		assumption.Fvalue = assumption.Fvalue / 100.0
	}
}
func (assumption selectassumptions) inputsave() {
	assumption.Option = assumption.Select.SelectedText()
}

func newAssumptionE(name string, unit string) entryassumptions {
	assumption := entryassumptions{}
	assumption.Name = name
	assumption.Unit = unit
	assumption.Fvalue = 0
	assumption.Ovalue = 0
	assumption.Entry = widget.NewEntry()
	return assumption
}
func newAssumptionS(name string, options []string) selectassumptions {
	assumption := selectassumptions{}
	assumption.Name = name
	assumption.Select = widget.NewSelectEntry(options)
	assumption.Option = options[0]
	return assumption
}

func (assumption entryassumptions) inputmaker() fyne.CanvasObject {
	assumption.Entry.PlaceHolder = "Enter the " + assumption.Name
	ac := widget.NewLabel(fmt.Sprint(assumption.Name + " in " + assumption.Unit))
	//theme.ColorForWidget("widgetColor", ac)
	abds := container.NewAdaptiveGrid(2, ac, assumption.Entry)
	return abds
}
func (assumption selectassumptions) inputmaker() fyne.CanvasObject {
	//assumption.Entry.PlaceHolder = "Enter the " + assumption.Name
	ac := widget.NewLabel(fmt.Sprint("Select ", assumption.Name, " Method"))
	//theme.ColorForWidget("widgetColor", ac)
	abds := container.NewAdaptiveGrid(2, ac, assumption.Select)
	return abds
}

func inputrenderer(as map[string][]assumptions, selectedTab string) []*fyne.CanvasObject {
	var wid1 []*fyne.CanvasObject
	for i := range as[selectedTab] {
		abds := (as[selectedTab][i]).inputmaker()
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
