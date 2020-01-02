package main

import "fmt"

func MinimumInt1(first int, rest ...int) (goal int) {
	goal = first
	for _, x := range rest {
		if x < first {
			goal = x
		}
	}
	return
}

func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first
	for _, x := range rest {
		switch x := x.(type) {
		case int:
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum = x
			}
		case string:
			if x < minimum.(string) {
				minimum = x
			}
		}
	}
	return minimum
}

func Index(xs interface{}, x interface{}) int {
	switch slice := xs.(type) {
	case []int:
		for i, y := range slice {
			if y == x.(int) {
				return i
			}
		}
	case []string:
		for i, y := range slice {
			if y == x.(string) {
				return i
			}
		}
	}
	return -1
}

func main() {
	//numbers := []int{7, 6, 2, -1, 7, -3, 9}
	//fmt.Println(Minimum(4, 5, 7, 8))

	xs := []int{2, 4, 6, 8}
	fmt.Println("5 @", Index(xs, 5), " 6 @", Index(xs, 6))
}
