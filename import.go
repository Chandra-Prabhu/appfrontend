package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func importassumptions() map[string]string {
	file, _ := excelize.OpenFile("input.xlsx")
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheets := file.GetSheetList()
	//section := make([]string, 0)
	values := make(map[string]string, 0)
	//units := make([]string, 0)
	//values:=make([]string,0)
	rows, _ := file.GetRows(sheets[0])
	for _, row := range rows {
		_, err := strconv.Atoi(row[0])
		if err == nil {
			//section = append(section, row[1])
			values[row[2]] = row[4]
			//units = append(units, row[3])
		}
	}
	//abind:=[][]string {section,description,units,values}
	return values
}

func inputbuilder(values map[string]string, as map[string][]assumptions) {
	for _, tabs := range as {
		for _, assumption := range tabs {
			assumption.inputupdate(values[assumption.inputname()])
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
