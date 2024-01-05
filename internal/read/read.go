package read

import (
	"encoding/csv"
	"os"
)

func ReadFile() ([][]string, error) {
	f, err := os.Open("././input/tickets.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
