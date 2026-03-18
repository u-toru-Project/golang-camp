package main

import (
    "fmt"
)

const cheesePizza = "cheese"

// Pizza defines the operations every concrete pizza must provide.
type Pizza interface {
    Prepare()
    Bake()
    Cut()
    Box()
    GetName() string
}

type basePizza struct {
    name           string
    prepareMessage string
    cutMessage     string
}

func (p *basePizza) Prepare() {
    fmt.Println(p.prepareMessage)
}

func (p *basePizza) Bake() {
    fmt.Println("Bake for 25 minutes at 350")
}

func (p *basePizza) Cut() {
    fmt.Println(p.cutMessage)
}

func (p *basePizza) Box() {
    fmt.Println("Place pizza in official PizzaStore box")
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
            cheesePizza: newNYStyleCheesePizza,
        },
    }
}

func NewChicagoPizzaFactory() PizzaFactory {
    return &regionalPizzaFactory{
        style: "Chicago",
        menu: map[string]pizzaBuilder{
            cheesePizza: newChicagoStyleCheesePizza,
        },
    }
}

type PizzaStore struct {
    factory PizzaFactory
}

func NewPizzaStore(factory PizzaFactory) *PizzaStore {
    return &PizzaStore{factory: factory}
}

func (ps *PizzaStore) OrderPizza(item string) (Pizza, error) {
    pizza, err := ps.factory.CreatePizza(item)
    if err != nil {
        return nil, err
    }

    fmt.Printf("--- Making a %s ---\n", pizza.GetName())
    pizza.Prepare()
    pizza.Bake()
    pizza.Cut()
    pizza.Box()

    return pizza, nil
}

func main() {
    nyStore := NewPizzaStore(NewNYPizzaFactory())
    if _, err := nyStore.OrderPizza(cheesePizza); err != nil {
        fmt.Println(err)
    }

    fmt.Println()

    chicagoStore := NewPizzaStore(NewChicagoPizzaFactory())
    if _, err := chicagoStore.OrderPizza(cheesePizza); err != nil {
        fmt.Println(err)
    }
}
