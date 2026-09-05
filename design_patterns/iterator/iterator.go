package iterator

import (
	"fmt"
	"io"
)

type MenuItem struct {
	Name        string
	Description string
	Vegetarian  bool
	Price       float64
}

func newMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		Name:        name,
		Description: description,
		Vegetarian:  vegetarian,
		Price:       price,
	}
}

type Iterator interface {
	HasNext() bool
	Next() (*MenuItem, bool)
}

type Menu interface {
	CreateIterator() Iterator
}

type menuItemIterator struct {
	items     []*MenuItem
	nextIndex int
}

func newMenuItemIterator(items []*MenuItem) Iterator {
	return &menuItemIterator{items: items}
}

func (it *menuItemIterator) HasNext() bool {
	return it.nextIndex < len(it.items)
}

func (it *menuItemIterator) Next() (*MenuItem, bool) {
	if !it.HasNext() {
		return nil, false
	}

	item := it.items[it.nextIndex]
	it.nextIndex++

	return item, true
}

type PancakeHouseMenu struct {
	items []*MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	return &PancakeHouseMenu{
		items: []*MenuItem{
			newMenuItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs and toast", true, 2.99),
			newMenuItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49),
		},
	}
}

func (m *PancakeHouseMenu) CreateIterator() Iterator {
	return newMenuItemIterator(m.items)
}

const maxMenuItems = 6

type DinerMenu struct {
	items         [maxMenuItems]*MenuItem
	numberOfItems int
}

func NewDinerMenu() *DinerMenu {
	menu := &DinerMenu{}
	menu.addItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	menu.addItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)

	return menu
}

func (m *DinerMenu) addItem(name, description string, vegetarian bool, price float64) {
	if m.numberOfItems >= maxMenuItems {
		return
	}

	m.items[m.numberOfItems] = newMenuItem(name, description, vegetarian, price)
	m.numberOfItems++
}

func (m *DinerMenu) CreateIterator() Iterator {
	return newMenuItemIterator(m.items[:m.numberOfItems])
}

type Waitress struct {
	breakfastMenu Menu
	lunchMenu     Menu
	out           io.Writer
}

func NewWaitress(breakfastMenu, lunchMenu Menu, out io.Writer) *Waitress {
	if out == nil {
		panic("waitress requires an output writer")
	}

	return &Waitress{
		breakfastMenu: breakfastMenu,
		lunchMenu:     lunchMenu,
		out:           out,
	}
}

func (w *Waitress) PrintMenu() {
	fmt.Fprintln(w.out, "MENU\n----\nBREAKFAST")
	w.printSection(w.breakfastMenu.CreateIterator())

	fmt.Fprintln(w.out, "\nLUNCH")
	w.printSection(w.lunchMenu.CreateIterator())
}

func (w *Waitress) printSection(iterator Iterator) {
	for {
		item, ok := iterator.Next()
		if !ok {
			return
		}

		fmt.Fprintf(w.out, "%s, $%.2f -- %s\n", item.Name, item.Price, item.Description)
	}
}
