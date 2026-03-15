package decorator

import (
	"testing"
)

func TestDecoratorPattern(t *testing.T) {
	tests := []struct {
		name     string
		beverage Beverage
		wantCost float64
		wantDesc string
	}{
		{
			name:     "Espresso",
			beverage: &Espresso{},
			wantCost: 1.99,
			wantDesc: "Espresso",
		},
		{
			name:     "House Blend",
			beverage: &HouseBlend{},
			wantCost: 0.89,
			wantDesc: "House Blend Coffee",
		},
		{
			name:     "Espresso with Mocha",
			beverage: &Mocha{Beverage: &Espresso{}},
			wantCost: 2.19,
			wantDesc: "Espresso, Mocha",
		},
		{
			name:     "Espresso with Whip",
			beverage: &Whip{Beverage: &Espresso{}},
			wantCost: 2.09,
			wantDesc: "Espresso, Whip",
		},
		{
			name:     "Espresso with Soy",
			beverage: &Soy{Beverage: &Espresso{}},
			wantCost: 2.14,
			wantDesc: "Espresso, Soy",
		},
		{
			name:     "House Blend with Soy, Mocha, Whip",
			beverage: &Whip{Beverage: &Mocha{Beverage: &Soy{Beverage: &HouseBlend{}}}},
			wantCost: 1.34,
			wantDesc: "House Blend Coffee, Soy, Mocha, Whip",
		},
		{
			name:     "Double Mocha",
			beverage: &Mocha{Beverage: &Mocha{Beverage: &Espresso{}}},
			wantCost: 2.39,
			wantDesc: "Espresso, Mocha, Mocha",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCost := tt.beverage.Cost()
			gotDesc := tt.beverage.GetDescription()

			if gotCost != tt.wantCost {
				t.Errorf("Cost() = %.2f, want %.2f", gotCost, tt.wantCost)
			}
			if gotDesc != tt.wantDesc {
				t.Errorf("GetDescription() = %q, want %q", gotDesc, tt.wantDesc)
			}
		})
	}
}
