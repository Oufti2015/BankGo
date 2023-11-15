package main

import (
	"github.com/oufti2015/html/input-csv"
	input_json "github.com/oufti2015/html/input-json"
	"github.com/oufti2015/html/model"
	"log/slog"
	"strconv"
)

func main() {
	const operationFilename = "C:\\zt974\\bank\\2023.csv"
	const categoryFilename = "C:\\zt974\\bank\\categories.json"

	slog.Info("starting")

	operations := readOperations(operationFilename)
	categories := readCategories(categoryFilename)

	slog.Info(strconv.Itoa(len(operations)) + " operations read")
	slog.Info(strconv.Itoa(len(categories)) + " operations read")
	slog.Info("end")
}

func readCategories(categoryFilename string) []model.Category {
	data, err := input_json.ReadFile(categoryFilename)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	categories := input_json.ParseCategories(data)
	return categories
}

func readOperations(operationFilename string) []model.Operation {
	data, err := input_csv.ReadFile(operationFilename)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	operations := input_csv.ParseOperations(data)
	return operations
}
