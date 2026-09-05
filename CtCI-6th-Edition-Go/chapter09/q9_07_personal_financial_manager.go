package chapter09

import (
	"fmt"
	"strings"
)

type Transaction struct {
	Amount      float64
	Description string
}

type PersonalFinancialManager struct {
	rules [][2]string
}

func NewPersonalFinancialManager() *PersonalFinancialManager {
	return &PersonalFinancialManager{}
}

func (m *PersonalFinancialManager) AddRule(keyword, category string) {
	if keyword == "" || category == "" {
		panic("keyword and category required")
	}
	m.rules = append(m.rules, [2]string{strings.ToLower(keyword), category})
}

func (m *PersonalFinancialManager) Categorize(transaction Transaction) string {
	description := strings.ToLower(transaction.Description)
	for _, rule := range m.rules {
		if strings.Contains(description, rule[0]) {
			return rule[1]
		}
	}
	return "uncategorized"
}

func (m *PersonalFinancialManager) Summary(transactions []Transaction) map[string]float64 {
	totals := map[string]float64{}
	for _, transaction := range transactions {
		category := m.Categorize(transaction)
		totals[category] += transaction.Amount
	}
	return totals
}

func RunQ907() {
	manager := NewPersonalFinancialManager()
	manager.AddRule("grocery", "food")
	manager.AddRule("uber", "transport")
	transactions := []Transaction{{-50.0, "Whole Foods grocery"}, {-20.0, "Uber ride"}}
	fmt.Println(manager.Categorize(transactions[0]))
	fmt.Printf("food=%v\n", manager.Summary(transactions)["food"])
}
