package main

import (
	"fmt"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {

	opt := option.WithCredentialsFile("service-account.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "ronsovia", opt)
	if err != nil {
		fmt.Println(err)
	}

	sheetName := "Answers"
	excelFile := excelize.NewFile()
	excelFile.NewSheet(sheetName)
	currRow := 1

	iter := client.Collection("answers").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		for k, v := range doc.Data() {
			fmt.Println("k:", k, "v:", v)
			if k == "soal_1" {
				excelFile.SetCellValue(sheetName, "B"+strconv.Itoa(currRow), v)
			} else if k == "soal_2" {
				excelFile.SetCellValue(sheetName, "C"+strconv.Itoa(currRow), v)
			} else if k == "soal_3" {
				excelFile.SetCellValue(sheetName, "D"+strconv.Itoa(currRow), v)
			} else if k == "soal_4" {
				excelFile.SetCellValue(sheetName, "E"+strconv.Itoa(currRow), v)
			} else if k == "soal_5" {
				excelFile.SetCellValue(sheetName, "F"+strconv.Itoa(currRow), v)
			}
		}
	}

	if err := excelFile.SaveAs("Answers.xlsx"); err != nil {
		fmt.Println(err)
	}
}
