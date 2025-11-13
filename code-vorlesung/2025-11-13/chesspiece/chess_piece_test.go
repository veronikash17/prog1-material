package chesspiece

import "fmt"

func ExampleChessPiece_MoveAllowed() {
	l1 := ChessPiece{
		pieceType: BISHOP,
		colour:    WHITE,
		row:       2,
		column:    2,
	}

	l2 := ChessPiece{
		pieceType: KNIGHT,
		colour:    WHITE,
		row:       2,
		column:    2,
	}

	l3 := ChessPiece{
		pieceType: QUEEN,
		colour:    WHITE,
		row:       2,
		column:    2,
	}

	l4 := ChessPiece{
		pieceType: KING,
		colour:    WHITE,
		row:       2,
		column:    2,
	}
	fmt.Println(l1.MoveAllowed(3, 4))
	fmt.Println(l1.MoveAllowed(0, 0))

	fmt.Println(l2.MoveAllowed(3, 4))
	fmt.Println(l2.MoveAllowed(0, 0))
	fmt.Println(l2.MoveAllowed(2, 0))

	fmt.Println(l3.MoveAllowed(3, 4))
	fmt.Println(l3.MoveAllowed(0, 0))
	fmt.Println(l3.MoveAllowed(2, 0))

	fmt.Println(l4.MoveAllowed(1, 2))
	fmt.Println(l4.MoveAllowed(1, 1))
	fmt.Println(l4.MoveAllowed(1, 2))
	// Output:
	// false
	// true
	// true 
	// false
	// false
	// true
	// true

}
