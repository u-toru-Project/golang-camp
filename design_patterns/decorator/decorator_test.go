package decorator

import "testing"

func TestDecoratorPattern(t *testing.T) {
	tests := []struct {
		name     string
		beverage Beverage
		wantCost float64
		wantDesc string
	}{
		{
			name:     "Espresso",
			beverage: NewEspresso(),
			wantCost: 1.99,
			wantDesc: "Espresso",
		},
		{
			name:     "House Blend",
			beverage: NewHouseBlend(),
			wantCost: 0.89,
			wantDesc: "House Blend Coffee",
		},
		{
			name:     "Espresso with Mocha",
			beverage: NewMocha(NewEspresso()),
			wantCost: 2.19,
			wantDesc: "Espresso, Mocha",
		},
		{
			name:     "Espresso with Whip",
			beverage: NewWhip(NewEspresso()),
			wantCost: 2.09,
			wantDesc: "Espresso, Whip",
		},
		{
			name:     "Espresso with Soy",
			beverage: NewSoy(NewEspresso()),
			wantCost: 2.14,
			wantDesc: "Espresso, Soy",
		},
		{
			name:     "House Blend with Soy, Mocha, Whip",
			beverage: NewWhip(NewMocha(NewSoy(NewHouseBlend()))),
			wantCost: 1.34,
			wantDesc: "House Blend Coffee, Soy, Mocha, Whip",
		},
		{
			name:     "Double Mocha",
			beverage: NewMocha(NewMocha(NewEspresso())),
			wantCost: 2.39,
			wantDesc: "Espresso, Mocha, Mocha",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.beverage.Cost(); got != tt.wantCost {
				t.Errorf("Cost() = %.2f, want %.2f", got, tt.wantCost)
			}
			if got := tt.beverage.GetDescription(); got != tt.wantDesc {
				t.Errorf("GetDescription() = %q, want %q", got, tt.wantDesc)
			}
		})
	}
}
