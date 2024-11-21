package main

import (
	"errors"
	"strconv"

	//"fmt"
	"math"
)

func convfloat(s string) float64 {
	if string(s[len(s)-1]) == "%" {
		s = s[:len(s)-1]
		c, _ := strconv.ParseFloat(s, 64)
		return c / 100
	} else {
		c, _ := strconv.ParseFloat(s, 64)
		return c
	}
}

func convint(s string) int {
	c, _ := strconv.Atoi(s)
	return c
}

func IRRmodel(values []inputs) map[string][]float64 {
	inputs := make(map[string]string)
	for _, k := range values {
		inputs[k.Attribute] = k.Value
	}
	//fmt.Println("Enter PPA length, tariff")
	pPALength := convint(inputs["PPA Length"])
	constrperiod := convint(inputs["Construction Period"]) / 12
	tariff := convfloat(inputs["Tariff"])
	//var tariff float64 = inputs["Tariff"]
	intrate := convfloat(inputs["Interest rate"])
	mindebtrepay := convfloat(inputs["Minimum Debt repayment"])
	repaymethod := convint(inputs["Repayment method"])
	//fmt.Println(repaymethod)
	dscr := convfloat(inputs["Minimum DSCR"])
	payables := convfloat(inputs["Payables"]) / 365
	receivables := convfloat(inputs["Receivables"]) / 365
	debttenure := convint(inputs["Debt Tenure"]) / 12
	capacity := convfloat(inputs["Capacity"])
	unitCapex := convfloat(inputs["Unit Capex"])
	unitOpex := convfloat(inputs["Unit Opex"])
	costunit := 10000000.0
	cuf := convfloat(inputs["CUF"])
	degradation := convfloat(inputs["Degradation"])
	tariffEscalation := convfloat(inputs["Tariff Escalation"])
	opexEscalation := convfloat(inputs["Opex escalation"])
	gstrate := convfloat(inputs["O&M GST"])
	taxrate := convfloat(inputs["Corporate tax"])
	de := convfloat(inputs["Debt as % of Capex"])
	deprerate := convfloat(inputs["Book Depreciation rate"])
	txdeprerate := convfloat(inputs["Tax Depreciation rate"])
	nondeprecap := convfloat(inputs["Non Depreciable Value"])
	dsra := convfloat(inputs["DSRA"]) / 12
	model := make(map[string][]float64)
	model["Generation"] = constrappend(gencal(capacity, cuf, degradation, pPALength), constrperiod)
	model["Tariff"] = constrappend(tariffcal(tariff, tariffEscalation, pPALength), constrperiod)
	model["Opex"] = constrappend(tariffcal(unitOpex*capacity*(1.0+gstrate)*costunit, opexEscalation, pPALength), constrperiod)
	model["Revenue"] = revenuecal(model["Generation"], model["Tariff"])
	model["EBITDA"] = minus(model["Revenue"], model["Opex"])
	capex := capacity * unitCapex * costunit
	model["Capex"] = make([]float64, constrperiod)
	model["Capex"][0] = capex / float64(constrperiod)
	model["Capex"][1] = capex / float64(constrperiod)
	initialloan := capex * de
	model["debt repayment"], model["debtopening"], model["debtoutstanding"], model["Interest paid"], model["DSCR"] = debtrepay(initialloan, debttenure, repaymethod, model["EBITDA"][2:], intrate, dscr, mindebtrepay)
	model["debt repayment"] = constrappend(model["debt repayment"], constrperiod)
	model["debtopening"] = constrappend(model["debtopening"], constrperiod)
	model["debtoutstanding"] = constrappend(model["debtoutstanding"], constrperiod)
	model["Interest paid"] = constrappend(model["Interest paid"], constrperiod)
	model["PBDT"] = minus(model["EBITDA"], model["Interest paid"])
	model["Working capital"], model["Change in WC"] = workingcapcal(model["Revenue"], model["Opex"], payables, receivables)
	_, model["Depreciation"] = depreciationslm(capex-nondeprecap, deprerate, pPALength)
	model["Depreciation"] = constrappend(model["Depreciation"], constrperiod)
	model["PBT"] = minus(model["PBDT"], model["Depreciation"])
	model["DSRA opening"], model["DSRA closing"], model["DSRA change"] = dsracal(model["Interest paid"], model["debt repayment"], dsra)
	model["debt repayment"][0] = -model["Capex"][0] * de
	model["debt repayment"][1] = -model["Capex"][1] * de
	_, model["Tax depreciation"] = depreciationslm(capex-nondeprecap, txdeprerate, pPALength)
	model["Tax depreciation"] = constrappend(model["Tax depreciation"], constrperiod)
	model["Taxable income"] = minus(model["PBDT"], model["Tax depreciation"])
	model["Tax"] = tax(model["Taxable income"], taxrate)
	model["Profits before Dividend"] = minus(model["PBT"], model["Tax"])
	model["FCFE"] = minus(minus(model["EBITDA"], model["Interest paid"]), add(add(model["Capex"], model["debt repayment"]), minus(model["Tax"], add(model["Change in WC"], model["DSRA change"]))))
	//fmt.Println(model["Profits before Dividend"][0])
	return model
}

func workingcapcal(revenue []float64, opex []float64, payables float64, receivables float64) ([]float64, []float64) {
	workingcap := make([]float64, len(revenue))
	changeinwc := make([]float64, len(revenue))
	for i := 0; i < len(revenue); i++ {
		workingcap[i] = revenue[i]*payables - opex[i]*receivables
		if i == 0 {
			changeinwc[i] = 0
		} else {
			changeinwc[i] = workingcap[i-1] - workingcap[i]
		}
	}
	return workingcap, changeinwc
}

func revenuecal(gen []float64, tariff []float64) []float64 {
	revenue := make([]float64, len(gen))
	for i := 0; i < len(gen); i++ {
		revenue[i] = gen[i] * tariff[i]
	}
	return revenue
}

func gencal(cap float64, cuf float64, degrad float64, ppaLength int) []float64 {
	gen := make([]float64, ppaLength)
	for i := 0; i < ppaLength; i++ {
		gen[i] = cap * cuf * 8760.0 * 1000 * math.Pow(1.0-degrad, float64(i))
	}
	return gen
}

func tariffcal(tariff float64, escalation float64, ppaLength int) []float64 {
	tariffts := make([]float64, ppaLength)
	for i := 0; i < ppaLength; i++ {
		tariffts[i] = tariff * math.Pow((1.0+escalation), float64(i))
	}
	return tariffts
}

func minus(from []float64, sub []float64) []float64 {
	size := min(len(from), len(sub))
	to := make([]float64, 0)
	for i := 0; i < size; i++ {
		to = append(to, (from[i] - sub[i]))
	}
	if len(from) < len(sub) {
		for i := size; i < len(sub); i++ {
			to = append(to, -sub[i])
		}
	} else {
		for i := size; i < len(from); i++ {
			to = append(to, from[i])
		}
	}
	return to
}

func add(from []float64, sub []float64) []float64 {
	size := min(len(from), len(sub))
	to := make([]float64, 0)
	for i := 0; i < size; i++ {
		to = append(to, from[i]+sub[i])
	}
	if len(from) < len(sub) {
		for i := size; i < len(sub); i++ {
			to = append(to, sub[i])
		}
	} else {
		for i := size; i < len(from); i++ {
			to = append(to, from[i])
		}
	}
	return to
}

func depreciationslm(deprecapex float64, deprerate float64, pPALength int) ([]float64, []float64) {
	fadp := make([]float64, pPALength)
	deprec := make([]float64, pPALength)
	//var i float64
	fadp[0] = deprecapex - deprecapex*deprerate
	deprec[0] = deprecapex * deprerate
	for j := 1; j < pPALength; j++ {
		if fadp[j-1] > deprecapex*deprerate {
			fadp[j] = fadp[j-1] - deprecapex*deprerate
		} else {
			fadp[j] = 0
		}
		deprec[j] = fadp[j-1] - fadp[j]
	}
	return fadp, deprec
}

func tax(taxableincome []float64, taxrate float64) []float64 {
	taxes := make([]float64, len(taxableincome))
	losscarry := make([]float64, len(taxableincome))
	for i := 0; i < len(taxableincome); i++ {
		if i == 0 {
			if taxableincome[i] < 0 {
				losscarry[i] = taxableincome[i]
			} else {
				taxes[i] = taxableincome[i] * taxrate
			}
		} else if taxableincome[i] < 0 {
			losscarry[i] = losscarry[i-1] + taxableincome[i]
			taxes[i] = 0
		} else if taxableincome[i] < losscarry[i-1] {
			losscarry[i] = losscarry[i-1] - taxableincome[i]
			taxes[i] = 0
		} else {
			losscarry[i] = 0
			taxes[i] = (taxableincome[i] - losscarry[i-1]) * taxrate
		}
	}
	return taxes
}

func constrappend(series []float64, constrperiod int) []float64 {
	toseries := make([]float64, constrperiod+len(series))
	for i := 0; i < (constrperiod + len(series)); i++ {
		if i < constrperiod {
			toseries[i] = 0
		} else {
			toseries[i] = series[i-constrperiod]
		}
	}
	return toseries
}

const (
	irrMaxInterations = 20
	irrAccuracy       = 1e-7
	irrInitialGuess   = 0
)

func IRRmake(model map[string][]float64) float64 {
	s, _ := IRR(model["FCFE"])
	//fmt.Printf("%f is the EIRR for given assumptions", s)
	return s
}

// IRR returns the Internal Rate of Return (IRR).
func IRR(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("values must include the initial investment (usually negative number) and period cash flows")
	}
	x0 := float64(irrInitialGuess)
	var x1 float64
	for i := 0; i < irrMaxInterations; i++ {
		fValue := float64(0)
		fDerivative := float64(0)
		for k := 0; k < len(values); k++ {
			fk := float64(k)
			fValue += values[k] / math.Pow(1.0+x0, fk)
			fDerivative += -fk * values[k] / math.Pow(1.0+x0, fk+1.0)
		}
		x1 = x0 - fValue/fDerivative
		if math.Abs(x1-x0) <= irrAccuracy {
			return x1, nil
		}
		x0 = x1
	}
	return x0, errors.New("could not find irr for the provided values")
}

func debtrepay(initialloan float64, debttenure int, method int, ebitda1 []float64, intrate float64, dscr float64, mindebtrepay float64) ([]float64, []float64, []float64, []float64, []float64) {
	//fmt.Println(ebitda1)
	debtrepayment := make([]float64, debttenure)
	debtoutstanding := make([]float64, debttenure)
	debtopening := make([]float64, debttenure)
	interest := make([]float64, debttenure)
	dscrts := make([]float64, debttenure)
	if method == 0 {
		debtopening[0] = initialloan
		debtrepayment[0] = initialloan / float64(debttenure)
		debtoutstanding[0] = debtopening[0] - debtrepayment[0]
		interest[0] = (debtopening[0] + debtoutstanding[0]) / 2.0 * intrate
		dscrts[0] = ebitda1[0] / (debtrepayment[0] + interest[0])
		for i := 1; i < debttenure; i++ {
			debtrepayment[i] = initialloan / float64(debttenure)
			debtopening[i] = debtoutstanding[i-1]
			debtoutstanding[i] = debtopening[i] - debtrepayment[i]
			interest[i] = (debtopening[i] + debtoutstanding[i]) / 2.0 * intrate
			dscrts[i] = ebitda1[i] / (debtrepayment[i] + interest[i])
		}
	} else {
		maxpay := maxrepay(ebitda1[:debttenure], dscr, intrate)
		for i := 0; i < debttenure; i++ {
			if i == 0 {
				debtopening[i] = initialloan
			} else {
				debtopening[i] = debtoutstanding[i-1]
			}
			if (debtopening[i] < maxpay[i]) && ((debtopening[i] - mindebtrepay*initialloan) < maxpay[i+1]) {
				debtrepayment[i] = mindebtrepay * initialloan
			} else if debtopening[i] < maxpay[i] {
				debtrepayment[i] = debtopening[i] - maxpay[i+1]
			} else {
				debtrepayment[i] = (ebitda1[i]/dscr - intrate*debtopening[i]) / (1 - intrate/2.0)
			}
			debtoutstanding[i] = debtopening[i] - debtrepayment[i]
			interest[i] = (debtopening[i] + debtoutstanding[i]) / 2.0 * intrate
			dscrts[i] = ebitda1[i] / (debtrepayment[i] + interest[i])
		}
	}
	return debtrepayment, debtopening, debtoutstanding, interest, dscrts
}

func maxrepay(ebitda []float64, dscr float64, intrate float64) []float64 {
	maxpay := make([]float64, len(ebitda))
	maxpay[len(ebitda)-1] = 0.0
	for i := len(ebitda) - 2; i >= 0; i-- {
		maxpay[i] = (2*ebitda[i]/dscr + maxpay[i+1]*(2.0-intrate)) / (2.0 + intrate)
	}
	//fmt.Println(maxdebtopening)
	return maxpay
}

func dsracal(interestpayment []float64, debtrepayment []float64, dsra float64) ([]float64, []float64, []float64) {
	dsraopening := make([]float64, len(interestpayment))
	dsraclosing := make([]float64, len(interestpayment))
	dsrachange := make([]float64, len(interestpayment))
	for i := 0; i < len(interestpayment); i++ {
		if i < 1 {
			dsraopening[i] = 0
		} else {
			dsraopening[i] = dsraclosing[i-1]
		}
		dsraclosing[i] = (interestpayment[i] + debtrepayment[i]) * dsra
		dsrachange[i] = dsraopening[i] - dsraclosing[i]
	}
	return dsraopening, dsraclosing, dsrachange
}
