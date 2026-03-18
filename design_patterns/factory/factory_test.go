package main

import (
	"bytes"
	"io"
	"os"
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pizza, err := tt.factory.CreatePizza(cheesePizza)
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
	store := NewPizzaStore(NewChicagoPizzaFactory())

	output := captureStdout(t, func() {
		pizza, err := store.OrderPizza(cheesePizza)
		if err != nil {
			t.Fatalf("OrderPizza() error = %v", err)
		}

		if got, want := pizza.GetName(), "Chicago Style Deep Dish Cheese Pizza"; got != want {
			t.Fatalf("GetName() = %q, want %q", got, want)
		}
	})

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
	})

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

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	originalStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() error = %v", err)
	}

	os.Stdout = writer

	outputCh := make(chan string, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, reader)
		outputCh <- buf.String()
	}()

	fn()

	os.Stdout = originalStdout
	_ = writer.Close()

	return <-outputCh
}
