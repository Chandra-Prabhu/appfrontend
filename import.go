package main

/*
1)picking up right input into the right section while reading from excel
2) screen design when there are more inputs in the section
3)handling blank submission
4) adding revenue calculation
5) output dashboard
6) adding formula to excel
7) compare two cases
*/
import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func importassumptions() []inputs {
	file, _ := excelize.OpenFile("input.xlsx")
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheets := file.GetSheetList()
	//section := make([]string, 0)
	values := make([]inputs, 0)
	var value inputs
	//units := make([]string, 0)
	//values:=make([]string,0)
	rows, _ := file.GetRows(sheets[0])
	for _, row := range rows {
		_, err := strconv.Atoi(row[0])
		if err == nil {
			var k string
			if (row[3] == "years") || (row[3] == "months") {
				k = "int"
			} else if row[3] == "select" {
				k = "select"
			} else if row[3] == "" {
				k = "string"
			} else {
				k = "float"
			}
			value.Section = row[1]
			value.Attribute = row[2]
			value.Value = row[4]
			value.Type = k
			values = append(values, value)
		}
	}
	return values
}

func structToMap(values []inputs) map[string][]string {
	k := make(map[string][]string, 0)
	for _, l := range values {
		k[l.Section+l.Attribute] = []string{l.Value, l.Type}
	}
	return k
}

func inputbuilder(values []inputs, as map[string][]assumptions) {
	k := structToMap(values)
	for section, tabs := range as {
		for _, assumption := range tabs {
			assumption.inputupdate(k[section+assumption.inputname()][0])
		}
	}
}

// submit the entries onto the screen
func (assumption entryassumptions) inputupdate(value string) {
	assumption.Entry.SetText(value)
}
func (assumption selectassumptions) inputupdate(value string) {
	assumption.Select.SetText(value)
}
func (assumption textassumptions) inputupdate(value string) {
	assumption.Entry.SetText(value)
}
