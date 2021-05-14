package main

import (
	"fmt"
	"math"
	"net/http"
)

// adds two values
func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a/b, a%b
}

func double_by_val(n int) {
	n *= 2
}

// missing out on not allow to say it's a ref with &
// but I guess then you are never asking, is this really a ptr?
func double_by_ref(n *int) {
	*n *= 2
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
	}

	return math.Sqrt(n), nil
}

func cleanup(name string) {
	fmt.Println(name)
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("working.")
}

func contentType(url string) (string, error) {
	
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	return ct, err
}

func main() {
	val := add(3,3)
	fmt.Println(val)

	div, mod := divmod(7,2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)

	val = 10
	double_by_val(val)
	fmt.Println(val)

	double_by_ref(&val)
	fmt.Println(val)

	// shows error
	s1, err := sqrt(-2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(s1)
	}
	s2, err := sqrt(2.0)
	fmt.Print(s2)
	
	worker()

	ct, err := contentType("https://mikelynchgames.com")
	fmt.Println(ct)

}

