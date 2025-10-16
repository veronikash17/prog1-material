package list

import "fmt"

func Find(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}
func ExampleFind() {

	l1 := []int{16, 26, 35, -4, 8, -12}

	pos1 := Find(l1, 26)
	pos2 := Find(l1, 13)

	fmt.Println(pos1)
	fmt.Println(pos2)

	//Output:

}
