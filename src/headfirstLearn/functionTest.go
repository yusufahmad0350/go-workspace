package main

import (
	"fmt"
)

func calcTotalPaintNeeded(height float64, width float64) (float64, error) {
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	if width < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", width)
	}

	area := height * width

	return area / 10, nil

}

func main() {
	var total float64

	Neededpaint1, err := calcTotalPaintNeeded(6.2, 4.5)
	fmt.Printf(" total Lit Needed to paint is %0.2f\n", Neededpaint1)

	total += Neededpaint1

	Neededpaint2, err := calcTotalPaintNeeded(-3.4, 5.6)
	//fmt.Print("The error message is", err)
	//total += Neededpaint2
	//fmt.Printf("The total paint  for both walls are %0.2f\n", total)

	//fmt.Printf("%12s | %s\n", "Total Area", "Cost in total")
	//fmt.Println("---------------------------------")
	fmt.Printf("%12s | %0.2f\n", "Second Room", Neededpaint2)
	//fmt.Printf("%12s | %5.2f\n", "First Room", Neededpaint1)

	//fmt.Printf("the string os %0.2f\n", Neededpaint2)
	//fmt.Printf("%f The value is \n", Neededpaint2)

	//error
	//err = errors.New("Height cannot be negative")
	fmt.Printf("The error message is %s\n", err)
	a, b, c := multiReturnFunct()
	fmt.Printf("The multi return function returns 3 type %v %v %vof values \n", a, b, c)

}

func multiReturnFunct() (int, bool, string) {
	return 1, false, "Yusuf"

}
