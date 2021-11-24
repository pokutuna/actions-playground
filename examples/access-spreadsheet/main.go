package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()
	client, err := sheets.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1Vrmv8_h2npEkqVLZOjUlS-tKvy0vQW-hiY03_xWJuxo"
	resp, err := client.Spreadsheets.Values.Get(spreadsheetId, "シート1!A1:B2").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	for _, row := range resp.Values {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Printf("\n")
	}
}
