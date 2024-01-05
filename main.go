package main

import (
	"DESAFIO-GO-BASES/internal/read"
	"DESAFIO-GO-BASES/internal/tickets"
	"fmt"
)

func main() {
	// read file and store it in a variable

	records, err := getRecords()
	if err != nil {
		return
	}

	// format records to tickets as type ticket
	//send records to getTickets function and store in variable GetTickets
	GetTickets, err := tickets.GetTickets(records)
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
