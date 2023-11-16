package main

import (
	"fmt"
	"github.com/oufti2015/html/dispatch"
	"github.com/oufti2015/html/input-csv"
	input_json "github.com/oufti2015/html/input-json"
	"github.com/oufti2015/html/model"
	"github.com/oufti2015/html/repo"
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

	/*	for _, category := range categories {
		if category.Name == "Voyages" {
			printCategory(category)
		}
	}*/

	dataRepository, _ := repo.NewRepository(categories, operations)

	err := dispatch.Categories(*dataRepository)
	if err != nil {
		panic(err)
	}
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

func printCategory(category model.Category) {
	slog.Info("name     : " + category.Name)
	slog.Info(fmt.Sprintf("priority : " + strconv.Itoa(category.Priority)))
	slog.Info(fmt.Sprintf("income   : %v", category.Income))
	slog.Info(fmt.Sprintf("saving   : %v", category.Saving))
	for i, criteria := range category.Criteria {
		slog.Info("Criteria " + strconv.Itoa(i) + " : " + criteria.Criteria + " - " + criteria.Comment)

		if !criteria.Period.StartTime().IsZero() {
			slog.Info(fmt.Sprintf("Period : %v -> %v",
				criteria.Period.StartTime().Format("2006-01-02"),
				criteria.Period.EndTime().Format("2006-01-02")))
		}
	}
}
