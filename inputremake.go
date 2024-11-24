package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// assumption set
func assumptionbuild() map[string][]assumptions {
	as := make(map[string][]assumptions, 0)
	//debtscpoption := []string{"Equal", "Sculpted"}
	//tabs := []string{"Commercial", "Projects", "Financing", "Others"}
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Capacity", "MW"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("PPA Length", "years"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Construction Period", "month"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Tariff", "Rs./KWh"))
	as["Commercial"] = append(as["Commercial"], newAssumptionE("Tariff Escalation", "% p.a"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Interest rate", "%"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Debt as % of Capex", "%"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Minimum Debt repayment", "% p.a"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Minimum DSCR", "x of EBITDA"))
	as["Financing"] = append(as["Financing"], newAssumptionE("DSRA", "months"))
	as["Financing"] = append(as["Financing"], newAssumptionE("Debt Tenure", "months"))
	as["Financing"] = append(as["Financing"], newAssumptionS("Repayment method", []string{"Equal", "Sculpted"}))
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
	as = solar(as)
	as = wind(as)
	return as
}

func solar(as map[string][]assumptions) map[string][]assumptions {
	as["Solar"] = append(as["Solar"], newAssumptionT("Project Name"))
	as["Solar"] = append(as["Solar"], newAssumptionE("AC Capacity", "MW"))
	as["Solar"] = append(as["Solar"], newAssumptionE("DC Capacity", "MWdc"))
	as["Solar"] = append(as["Solar"], newAssumptionE("T/L length", "ckm"))
	as["Solar"] = append(as["Solar"], newAssumptionE("PV Degradation", "%"))
	as["Solar"] = append(as["Solar"], newAssumptionE("Capex phasing", "%"))
	as["Solar"] = append(as["Solar"], newAssumptionE("Land Capex", "Rs.L/MWdc"))
	as["Solar"] = append(as["Solar"], newAssumptionE("PV module Capex", "Rs.L/MWdc"))
	as["Solar"] = append(as["Solar"], newAssumptionE("BoS Capex", "Rs.L/MWdc"))
	as["Solar"] = append(as["Solar"], newAssumptionE("Currency Exposure", "%"))
	as["Solar"] = append(as["Solar"], newAssumptionE("Hedging rate", "%"))
	as["Solar"] = append(as["Solar"], newAssumptionE("Unit Opex", "Rs.L/MW/yr"))
	as["Solar"] = append(as["Solar"], newAssumptionE("CUF", "%"))
	as["Solar"] = append(as["Solar"], newAssumptionE("Opex escalation", "% p.a"))
	return as
}

func wind(as map[string][]assumptions) map[string][]assumptions {
	as["Wind"] = append(as["Wind"], newAssumptionT("Project Name"))
	as["Wind"] = append(as["Wind"], newAssumptionE("AC Capacity", "MW"))
	as["Wind"] = append(as["Wind"], newAssumptionE("T/L length", "ckm"))
	as["Wind"] = append(as["Wind"], newAssumptionE("Capex phasing", "%"))
	as["Wind"] = append(as["Wind"], newAssumptionE("Land Capex", "Rs.L/MW"))
	as["Wind"] = append(as["Wind"], newAssumptionE("WTG Capex", "Rs.L/MW"))
	as["Wind"] = append(as["Wind"], newAssumptionE("BoS Capex", "Rs.L/MW"))
	as["Wind"] = append(as["Wind"], newAssumptionE("Currency Exposure", "%"))
	as["Wind"] = append(as["Wind"], newAssumptionE("Hedging rate", "%"))
	as["Wind"] = append(as["Wind"], newAssumptionE("Unit Opex", "Rs.L./MW/yr"))
	as["Wind"] = append(as["Wind"], newAssumptionE("CUF", "%"))
	as["Wind"] = append(as["Wind"], newAssumptionE("Opex escalation", "% p.a"))
	return as
}

// Project renderer
func offtake(as map[string][]assumptions) {
	scenarioCasecounter = "New Project Setup"
	scenariotitle()
	k1 := newle("Offtake Name")
	var offtakeName, offtakePeriod string
	k1.Entry.OnSubmitted = func(a string) { offtakeName = a }
	k2 := newle("Offtake Period")
	k2.Entry.OnSubmitted = func(a string) { offtakePeriod = a }
	i := 1
	var l2 []fyne.CanvasObject
	l1 := widget.NewCard("Offtake commitment", "", widget.NewButtonWithIcon("Add new", theme.ContentAddIcon(), func() {
		l2 = append(l2, createCommitmentrules(i))
		i++
		if i == 3 {
			content1 = *inputrenderer(as, "Commercial")
		}
		content1.Refresh()
	}))
	content1 = *container.NewVBox(k1.widgetmaker(), k2.widgetmaker(), container.NewHBox(l1), container.NewVBox(l2...))
	fmt.Println(offtakeName, offtakePeriod)
}

type labelAndEntry struct {
	Label *widget.Label
	Entry *widget.Entry
}

type labelAndOption struct {
	Label  *widget.Label
	Select *widget.SelectEntry
}

func createCommitmentrules(i int) *fyne.Container {
	k1 := newle("Constraints -" + string(i) + " Name")
	k2 := newle("Constraint -" + string(i) + " CUF")
	k3 := newle("Constraint -" + string(i) + " time interval")
	k4 := newls("Constraint -"+string(i)+" type", []string{"hard constraint", "soft constraint"})
	return container.NewGridWithColumns(2, k1.widgetmaker(), k2.widgetmaker(), k3.widgetmaker(), k4.widgetmaker())
}

func (c labelAndEntry) widgetmaker() *fyne.Container {
	return container.NewGridWithRows(1, c.Label, c.Entry)
}

func (c labelAndOption) widgetmaker() *fyne.Container {
	return container.NewGridWithRows(1, c.Label, c.Select)
}

func newle(s string) labelAndEntry {
	return labelAndEntry{Label: widget.NewLabel(s), Entry: widget.NewEntry()}
}

func newls(s string, options []string) labelAndOption {
	return labelAndOption{Label: widget.NewLabel(s), Select: widget.NewSelectEntry(options)}
}

/*func projectConfig(as map[string][]assumptions) map[string][]assumptions {
	as["Projects"]
}*/
