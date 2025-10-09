package main

import (

	"fmt"
	"math/rand"
)

func main() {
	//my_number = 17
	for n:= 0, n<3, n = n + 1 {
       guess:= ReadNumber()
	   if NumberGood(guess) {
		fmt.Println("Das war richtig!")
		return
	   }
	   fmt.Printf("%d war nicht korrekt",guess)

	}
	fmt.Println("Game Ower")
}

func ReadNumber() int {
	var n int 
	fmt.Print("Rate eine Zahl")
	fmt.Scan(&n)
	return n
}

 func NumberGood(x int) bool {
my_number := rand.Intn(100)+1
if x == my_number{
	return true
}
return true
 }
