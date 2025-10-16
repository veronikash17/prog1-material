package list

import "fmt"

func Example_liste() {
	l1 := []int{1,3,6,24,54}
	fmt.Println(l1[3])
	fmt.Println(l1[0])
	l1[3] = 52
    fmt.Println(l1[3])
	for i := 0; i < len(l1); i++{
		fmt.Println(l1[1])
	}




	//Output:
}
