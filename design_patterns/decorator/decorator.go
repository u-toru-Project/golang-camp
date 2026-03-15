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

type Mocha struct {
	Beverage
}

func NewMocha(b Beverage) Beverage {
	return Mocha{Beverage: b}
}

func (m Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", Mocha"
}

func (m Mocha) Cost() float64 {
	return m.Beverage.Cost() + 0.20
}

type Whip struct {
	Beverage
}

func NewWhip(b Beverage) Beverage {
	return Whip{Beverage: b}
}

func (w Whip) GetDescription() string {
	return w.Beverage.GetDescription() + ", Whip"
}

func (w Whip) Cost() float64 {
	return w.Beverage.Cost() + 0.10
}

type Soy struct {
	Beverage
}

func NewSoy(b Beverage) Beverage {
	return Soy{Beverage: b}
}

func (s Soy) GetDescription() string {
	return s.Beverage.GetDescription() + ", Soy"
}

func (s Soy) Cost() float64 {
	return s.Beverage.Cost() + 0.15
}
