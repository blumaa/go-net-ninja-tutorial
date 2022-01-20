package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main() {
	// sayGreeting("john")
	// sayGreeting("paul")
	// sayBye("ringo")
	
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

  //	a1 := circleArea(10.5)
  //	a2 := circleArea(15)

  //	fmt.Println(a1, a2)
  //	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
	
	
}
