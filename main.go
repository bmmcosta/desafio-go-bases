package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/read"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// read file and store in variable records
	records, err := getRecords()
	if err != nil {
		return
	}

	// format records to tickets
	GetTickets, err := tickets.GetTickets(records) //send records to getTickets function and store in variable tickets
	if err != nil {
		return
	}

	// get total tickets by destination
	fmt.Printf("\n****** Total tickets by destination: \n")
	totalTickets, err := tickets.GetTotalTickets(GetTickets)
	if err != nil {
		return
	}

	for Country, Count := range totalTickets {
		fmt.Printf("%s: %d\n", Country, Count)
	}

	// get total tickets by Time
	fmt.Printf("\n****** Total tickets by Time Range: \n")
	totalByTime, err := tickets.GetTimes(GetTickets)
	if err != nil {
		return
	}
	for Time, Count := range totalByTime {
		fmt.Printf("%s: %d\n", Time, Count)
	}
	// get average tickets by country
	averageTickets, err := tickets.AverageDestination(GetTickets)
	if err != nil {
		return
	}
	fmt.Printf("\n****** Average tickets by country: \n")
	for Country, Average := range averageTickets {
		fmt.Printf("%s: %.2f %% \n", Country, Average)
	}
}

// getRecords reads the file and returns the records
func getRecords() ([][]string, error) {
	records, err := read.ReadFile()
	if err != nil {
		return nil, err
	}

	return records, nil
}
