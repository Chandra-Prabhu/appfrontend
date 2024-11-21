package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type entryassumptions struct {
	Name  string
	Unit  string
	Entry *widget.Entry
	Type  string
}

type textassumptions struct {
	Name  string
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
	Inputs []inputs // first is section header, second is value, third is type int or float
	Model  map[string][]float64
	Irr    float64
	//Select *widget.Button
	//Delete *widget.Button
	//Container *fyne.Container
}
type inputs struct {
	Section   string
	Attribute string
	Value     string
	Type      string
}

type assumptions interface {
	inputmaker() fyne.CanvasObject
	//inputsave() float64
	inputname() string
	inputupdate(string)
	inputSaveAsStr() string
	inputType() string
}

// new assumption with inputable struct is created
func newAssumptionE(name string, unit string) entryassumptions {
	assumption := entryassumptions{}
	assumption.Name = name
	assumption.Unit = unit
	assumption.Entry = widget.NewEntry()
	if (unit == "years") || (unit == "months") {
		assumption.Type = "int"
	} else {
		assumption.Type = "float"
	}
	return assumption
}

func newAssumptionT(name string) textassumptions {
	assumption := textassumptions{}
	assumption.Name = name
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

/*
// once submit button is triggered it fetches inputs from each entries
func inputgrab(as map[string][]assumptions) map[string]float64 {
	inputs := make(map[string]float64, 0)
	for _, tabs := range as {
		for _, assumption := range tabs {
			inputs[assumption.inputname()] = assumption.inputsave()
		}
	}
	return inputs
}*/

// fetch the label of the assumption
func (assumption entryassumptions) inputname() string {
	return assumption.Name
}
func (assumption selectassumptions) inputname() string {
	return assumption.Name
}
func (assumption textassumptions) inputname() string {
	return assumption.Name
}

/*
// all the inputs are getting registered
func (assumption entryassumptions) inputsave() float64 {
	//var input1 float64
	a := assumption.Entry.Text
	if string(a[len(a)-1]) == "%" {
		a = a[:len(a)-1]
	}
	input1, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Please enter valid data, using default values")
		assumption.Entry.SetText("")
		values := importassumptions()
		a := values[assumption.inputname()]
		assumption.Entry.SetText(a)
		if string(a[len(a)-1]) == "%" {
			a = a[:len(a)-1]
		}
		input1, _ = strconv.ParseFloat(a, 64)
	}
	return input1
}

func (assumption textassumptions) inputsave() float64 {
	//var input1 float64
	return 0
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
}*/

// creates input & label
func (assumption entryassumptions) inputmaker() fyne.CanvasObject {
	assumption.Entry.PlaceHolder = "Enter the " + assumption.Name
	ac := widget.NewLabel(fmt.Sprint(assumption.Name + " in " + assumption.Unit))
	ad := canvas.NewText("", hexColor("#585858FD"))
	assumption.Entry.OnChanged = func(a string) {
		if len(a) > 1 {
			if string(a[len(a)-1]) == "%" {
				a = a[:len(a)-1]
			}
		}
		_, err := strconv.ParseFloat(a, 64)
		if err != nil {
			ad.Text = "Please enter the correct data"
			ad.Color = hexColor("#FA2A2AFD")
			assumption.Entry.SetText("")
			ad.Refresh()
		} else {
			ad.Text = ""
			ad.Color = hexColor("#585858FD")
			ad.Refresh()
		}
	}
	abds := container.New(newAsLayout(ac, assumption.Entry, ad), ac, assumption.Entry, ad)
	return abds
}

// creates input with options & Label
func (assumption selectassumptions) inputmaker() fyne.CanvasObject {
	ac := widget.NewLabel(fmt.Sprint("Select ", assumption.Name, " Method"))
	ad := canvas.NewText("", hexColor("#585858FD"))
	abds := container.New(newAsLayout(ac, assumption.Select, ad), ac, assumption.Select, ad)
	return abds
}

// creates input which would have text & Label
func (assumption textassumptions) inputmaker() fyne.CanvasObject {
	assumption.Entry.PlaceHolder = "Enter the " + assumption.Name
	ac := widget.NewLabel(assumption.Name)
	//ac := widget.NewLabel(fmt.Sprint("Select ", assumption.Name, " Method"))
	ad := canvas.NewText("", hexColor("#585858FD"))
	abds := container.New(newAsLayout(ac, assumption.Entry, ad), ac, assumption.Entry, ad)
	return abds
}

// once submit button is triggered it fetches inputs from each entries
func inputGrabAsStr(as map[string][]assumptions) []inputs {
	inputs1 := make([]inputs, 0)
	var input1 inputs
	for k, tabs := range as {
		for _, assumption := range tabs {
			input1.Attribute = assumption.inputname()
			input1.Section = k
			input1.Value = assumption.inputSaveAsStr()
			input1.Type = assumption.inputType()
			inputs1 = append(inputs1, input1)
		}
	}
	return inputs1
}

// all the inputs are getting registered
func (assumption entryassumptions) inputSaveAsStr() string {
	return assumption.Entry.Text
}

// all the inputs are getting registered
func (assumption textassumptions) inputSaveAsStr() string {
	return assumption.Entry.Text
}

// all the options selected are getting registered
func (assumption selectassumptions) inputSaveAsStr() string {
	return assumption.Select.Text
}

// all the inputs are getting registered
func (assumption entryassumptions) inputType() string {
	return assumption.Type
}

// all the inputs are getting registered
func (assumption textassumptions) inputType() string {
	return "string"
}

// all the options selected are getting registered
func (assumption selectassumptions) inputType() string {
	return "select"
}
