package tickets

import "time"

type Ticket struct {
	id          string
	name        string
	email       string
	destination string
	time        string
	Price       string
}

type Tickets struct {
	ticketRecord []Ticket
}

func GetTickets(records [][]string) ([]Ticket, error) {
	var allTickets Tickets

	for _, record := range records {
		ticketRow := Ticket{
			id:          record[0],
			name:        record[1],
			email:       record[2],
			destination: record[3],
			time:        record[4],
			Price:       record[5],
		}
		allTickets.ticketRecord = append(allTickets.ticketRecord, ticketRow)
	}
	return allTickets.ticketRecord, nil
}

// ejemplo 1
func GetTotalTickets(tickets []Ticket) (map[string]int, error) {
	countryCount := make(map[string]int)
	for _, ticket := range tickets {
		countryCount[ticket.destination]++
	}

	return countryCount, nil
}

// ejemplo 2
func GetTimes(tickets []Ticket) (map[string]int, error) {
	timeMap := make(map[string]int)
	timeMap["morning"] = 0
	timeMap["afternoon"] = 0
	timeMap["night"] = 0
	timeMap["late_night"] = 0

	for _, ticket := range tickets {
		t, err := time.Parse("15:04", ticket.time)
		if err != nil {
			return nil, err
		}

		hour := t.Hour()
		switch {
		case hour >= 0 && hour < 6:
			timeMap["late_night"]++
		case hour >= 6 && hour < 12:
			timeMap["morning"]++
		case hour >= 12 && hour < 18:
			timeMap["afternoon"]++
		case hour >= 18 && hour < 24:
			timeMap["night"]++
		}
	}
	return timeMap, nil

}

// ejemplo 3
func AverageDestination(Tickets []Ticket) (map[string]float64, error) {
	destinationCount := make(map[string]int)
	totalTickets := len(Tickets)
	averageDestination := make(map[string]float64)

	for _, ticket := range Tickets {
		destinationCount[ticket.destination]++
	}

	for destination, count := range destinationCount {
		percentage := float64(count) / float64(totalTickets) * 100
		averageDestination[destination] = percentage
	}

	return averageDestination, nil
}
