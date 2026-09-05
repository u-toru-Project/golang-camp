package factory

import (
	"fmt"
	"io"
)

const CheesePizza = "cheese"

// Pizza defines the operations every concrete pizza must provide.
type Pizza interface {
	Prepare(w io.Writer)
	Bake(w io.Writer)
	Cut(w io.Writer)
	Box(w io.Writer)
	GetName() string
}

type basePizza struct {
	name           string
	prepareMessage string
	cutMessage     string
}

func (p *basePizza) Prepare(w io.Writer) {
	fmt.Fprintln(w, p.prepareMessage)
}

func (p *basePizza) Bake(w io.Writer) {
	fmt.Fprintln(w, "Bake for 25 minutes at 350")
}

func (p *basePizza) Cut(w io.Writer) {
	fmt.Fprintln(w, p.cutMessage)
}

func (p *basePizza) Box(w io.Writer) {
	fmt.Fprintln(w, "Place pizza in official PizzaStore box")
}

func (p *basePizza) GetName() string {
	return p.name
}

type NYStyleCheesePizza struct {
	basePizza
}

func newNYStyleCheesePizza() Pizza {
	return &NYStyleCheesePizza{
		basePizza: basePizza{
			name:           "NY Style Sauce and Cheese Pizza",
			prepareMessage: "Preparing NY style dough, sauce, and cheese",
			cutMessage:     "Cutting the pizza into diagonal slices",
		},
	}
}

type ChicagoStyleCheesePizza struct {
	basePizza
}

func newChicagoStyleCheesePizza() Pizza {
	return &ChicagoStyleCheesePizza{
		basePizza: basePizza{
			name:           "Chicago Style Deep Dish Cheese Pizza",
			prepareMessage: "Preparing Chicago style dough, sauce, and cheese",
			cutMessage:     "Cutting the pizza into square slices",
		},
	}
}

type PizzaFactory interface {
	CreatePizza(item string) (Pizza, error)
}

type pizzaBuilder func() Pizza

type regionalPizzaFactory struct {
	style string
	menu  map[string]pizzaBuilder
}

func (f *regionalPizzaFactory) CreatePizza(item string) (Pizza, error) {
	builder, ok := f.menu[item]
	if !ok {
		return nil, fmt.Errorf("%s factory: unknown pizza type %q", f.style, item)
	}

	return builder(), nil
}

func NewNYPizzaFactory() PizzaFactory {
	return &regionalPizzaFactory{
		style: "NY",
		menu: map[string]pizzaBuilder{
			CheesePizza: newNYStyleCheesePizza,
		},
	}
}

func NewChicagoPizzaFactory() PizzaFactory {
	return &regionalPizzaFactory{
		style: "Chicago",
		menu: map[string]pizzaBuilder{
			CheesePizza: newChicagoStyleCheesePizza,
		},
	}
}

type PizzaStore struct {
	factory PizzaFactory
	out     io.Writer
}

func NewPizzaStore(factory PizzaFactory, out io.Writer) *PizzaStore {
	if out == nil {
		panic("pizza store requires an output writer")
	}

	return &PizzaStore{factory: factory, out: out}
}

func (ps *PizzaStore) OrderPizza(item string) (Pizza, error) {
	pizza, err := ps.factory.CreatePizza(item)
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(ps.out, "--- Making a %s ---\n", pizza.GetName())
	pizza.Prepare(ps.out)
	pizza.Bake(ps.out)
	pizza.Cut(ps.out)
	pizza.Box(ps.out)

	return pizza, nil
}
