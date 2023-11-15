package input_csv

import (
	"github.com/oufti2015/html/model"
	"log/slog"
	"strconv"
	"time"
)

func ParseOperations(data [][]string) []model.Operation {
	var operations []model.Operation
	for i, line := range data {

		if i > 0 {
			var operation model.Operation
			for j, field := range line {
				if j == 0 {
					operation.Id = field
				}
				layout := "02/01/2006"
				if j == 1 {
					if len(field) > 0 {
						parse, err := time.Parse(layout, field)
						operation.ExecutionDate = parse
						if err != nil {
							slog.Error("cannot parse Execution Date", field)
							panic(err)
						}
					}
				}
				if j == 2 {
					if len(field) > 0 {
						parse, err := time.Parse(layout, field)
						operation.ValueDate = parse
						if err != nil {
							slog.Error("cannot parse Value Date", field)
							panic(err)
						}
					}
				}
				if j == 3 {
					parse, err := strconv.ParseFloat(field, 64)
					operation.Amount = parse
					if err != nil {
						slog.Error("cannot parse ", field)
						panic(err)
					}
				}
				if j == 4 {
					operation.Currency = field
				}
				if j == 5 {
					operation.AccountNumber = field
				}
				if j == 6 {
					operation.TransactionType = field
				}
				if j == 7 {
					operation.CounterpartyAccountNumber = field
				}
				if j == 8 {
					operation.CounterpartyName = field
				}
				if j == 9 {
					operation.Communication = field
				}
				if j == 10 {
					operation.Details = field
				}
				if j == 11 {
					operation.Status = field
				}
				if j == 12 {
					operation.ReasonForRefusal = field
				}
			}
			operations = append(operations, operation)
		}
	}
	return operations
}
