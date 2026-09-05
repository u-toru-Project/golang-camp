package factory

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestRegionalFactoriesCreateCheesePizza(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		factory  PizzaFactory
		wantName string
		wantType any
	}{
		{
			name:     "NY",
			factory:  NewNYPizzaFactory(),
			wantName: "NY Style Sauce and Cheese Pizza",
			wantType: &NYStyleCheesePizza{},
		},
		{
			name:     "Chicago",
			factory:  NewChicagoPizzaFactory(),
			wantName: "Chicago Style Deep Dish Cheese Pizza",
			wantType: &ChicagoStyleCheesePizza{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pizza, err := tt.factory.CreatePizza(CheesePizza)
			if err != nil {
				t.Fatalf("CreatePizza() error = %v", err)
			}

			if got := pizza.GetName(); got != tt.wantName {
				t.Fatalf("GetName() = %q, want %q", got, tt.wantName)
			}

			if got, want := reflect.TypeOf(pizza), reflect.TypeOf(tt.wantType); got != want {
				t.Fatalf("CreatePizza() type = %v, want %v", got, want)
			}
		})
	}
}

func TestRegionalFactoryCreatePizzaUnknownType(t *testing.T) {
	t.Parallel()

	factory := &regionalPizzaFactory{
		style: "Test",
		menu:  map[string]pizzaBuilder{},
	}

	pizza, err := factory.CreatePizza("unknown")
	if err == nil {
		t.Fatal("CreatePizza() error = nil, want error")
	}

	if pizza != nil {
		t.Fatalf("CreatePizza() pizza = %v, want nil", pizza)
	}

	if got, want := err.Error(), `Test factory: unknown pizza type "unknown"`; got != want {
		t.Fatalf("CreatePizza() error = %q, want %q", got, want)
	}
}

func TestPizzaStoreOrderPizza(t *testing.T) {
	var out bytes.Buffer
	store := NewPizzaStore(NewChicagoPizzaFactory(), &out)

	pizza, err := store.OrderPizza(CheesePizza)
	if err != nil {
		t.Fatalf("OrderPizza() error = %v", err)
	}

	if got, want := pizza.GetName(), "Chicago Style Deep Dish Cheese Pizza"; got != want {
		t.Fatalf("GetName() = %q, want %q", got, want)
	}

	output := out.String()
	for _, want := range []string{
		"--- Making a Chicago Style Deep Dish Cheese Pizza ---",
		"Preparing Chicago style dough, sauce, and cheese",
		"Bake for 25 minutes at 350",
		"Cutting the pizza into square slices",
		"Place pizza in official PizzaStore box",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("OrderPizza() output missing %q in %q", want, output)
		}
	}
}

func TestPizzaStoreOrderPizzaPropagatesFactoryError(t *testing.T) {
	t.Parallel()

	store := NewPizzaStore(&regionalPizzaFactory{
		style: "Test",
		menu:  map[string]pizzaBuilder{},
	}, &bytes.Buffer{})

	pizza, err := store.OrderPizza("unknown")
	if err == nil {
		t.Fatal("OrderPizza() error = nil, want error")
	}

	if pizza != nil {
		t.Fatalf("OrderPizza() pizza = %v, want nil", pizza)
	}

	if got, want := err.Error(), `Test factory: unknown pizza type "unknown"`; got != want {
		t.Fatalf("OrderPizza() error = %q, want %q", got, want)
	}
}

func TestNewPizzaStorePanicsWhenWriterIsMissing(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic when writer is missing")
		}
	}()

	NewPizzaStore(NewNYPizzaFactory(), nil)
}
