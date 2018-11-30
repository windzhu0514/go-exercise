package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

func TestReadxlsx(t *testing.T) {

	f, err := xlsx.OpenFile("123.xlsx")
	if err != nil {
		panic(err)
	}

	for _, sheet := range f.Sheets {
		for rowIndex, row := range sheet.Rows {
			fmt.Printf("%d\t", rowIndex)
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\t", text)
			}
			fmt.Println()
		}
	}
}

func TestReadcsv(t *testing.T) {

	f, err := os.Open("123.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

func TestReadxls(t *testing.T) {
	wb, err := xls.Open("123.xls", "utf-8")
	if err != nil {
		panic(err)
	}

	sheet := wb.GetSheet(0)
	maxRow := int(sheet.MaxRow)
	for i := 0; i < maxRow; i++ {
		fmt.Printf("%d", i)
		maxCol := sheet.Row(i).LastCol()
		for j := 0; j < maxCol; j++ {
			text := sheet.Row(i).Col(j)
			fmt.Printf("\t%s", text)
		}

		fmt.Println()
	}

}
