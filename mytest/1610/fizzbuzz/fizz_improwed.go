package fizzbuzz

import "fmt"

func FizzImprowed(x, y, n int) {
	for i := 1; i <= n; i++ {
		//fizzbuss durch 3 und 5 teilbar
		if i%y == 0 && i%x == 0 {
			fmt.Println("fizzbuss")
		} else if i%x == 0 { //fizz - 3
			fmt.Println("fizz")
		} else if i%y == 0 {//buss - 5
			fmt.Println("buss")//sonst i
		}else {
			fmt.Println(i)
		}
	}

}