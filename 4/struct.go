package main

import "fmt"

/* Trade */
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if (t.Buy) {
		value = -value
	}

	return value
}

/* main */
func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)
	fmt.Println(t1.Value())

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
}