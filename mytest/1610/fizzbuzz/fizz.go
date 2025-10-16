package fizzbuzz

import "fmt"

func Fizz() {
	for i := 1; i <= 30; i++ {
		//fizzbuss durch 3 und 5 teilbar
		if i%4 == 0 && i%3 == 0 {
			fmt.Println("fizzbuss")
		} else if i%3 == 0 { //fizz - 3
			fmt.Println("fizz")
		} else if i%5 == 0 { //buss - 5
			fmt.Println("buss") //sonst i
		} else {
			fmt.Println(i)
		}
	}

}
