package main

import (
	"fmt"
	"image/color"
	"strconv"

	//"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"

	//"fyne.io/fyne/v2/internal/theme"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var content1, casewindow, sctitle, tabsbuild, submitbtns, finaldisplay fyne.Container
var titelcol color.RGBA = hexColor("#B9C3FDFD")
var casecounter int = 1
var scenarioCasecounter string = "New Scenario"

// var sctitle fyne.Container
var cases = make(map[string]scenarios)

//var scenariotitle string = "New Scenario"

func main() {
	a := app.New()
	a.Settings().SetTheme(newFysionTheme())
	w := a.NewWindow("HW")
	w.Title()
	w.Resize(fyne.NewSize(400, 400))
	viewmaker(w)
	w.SetContent(&finaldisplay)
	w.ShowAndRun()
}

/*func projectConfig(){
	//offtake contract name
	//offtakecontractlength
	//compliances add
	//complaince1 name,value,timeline
	//ProjectConfig add
	//Element one name,type, capacity,location,interconnection,delete,comment
	//Element two same
}*/

func viewmaker(w fyne.Window) {

	as := assumptionbuild()
	inputtabmaker(as)
	actionbuttons(w, as)
	if scenarioCasecounter == "New Scenario" {
		inputbuilder(importassumptions(), as)
	} else {
		inputbuilder(cases[scenarioCasecounter].Inputs, as)
	}
	offtake(as)
	scenariotitle()
	caserenderer(as)
	makeGUI()
	finaldisplay.Refresh()
}

func inputtabmaker(as map[string][]assumptions) {
	tabs := make([]string, 0)
	for i := range as {
		tabs = append(tabs, i)
	}
	inputrenderer(as, tabs[0])
	click := tabs[0]
	tabsbtn := make([]fyne.CanvasObject, len(tabs))
	for i := 0; i < len(tabs); i++ {
		tabsbtn[i] = widget.NewButton(tabs[i], func() {
			if click != tabs[i] {
				content1 = *inputrenderer(as, tabs[i])
				finaldisplay.Refresh()
				//makeGUI()
				click = tabs[i]
			}
		})
	}
	tabsbuild = *container.NewGridWithRows(2, tabsbtn...)
}

func actionbuttons(w fyne.Window, as map[string][]assumptions) {
	inputs := make([]inputs, 0)
	submitbtn := make([]fyne.CanvasObject, 0)
	submitbtn = append(submitbtn, widget.NewButton("Run Scenario", func() {
		inputs = inputGrabAsStr(as)
		IRRmake(IRRmodel(inputs))
	}))
	submitbtn = append(submitbtn, widget.NewButton("Save Scenario", func() {
		inputs = inputGrabAsStr(as)
		var case1 scenarios
		case1.Inputs = inputGrabAsStr(as)
		case1.Model = IRRmodel(inputs)
		case1.Irr = IRRmake(case1.Model)
		if scenarioCasecounter == "New Scenario" {
			a := scenarioNamePopup(w, case1, as)
			a.Show()
		} else {
			case1.Name = cases[scenarioCasecounter].Name
			cases[scenarioCasecounter] = case1
			caserenderer(as)
			//makeGUI()
		}
	}))
	submitbtn = append(submitbtn, widget.NewButton("Save As NewScenario", func() {
		inputs = inputGrabAsStr(as)
		var case1 scenarios
		case1.Inputs = inputGrabAsStr(as)
		case1.Model = IRRmodel(inputs)
		case1.Irr = IRRmake(case1.Model)
		a := scenarioNamePopup(w, case1, as)
		a.Show()
	}))
	submitbtn = append(submitbtn, widget.NewButton("Save As Excel", func() {
		inputs = inputGrabAsStr(as)
		excelfill(IRRmodel(inputs))
	}))
	submitbtns = *container.NewGridWithRows(1, submitbtn...)
}

// takes assumptions as per category into widgets and displays
func inputrenderer(as map[string][]assumptions, selectedTab string) *fyne.Container {
	var wid1 []fyne.CanvasObject
	for i := range as[selectedTab] {
		abds := (as[selectedTab][i]).inputmaker()
		wid1 = append(wid1, abds)
	}
	return container.NewVBox(wid1...)
}

func scenariotitle() {
	k := widget.NewLabel("")
	_, err := strconv.Atoi(scenarioCasecounter)
	if err != nil {
		k.SetText(scenarioCasecounter)
	} else {
		k.SetText(cases[scenarioCasecounter].Name)
	}
	k.Alignment = fyne.TextAlignCenter
	k.TextStyle = fyne.TextStyle{Bold: true, Underline: true}
	kb := canvas.NewRectangle(titelcol)
	sctitle = *container.NewVBox(widget.NewSeparator(), container.New(layout.NewStackLayout(), kb, k), widget.NewSeparator())
	//sctitle.Move(fyne.Position{X: 0, Y: 45})
	//sctitle.Resize(fyne.NewSize(460, 30))
}

// create formdialog for scenario name
func scenarioNamePopup(w fyne.Window, case1 scenarios, as map[string][]assumptions) *dialog.FormDialog {
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
				cases[strconv.Itoa(casecounter)] = case1
				scenarioCasecounter = strconv.Itoa(casecounter)
				caserenderer(as)
				scenariotitle()
				casecounter++
				finaldisplay.Refresh()
				//makeGUI()
			}
		}, w)
	return wap
}

func removecase(k string) map[string]scenarios {
	cases1 := make(map[string]scenarios, 0)
	if len(cases) > 1 {
		for i := range cases {
			if i != k {
				cases1[i] = cases[i]
			}
		}
	}
	return cases1
}

func removecasecont(casedisplay map[string]fyne.CanvasObject, k string) map[string]fyne.CanvasObject {
	cases1 := make(map[string]fyne.CanvasObject)
	//kint, _ := strconv.Atoi(k)
	if len(casedisplay) > 1 {
		for i := range casedisplay {
			if i != k {
				cases1[i] = casedisplay[i]
			}
		}
	}
	return cases1
}

// creates the case window on the left side as it gets saved
func caserenderer(as map[string][]assumptions) {
	casewindow.RemoveAll()
	outer := make(map[string]fyne.CanvasObject, 0)
	//var name string
	//casebtn := make([]*widget.Button, 0)
	//casecontainer := make([]*fyne.Container, 0)
	for k, i := range cases {
		a := widget.NewLabel(trunc(i.Name))
		b := widget.NewLabel(fmt.Sprintf("%.1f", i.Irr*100) + " %")
		//a.Resize(fyne.NewSize(100, 100))
		//b.Resize(fyne.NewSize(50, 100))
		c := widget.NewButton("Select", func() {
			scenarioCasecounter = k
			scenariotitle()
			inputbuilder(i.Inputs, as)
			casewindowbuild(outer)
			//name = i.Name
		})
		d := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			cases = removecase(k)
			fmt.Println(len(cases))
			//caserenderer(as, w)
			//viewmaker(w)
			outer = removecasecont(outer, k)
			if (scenarioCasecounter == k) || (len(cases) == 0) {
				scenarioCasecounter = "New Scenario"
				fmt.Println("New Scenario")
				scenariotitle()
				finaldisplay.Refresh()
				//	name = ""
			}
			casewindowbuild(outer)
		})
		//c.Resize(fyne.NewSize(100, 100))

		container1 := container.NewGridWithRows(1, a, b, c, d)
		outer[k] = container1
		//casecontainer = append(casecontainer, container2)
		//outer = append(outer, container2)
	}
	casewindowbuild(outer)
}
func casewindowbuild(outer map[string]fyne.CanvasObject) {
	casewindow.RemoveAll()
	ct := container.New(layout.NewStackLayout(), canvas.NewRectangle(titelcol), widget.NewLabelWithStyle("Saved Cases", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Underline: true}))
	casebg := canvas.NewRectangle(hexColor("#EBF7BDFF"))
	outer1 := make([]fyne.CanvasObject, 0)
	//casecount := 0
	keys := make([]string, 0)
	for i := range outer {
		keys = append(keys, i)
	}
	sortkeys(keys)
	for _, k := range keys {
		var containerColor *canvas.Rectangle
		if k == scenarioCasecounter {
			containerColor = canvas.NewRectangle(hexColor("#C2F4F7FF"))
		} else {
			containerColor = canvas.NewRectangle(hexColor("#F2F594FF"))
		}
		container2 := container.New(layout.NewStackLayout(), containerColor, outer[k])
		outer1 = append(outer1, container2, widget.NewSeparator())
		//casecount++
	}
	d := widget.NewSeparator()
	d.CreateRenderer()
	casewindow = *container.New(layout.NewStackLayout(), casebg, container.NewBorder(widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), nil, container.NewVBox(ct, widget.NewSeparator(), container.NewVBox(outer1...))))
	//casewindow.Resize(fyne.NewSize(200, 200))
	casewindow.Refresh()
	finaldisplay.Refresh()
	//makeGUI()
}

// sort the numbers which are stored as string
func sortkeys(keys []string) {
	//keysSorted := make([]string, len(keys))
	keysAsInt := make([]int, 0)
	for _, k := range keys {
		i, _ := strconv.Atoi(k)
		keysAsInt = append(keysAsInt, i)
	}
	BubbleSort(keysAsInt)
	for i, k := range keysAsInt {
		keys[i] = strconv.Itoa(k)
	}
}

func BubbleSort(data []int) {
	size := len(data)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1; j++ {
			Swap(data, j)
		}
	}
}

func Swap(data []int, index int) {
	if data[index] > data[index+1] {
		data[index], data[index+1] = data[index+1], data[index]
	}
}

func trunc(word string) string {
	if len(word) > 10 {
		return word[:10] + ".."
	} else {
		return word
	}
}
