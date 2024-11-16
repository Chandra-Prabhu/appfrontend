package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
