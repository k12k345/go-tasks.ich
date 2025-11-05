package listprops

import "fmt"

func ExampleContainsValue() {
	list := []int{1, 3, 1, 4, 1, 5, 2, 5, 3, 5}
	emptyList := []int{}

	fmt.Println(ContainsValue(list, 1))
	fmt.Println(ContainsValue(list, 5))
	fmt.Println(ContainsValue(list, 4))
	fmt.Println(ContainsValue(list, 3))

	fmt.Println()

	fmt.Println(ContainsValue(emptyList, 1))
	fmt.Println(ContainsValue(emptyList, 5))
	fmt.Println(ContainsValue(emptyList, 4))
	fmt.Println(ContainsValue(emptyList, 3))

	// Output:
	// true
	// true
	// true
	// true
	//
	// false
	// false
	// false
	// false
}
