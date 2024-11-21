package main

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
	return as
}
