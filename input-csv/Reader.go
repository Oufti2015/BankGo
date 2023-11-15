package input_csv

import (
	"encoding/csv"
	"log/slog"
	"os"
)

func ReadFile(filename string) (data [][]string, err error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("cannot read " + filename)
	}
	defer file.Close()

	// Read CSV file using csv.Reader
	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	data, err = csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return data, err
}
