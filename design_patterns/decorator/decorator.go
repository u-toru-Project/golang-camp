package decorator

type Beverage interface {
	GetDescription() string
	Cost() float64
}

type Espresso struct{}

func (Espresso) GetDescription() string {
	return "Espresso"
}

func (Espresso) Cost() float64 {
	return 1.99
}

func NewEspresso() Beverage {
	return Espresso{}
}

type HouseBlend struct{}

func (HouseBlend) GetDescription() string {
	return "House Blend Coffee"
}

func (HouseBlend) Cost() float64 {
	return 0.89
}

func NewHouseBlend() Beverage {
	return HouseBlend{}
}

type topping struct {
	Beverage
	label   string
	addCost float64
}

func newTopping(b Beverage, label string, addCost float64) Beverage {
	return topping{Beverage: b, label: label, addCost: addCost}
}

func (t topping) GetDescription() string {
	return t.Beverage.GetDescription() + ", " + t.label
}

func (t topping) Cost() float64 {
	return t.Beverage.Cost() + t.addCost
}

func NewMocha(b Beverage) Beverage {
	return newTopping(b, "Mocha", 0.20)
}

func NewWhip(b Beverage) Beverage {
	return newTopping(b, "Whip", 0.10)
}

func NewSoy(b Beverage) Beverage {
	return newTopping(b, "Soy", 0.15)
}
