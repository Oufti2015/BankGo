package dispatch

import (
	"github.com/oufti2015/html/model"
	"github.com/oufti2015/html/repo"
	"sort"
	"strings"
)

func Categories(repository repo.Repository) error {
	sort.Slice(repository.Categories, func(i, j int) bool {
		return repository.Categories[i].Priority < repository.Categories[j].Priority
	})
	sort.Slice(repository.Operations, func(i, j int) bool {
		return repository.Operations[i].ExecutionDate.Before(repository.Operations[j].ExecutionDate)
	})

	for _, operation := range repository.Operations {
		chooseACategory(repository.Categories, &operation)
	}
	return nil
}

func chooseACategory(categories []model.Category, operation *model.Operation) {
	for _, category := range categories {
		if !operation.Dispatched {
			if dispatchCategory(category, operation) {
				return
			}
		}
	}
}

func dispatchCategory(category model.Category, operation *model.Operation) bool {
	for _, criteria := range category.Criteria {
		if positiveAmount(category, operation) && checkCriteria(criteria, operation) {
			operation.Category = category
			return true
		}
	}
	return false
}

func checkCriteria(criteria model.Criteria, operation *model.Operation) bool {
	if criteria.ImpactedField == "COUNTERPARTY_ACCOUNT" {
		if operation.CounterpartyAccountNumber == criteria.Criteria {
			return attachCriteria(criteria, operation)
		}
	}
	if criteria.ImpactedField == "COUNTERPARTY_NAME" {
		if strings.Contains(operation.CounterpartyName, criteria.Criteria) {
			return attachCriteria(criteria, operation)
		}
	}
	if criteria.ImpactedField == "DETAIL" {
		if strings.Contains(operation.Details, criteria.Criteria) {
			return attachCriteria(criteria, operation)
		}
	}
	if criteria.ImpactedField == "VALUE_DATE" {
		if (operation.ValueDate.Equal(criteria.Period.StartTime()) || operation.ValueDate.After(criteria.Period.StartTime())) &&
			(operation.ValueDate.Equal(criteria.Period.EndTime()) || operation.ValueDate.Before(criteria.Period.EndTime())) {
			return attachCriteria(criteria, operation)
		}
	}
	if criteria.ImpactedField == "TRANSACTION_TYPE" {
		if operation.TransactionType == criteria.Criteria {
			return attachCriteria(criteria, operation)
		}
	}
	if criteria.ImpactedField == "ALL" {
		return true
	}
	return false
}

func attachCriteria(criteria model.Criteria, operation *model.Operation) bool {
	operation.Dispatched = true
	operation.Criteria = criteria
	return true
}

func positiveAmount(category model.Category, operation *model.Operation) bool {
	return !category.Income || operation.Amount > 0.00
}
