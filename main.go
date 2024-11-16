package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	//"fyne.io/fyne/v2/internal/theme"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var content1, casewindow, sctitle fyne.Container

// var sctitle fyne.Container
var cases []scenarios

//var scenariotitle string = "New Scenario"

func main() {
	//fmt.Println("HW")
	a := app.New()
	a.Settings().SetTheme(newFysionTheme())
	w := a.NewWindow("HW")
	w.Title()
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(viewmaker(w))
	w.ShowAndRun()
}

func viewmaker(w fyne.Window) fyne.CanvasObject {
	as, debstscpoption, tabs := assumptionbuild()
	tabsbuild := inputtabmaker(as, tabs)
	submits := actionbuttons(w, as, debstscpoption)
	inputbuilder(importassumptions(), as)
	content1 = *inputrenderer(as, "Commercial")
	scenariotitle("New Scenario")
	caserenderer(as)
	return container.NewStack(container.NewBorder(tabsbuild, container.NewGridWithRows(1, submits...), nil, nil, container.NewVBox(&sctitle, &content1)), &casewindow)
}

// assumption set
func assumptionbuild() (map[string][]assumptions, []string, []string) {
	as := make(map[string][]assumptions, 0)
	debtscpoption := []string{"Equal", "Sculpted"}
	tabs := []string{"Commercial", "Projects", "Financing", "Others"}
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Capacity", "MW"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("PPA Length", "years"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Construction Period", "months"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Tariff", "Rs./KWh"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Tariff Escalation", "% p.a"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Interest rate", "%"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Debt as % of Capex", "%"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Minimum Debt repayment p.a", "% p.a"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Minimum DSCR", "x of EBITDA"))
	as["Financing"] = append(as["Financing"], newAssumptionE("DSRA", "months"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Debt Tenure", "months"))
	as["Financing"] = append(as["Financing"], newAssumptionS("Repayment method", debtscpoption))
	//as["Financing"] = append(as["Financing"], newAssumptionE("Min Debt repayment", "% p.a"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Capacity", "MW"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Unit Capex", "Rs.Cr./MW"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Unit Opex", "Rs.Cr./MW/yr"))
	as["Projects"] = append(as["Projects"], newAssumptionE("CUF", "%"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Degradation", "% p.a"))
	as["Projects"] = append(as["Projects"], newAssumptionE("Opex escalation", "% p.a"))
	as["Others"] = append(as["Others"], newAssumptionE("Corporate tax", "%"))
	as["Others"] = append(as["Others"], newAssumptionE("O&M GST", "%"))
	as["Others"] = append(as["Others"], newAssumptionS("Depreciation method", []string{"SLM", "Diminshing Balance"}))
	as["Others"] = append(as["Others"], newAssumptionE("Book Depreciation rate", "%"))
	as["Others"] = append(as["Others"], newAssumptionE("Tax Depreciation rate", "%"))
	as["Others"] = append(as["Others"], newAssumptionE("Non Depreciable Value", "%"))
	as["Others"] = append(as["Others"], newAssumptionE("Payables", "days"))
	as["Others"] = append(as["Others"], newAssumptionE("Receivables", "days"))
	return as, debtscpoption, tabs
}

func inputtabmaker(as map[string][]assumptions, tabs []string) *fyne.Container {
	inputrenderer(as, tabs[0])
	click := tabs[0]
	tabsbtn := make([]fyne.CanvasObject, len(tabs))
	for i := 0; i < len(tabs); i++ {
		tabsbtn[i] = widget.NewButton(tabs[i], func() {
			if click != tabs[i] {
				content1 = *inputrenderer(as, tabs[i])
				click = tabs[i]
			}
		})
	}
	return container.NewGridWithRows(1, tabsbtn...)
}

func actionbuttons(w fyne.Window, as map[string][]assumptions, debtscpoption []string) []fyne.CanvasObject {
	inputs := make(map[string]float64, 0)
	var case1 scenarios
	submitbtn := make([]fyne.CanvasObject, 0)
	submitbtn = append(submitbtn, widget.NewButton("Run Scenario", func() {
		inputs = inputgrab(as)
		IRRmake(IRRmodel(inputs, debtscpoption))
	}))
	submitbtn = append(submitbtn, widget.NewButton("Run and Save Scenario", func() {
		inputs = inputgrab(as)
		case1.Inputs = inputGrabAsStr(as)
		case1.Model = IRRmodel(inputs, debtscpoption)
		case1.Irr = IRRmake(case1.Model)
		a := scenarioname(w, case1, as)
		a.Show()
	}))
	submitbtn = append(submitbtn, widget.NewButton("Save As Excel", func() {
		inputs = inputgrab(as)
		excelfill(IRRmodel(inputs, debtscpoption))
	}))
	return submitbtn
}

// takes assumptions as per category into widgets and displays
func inputrenderer(as map[string][]assumptions, selectedTab string) *fyne.Container {
	var wid1 []fyne.CanvasObject
	for i := range as[selectedTab] {
		abds := (as[selectedTab][i]).inputmaker()
		wid1 = append(wid1, abds)
	}
	fmt.Println(selectedTab, "was pressed")
	return container.NewVBox(wid1...)
}

func scenariotitle(title1 string) {
	//var k appLabelWidget
	k := widget.NewLabel(title1)
	k.Alignment = fyne.TextAlignCenter
	k.TextStyle = fyne.TextStyle{Bold: true,
		Underline: true}
	kb := canvas.NewRectangle(hexColor("#8AF3A4FF"))
	sctitle = *container.New(layout.NewStackLayout(), kb, k)
}

// creates input window
/*func inputwindow(w fyne.Window) *fyne.Container {


	//var toptab *fyne.Container
	//toptab = container.NewGridWithRows(1, tabsbtn...)
	scenariotitle("New Scenario")
	caserenderer(as)

	bottomtab := container.NewGridWithRows(1, widget.NewLabel(""), submit, savecase, excel, widget.NewLabel(""))
	inputbuilder(importassumptions(), as)
	middleportion := container.NewBorder(nil, bottomtab, nil, nil, container.NewVBox(toptab, &sctitle, &content1))
	//middleportion.Resize(fyne.NewSize(300, 400))
	return container.NewBorder(nil, nil, nil, &casewindow, middleportion)
}*/

type entryassumptions struct {
	Name  string
	Unit  string
	Entry *widget.Entry
}

type selectassumptions struct {
	Name   string
	Unit   string
	Select *widget.SelectEntry
	Option []string
}

type scenarios struct {
	Name   string
	Inputs map[string]string
	Model  map[string][]float64
	Irr    float64
}

type assumptions interface {
	inputmaker() fyne.CanvasObject
	inputsave() float64
	inputname() string
	inputupdate(string)
	inputSaveAsStr() string
}

// once submit button is triggered it fetches inputs from each entries
func inputgrab(as map[string][]assumptions) map[string]float64 {
	inputs := make(map[string]float64, 0)
	for _, tabs := range as {
		for _, assumption := range tabs {
			inputs[assumption.inputname()] = assumption.inputsave()
		}
	}
	return inputs
}

// fetch the label of the assumption
func (assumption entryassumptions) inputname() string {
	return assumption.Name
}
func (assumption selectassumptions) inputname() string {
	return assumption.Name
}

// all the inputs are getting registered
func (assumption entryassumptions) inputsave() float64 {
	var input1 float64
	a := assumption.Entry.Text
	if string(a[len(a)-1]) == "%" {
		a = a[:len(a)-1]
	}
	input1, _ = strconv.ParseFloat(a, 64)
	return input1
}

// all the options selected are getting registered
func (assumption selectassumptions) inputsave() float64 {
	var o float64
	for i, k := range assumption.Option {
		if k == assumption.Select.Text {
			o = float64(i)
		}
	}
	return o
}

// new assumption with inputable struct is created
func newAssumptionE(name string, unit string) entryassumptions {
	assumption := entryassumptions{}
	assumption.Name = name
	assumption.Unit = unit
	assumption.Entry = widget.NewEntry()
	return assumption
}

// new assumption with options to select struct is created
func newAssumptionS(name string, options []string) selectassumptions {
	assumption := selectassumptions{}
	assumption.Name = name
	assumption.Select = widget.NewSelectEntry(options)
	assumption.Option = options
	return assumption
}

// creates input & label
func (assumption entryassumptions) inputmaker() fyne.CanvasObject {
	assumption.Entry.PlaceHolder = "Enter the " + assumption.Name
	ac := widget.NewLabel(fmt.Sprint(assumption.Name + " in " + assumption.Unit))
	abds := container.NewAdaptiveGrid(2, ac, assumption.Entry)
	return abds
}

// creates input with options & Label
func (assumption selectassumptions) inputmaker() fyne.CanvasObject {
	ac := widget.NewLabel(fmt.Sprint("Select ", assumption.Name, " Method"))
	abds := container.NewAdaptiveGrid(2, ac, assumption.Select)
	return abds
}

// create formdialog for scenario name
func scenarioname(w fyne.Window, case1 scenarios, as map[string][]assumptions) *dialog.FormDialog {
	a := widget.NewEntry()
	a.SetPlaceHolder("Type Here..")
	scenarioinput := widget.FormItem{
		Text:   "Scenario Name",
		Widget: a,
	}
	wap := dialog.NewForm("Enter the Scenario name", "Okay", "Close",
		[]*widget.FormItem{&scenarioinput},
		func(b bool) {
			if b {
				// Get input text
				case1.Name = a.Text
				cases = append(cases, case1)
				//fmt.Println(cases[len(cases)-1].Name)
				caserenderer(as)
			}
		}, w)
	//fmt.Println(case1.Name)
	return wap
}

// creates the case window on the left side as it gets saved
func caserenderer(as map[string][]assumptions) {
	outer := make([]fyne.CanvasObject, 0)
	//var name string
	for k, i := range cases {
		a := widget.NewLabel(i.Name)
		b := widget.NewLabel(fmt.Sprintf("%.1f", i.Irr*100) + " %")
		a.Resize(fyne.NewSize(100, 100))
		b.Resize(fyne.NewSize(50, 100))
		c := widget.NewButton("Select"+i.Name, func() {
			scenariotitle(i.Name)
			inputbuilder(i.Inputs, as)
		})
		c.Resize(fyne.NewSize(100, 100))
		var containerColor *canvas.Rectangle
		if (k/2)*2 == k {
			containerColor = canvas.NewRectangle(hexColor("#F4F5D4FF"))
		} else {
			containerColor = canvas.NewRectangle(hexColor("#83E2C6FF"))
		}
		container1 := container.NewGridWithRows(1, a, b, c)
		containerColor.Resize(container1.Size())
		container2 := container.New(layout.NewStackLayout(), containerColor, container1)
		outer = append(outer, container2)
	}
	casewindow = *container.NewVBox(outer...)
	casewindow.Resize(fyne.NewSize(100, 200))
}

// once submit button is triggered it fetches inputs from each entries
func inputGrabAsStr(as map[string][]assumptions) map[string]string {
	inputs := make(map[string]string, 0)
	for _, tabs := range as {
		for _, assumption := range tabs {
			inputs[assumption.inputname()] = assumption.inputSaveAsStr()
		}
	}
	return inputs
}

// all the inputs are getting registered
func (assumption entryassumptions) inputSaveAsStr() string {
	return assumption.Entry.Text
}

// all the options selected are getting registered
func (assumption selectassumptions) inputSaveAsStr() string {
	return assumption.Select.Text
}
