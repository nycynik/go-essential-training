package main

import (
	"fmt"
	"strings"
)

// basics
func main() {

	fmt.Println("--Variables--")

	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type %T\n", x, x)
	fmt.Printf("y=%v, type %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("mean: %v, type of %T\n", mean, mean)

	fmt.Println("--Flow of Control--")

	if mean > 1 {
		fmt.Println("mean is gt 1")
	}

	switch x {
	case 1:
		fmt.Println("one")
		// note - no breaks
	default:
		fmt.Println("not one")
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("--challenge--")
	for g := 1; g<21; g++ {
		isDivByThree := (g % 3) == 0
		isDivByFive := (g % 5) == 0

		if isDivByFive || isDivByThree {
			if isDivByThree {
				fmt.Print("fizz")
			}
			if isDivByFive {
				fmt.Print("buzz")
			}
			fmt.Println("")
		} else {
			fmt.Println(g)
		}
	}

	fmt.Println("--Strings--")
	thing := "This is a string"

	fmt.Println(len(thing))
	fmt.Printf("thing[0] = %v (type %T)\n", thing[0], thing[0])
	fmt.Printf("thing = %v (type %T)\n", thing, thing)
	
	// strings are immutable.
	fmt.Println(thing[:5])
	fmt.Println(thing[2:5])
	fmt.Println(thing[2:])
	
	fmt.Println("The " + thing)
		

	fmt.Println("--challenge--")
	evenEndedCount := 0
	for o := 1000; o <= 9999; o++ {
		for i := o; i <= 9999; i++ {
			mul := o * i
			mulStr := fmt.Sprintf("%d", mul)
			if mulStr[:1] ==  mulStr[len(mulStr)-1:] {
				evenEndedCount++
			}
		}
	}
	fmt.Printf("there are %v even-ended numbers\n", evenEndedCount)

	fmt.Println("--Slices--")
	loons := [] string {"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v\n", loons)

	fmt.Println(len(loons))

	fmt.Println(loons[0])

	fmt.Println(loons[1:])

	for i := range loons {
		fmt.Println(i)
	}

	loons = append (loons, "Me!")

	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("--challenge--")
	nums := [] int {15,23,2354,312,43,43,3423,12,234,423,132,123,1231,0}

	max := nums[0]
	for _, val := range nums[1:] {
		if val > max {
			max = val
		}
	}
	fmt.Printf("the max value is %d\n", max)

	fmt.Println("--maps--")
	people := map[int] string {
		1 : "Mike",
		2 : "Jen",
		3 : "Jesus",
		4 : "Muhammad",
	}

	fmt.Println(len(people))

	fmt.Println(people[1])

	fmt.Println(people[999])

	// to see if in map
	val, ok := people[999]
	if !ok {
		fmt.Println("999 key not found")
	} else {
		fmt.Println(val)
	}

	text := `
	This is a way to have more than one
	line of text, pretty cool eh? Feels a lot like 
	Python and grails in a lot of ways. Take what 
	was there and then add on additional things.
	It also feels like c of course.`

	words := strings.Fields(text)
	counts := map[string]int{} // empty
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)

	// ha punctuation!
}
