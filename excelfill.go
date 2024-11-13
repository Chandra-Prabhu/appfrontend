package main

import (
	"fmt"
	"math"

	"github.com/xuri/excelize/v2"
)

func excelfill(model map[string][]float64) {
	wb := excelize.NewFile()
	defer func() {
		if err := wb.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheet := "Financials"
	wb.SetSheetName("Sheet1", sheet)
	excelstyling(wb, sheet)
	wb.SetColWidth(sheet, "D", "AF", 18)
	i := 2
	i = adddata(model, "Generation", wb, sheet, i)
	i = adddata(model, "Tariff", wb, sheet, i)
	i = adddata(model, "Revenue", wb, sheet, i)
	i++
	i = adddata(model, "Opex", wb, sheet, i)
	i = adddata(model, "EBITDA", wb, sheet, i) + 1
	i = adddata(model, "Interest paid", wb, sheet, i)
	i = adddata(model, "Depreciation", wb, sheet, i)
	i = adddata(model, "PBT", wb, sheet, i)
	i = adddata(model, "Tax", wb, sheet, i)
	i = adddata(model, "Profits before Dividend", wb, sheet, i) + 2
	i = adddata(model, "Capex", wb, sheet, i)
	i = adddata(model, "debtopening", wb, sheet, i)
	i = adddata(model, "debtoutstanding", wb, sheet, i)
	i = adddata(model, "debt repayment", wb, sheet, i)
	i = adddata(model, "DSCR", wb, sheet, i) + 1
	i = adddata(model, "Tax depreciation", wb, sheet, i)
	i = adddata(model, "Taxable income", wb, sheet, i) + 1
	i = adddata(model, "FCFE", wb, sheet, i) + 1
	i = adddata(model, "DSRA opening", wb, sheet, i)
	i = adddata(model, "DSRA closing", wb, sheet, i)
	i = adddata(model, "DSRA change", wb, sheet, i) + 1
	i = adddata(model, "Working capital", wb, sheet, i)
	_ = adddata(model, "Change in WC", wb, sheet, i)
	highlightrows := []int{4, 7, 11, 13, 25}
	highlightstyle(wb, sheet, highlightrows)
	wb.SaveAs("test2.1.xlsx")
}

func celladdress(col int, row int) string {
	var str string
	if col < 26 {
		str = fmt.Sprintf("%s%d", string(rune(64+col)), row)
	} else {
		str = fmt.Sprintf("%s%d", (string(rune(64+col/26)) + string(rune(65+int(math.Mod(float64(col), 26))))), row)
	}
	return str
}

func adddata(model map[string][]float64, dataname string, wb *excelize.File, sheet string, i int) int {
	data := model[dataname]
	wb.SetCellStr(sheet, celladdress(1, i), dataname)
	wb.SetSheetRow(sheet, celladdress(4, i), &data)
	return i + 1
}

func highlightstyle(wb *excelize.File, sheet string, k []int) {
	style1, _ := wb.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 12,
			Bold: true,
		},
		Border: []excelize.Border{
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 6},
		},
		NumFmt: 38,
	})
	for _, i := range k {
		wb.SetRowStyle(sheet, i, i, style1)
	}
}

func excelstyling(wb *excelize.File, sheet string) {
	styledef, _ := wb.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "FFFFFF", Style: 2},
			{Type: "right", Color: "FFFFFF", Style: 2},
			{Type: "top", Color: "FFFFFF", Style: 2},
			{Type: "bottom", Color: "FFFFFF", Style: 2},
		},
		NumFmt: 38,
	})
	styletitle, _ := wb.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size:   12,
			Italic: true,
		},
	})
	wb.SetCellStyle(sheet, "a1", "a100", styletitle)
	wb.SetRowStyle(sheet, 1, 100, styledef)
}
