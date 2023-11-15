package input_json

import (
	"encoding/json"
	"github.com/oufti2015/html/model"
	"log/slog"
)

func ParseCategories(data []byte) []model.Category {
	var categories model.Categories
	err := json.Unmarshal(data, &categories)
	if err != nil {
		slog.Error("cannot parse JSOM ")
		panic(err)
	}
	return categories.Categories
}
