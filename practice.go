package main

import (
	"fmt"
)

func main(){

	fmt.Println("\nHello, this is for practice.\n")

	a()
	tictactoe()


}

// ***************************************************************************************************
// Practicing concepts

func a(){
	fmt.Println("FUNCTION A: TYPES \n")
	a := 1
	b := 1.23
	c := 0.123
	d := "string"
	e := false

	fmt.Printf("a equals: %d and is of type %T\t\n",a,a)
	fmt.Printf("b equals: %v and is of type %T\t\n",b,b)
	fmt.Printf("c equals: %v and is of type %T\t\n",c,c)
	fmt.Printf("d equals: %v and is of type %T\t\n",d,d)
	fmt.Printf("e equals: %v and is of type %T\t\n",e,e)

}

func tictactoe() {

	fmt.Println("Lets fucking play some tic tac toe")

	//Game lobby lol
	p1 := "Player 1"
	p2 := "Player 2"
	fmt.Println("Enter the name of the first player:")
	fmt.Scanln(&p1)
	fmt.Println("Enter the name of the second player:")
	fmt.Scanln(&p2)
	fmt.Printf("Thanks %v and %v! Lets play some Tic Tac Toe\n", p1, p2)

	//positions on the board
	t1, t2, t3, m1, m2, m3, b1, b2, b3 := "-", "-", "-", "-", "-", "-", "-", "-", "-"


	//Empty board visual
	fmt.Print("    1   2   3\n")
	fmt.Printf("t|  %v | %v | %v\n", t1, t2, t3)
	fmt.Print("  -----------\n")
	fmt.Printf("m|  %v | %v | %v\n", m1, m2, m3)
	fmt.Print("  -----------\n")
	fmt.Printf("b|  %v | %v | %v\n", b1, b2, b3)

	//Playing the game
	for i := 0; i < 20; i++ {

		if 0 == i%2 {

			t1, t2, t3, m1, m2, m3, b1, b2, b3, p1 = makemove(t1 , t2, t3, m1, m2, m3, b1, b2, b3, p1, "o")

			} else {
			t1, t2, t3, m1, m2, m3, b1, b2, b3, p2 = makemove(t1 , t2, t3, m1, m2, m3, b1, b2, b3, p2, "x")
		}

		//Looping board per turn
		fmt.Print("    1   2   3\n")
		fmt.Printf("t|  %v | %v | %v\n", t1, t2, t3)
		fmt.Print("  -----------\n")
		fmt.Printf("m|  %v | %v | %v\n", m1, m2, m3)
		fmt.Print("  -----------\n")
		fmt.Printf("b|  %v | %v | %v\n", b1, b2, b3)


		}

	}


func makemove(t1 string, t2 string, t3 string, m1 string, m2 string, m3 string, b1 string, b2 string, b3 string, p1 string, xo string) (string, string, string, string, string, string, string, string, string, string) {

	fmt.Printf("Your move %v!\n", p1)
	move := ""
	fmt.Scanln(&move)

	switch move {
	case "t1":
		if t1 != "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			t1 = xo
		}
	case "t2":
		if t2 != "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			t2 = xo
		}
	case "t3":
		if t3!= "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			t3 = xo
		}
	case "m1":
		if m1 != "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			m1 = xo
		}
	case "m2":
		if m2 != "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			m2 = xo
		}
	case "m3":
		if m3 != "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			m3 = xo
		}
	case "b1":
		if b1 != "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			b1 = xo
		}
	case "b2":
		if b2 != "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			b2 = xo
		}
	case "b3":
		if b3 != "-" {
			fmt.Println("Nice try cheater")
			makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
		} else {
			b3 = xo
		}
	default:
		fmt.Println("You didn't make a move dingus")
		makemove(t1, t2, t3, m1, m2, m3, b1, b2, b3, p1, xo)
	}
	return t1, t2, t3, m1, m2, m3, b1, b2, b3, p1
}

