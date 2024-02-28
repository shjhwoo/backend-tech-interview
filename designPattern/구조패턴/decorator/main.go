package main

type IPizza interface {
	getToppings() []string
	addToppings(string)
	getPrice() int
}

type VeggieMania struct {
	Toppings []string
	Price    int
}

func (v *VeggieMania) getToppings() []string {
	return v.Toppings
}

func (v *VeggieMania) addToppings(topping string) {
	v.Toppings = append(v.Toppings, topping)
}

func (v *VeggieMania) getPrice() int {
	return v.Price
}

type TomatoTopping struct {
	IPizza
}

func (t *TomatoTopping) getToppings() []string {
	return t.IPizza.getToppings()
}

func (t *TomatoTopping) addToppings(topping string) {
	t.IPizza.addToppings(topping)
}

func (t *TomatoTopping) getPrice() int {
	return t.IPizza.getPrice() + 50
}

type CheeseTopping struct {
	IPizza
}

func (t *CheeseTopping) getToppings() []string {
	return t.IPizza.getToppings()
}

func (t *CheeseTopping) addToppings(topping string) {
	t.IPizza.addToppings(topping)
}

func (t *CheeseTopping) getPrice() int {
	return t.IPizza.getPrice() + 60
}

func main() {
	veggieMania := &VeggieMania{
		Toppings: []string{"onion", "tomato"},
		Price:    200,
	}

	pizzawithtomatoTopping := &TomatoTopping{veggieMania}
	pizzawithtomatoTopping.addToppings("tomato")
	pizzawithtomatoTopping.getPrice()

	pizzawithTomatoAndcheeseTopping := &CheeseTopping{pizzawithtomatoTopping}
	pizzawithTomatoAndcheeseTopping.addToppings("cheese")
	pizzawithTomatoAndcheeseTopping.getPrice()
}
