package fizzbuzz

import "fmt"

func FizzCompact() {
	for i := 1; i <= 30; i++ {
		//fizzbuss durch 3 und 5 teilbar
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}
		if i%3 == 0 { //fizz - 3
			fmt.Println("fizz")
		}
		if i%5 == 0 { //buss - 5
			fmt.Println("buss") //sonst i
		}
		fmt.Println()

	}

}
