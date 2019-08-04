package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/colinaaa/hust-cs-grade/internal"
)

const (
	idIndex   = 1
	nameIndex = 2
)

const output = "./data/out.xlsx"

func main() {
	husters, err := internal.LoadJSON("./data/cs18.json")
	f, err := internal.LoadXLS("./data/rank2018-2019.xlsx")
	if err != nil {
		log.Fatalln("error occurred:", err.Error())
	}
	rows := f.GetRows("sheet0")
	for index, cols := range rows[4:] {
		name := husters.FindNameByID(strings.TrimSpace(cols[idIndex]))
		f.SetCellStr("sheet0", "C"+fmt.Sprint(index+5), name)
	}
	if err := f.SaveAs(output); err != nil {
		panic(err.Error())
	}
}
