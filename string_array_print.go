package main

import "fmt"

func main() {

	// string
		// var nameOne string = "hello you fool"
		// var nameTwo = "yet another string"
		// var nameThree string

		// fmt.Println(nameOne, nameTwo, nameThree)
		
		// nameOne = "peach"
		// nameThree = "bowser"
		
		// fmt.Println(nameOne, nameTwo, nameThree)
		
		// short hand for declaring variable (doesn't work outside of main function)
		// nameFour := "yoshi"

		// fmt.Println(nameFour)

	// ints
		// var ageOne int = 20
		// var ageTwo = 30
		// ageThree := 40

		// fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
		// var numOne int8 = 25
		// var numTwo int8 = -128
		// var numThree uint8 = 255

		// fmt.Println(numOne, numTwo, numThree)

	// float

		// var scoreOne float32 = 25.98
		// var scoreTwo float64 = 234354548767829849238498239.7
		// scoreThree := 1.3

		// fmt.Println(scoreOne, scoreTwo, scoreThree)

	// Print (no new line)

		// fmt.Print("hello, ")
			// with new line:
		// fmt.Print("world! \n")
		// fmt.Print("new line! \n")

		// fmt.Println("hello fools!")
		// fmt.Println("goodbye fools!")

		// age := 40
		// name:= "aaron"
		// yes := false

		// fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier (v = variable, %T = type, %0.1f = float )

		// fmt.Printf("my age is %v and my name is %v \n", age, name)
		// fmt.Printf("my age is %q and my name is %q \n", age, name)
		// fmt.Printf("age is of type %T \n", age)
		// fmt.Printf("you scored %0.1f points! \n", 224.3)
	
	// Sprintf (save formatted strings)
		// var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)

		// fmt.Println("the saved string is:", str)
		// fmt.Printf("the world is round is %t \n", yes)

	// arrays
		// var ages [3]int = [3]int{20, 25, 30}
		// var ages = [3]int{20, 25, 30}

		// fmt.Println(ages, len(ages))

		names := [4]string{"john", "ringo", "george", "paul"}
		names[1] = "Desmond"

		// fmt.Println(names, len(names))

	//  slices (use arrays under the hood)

		// var scores = []int{100, 50, 60}
		// scores[2] = 25
		// scores = append(scores, 85)

		// fmt.Println(scores, len(scores))

	// slice ranges

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
