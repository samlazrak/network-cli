package main

import "fmt"

// used to loop through maps, arrays, slices, etc.
func main() {
	ids := []int{
		33,
		34,
		35,
		36,
	}

	// loop through ids with range
	for i, id := range ids {
		fmt.Printf("%d & %d\n", i, id)
	}
	// must use all declared variables, put _ otherwise
	for _, id := range ids {
		fmt.Printf("%d\n", id)
	}

	// sum ids
	sum := 0
	for _, id := range ids {
		// += is add to current number
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with Map - different because key value pairs ( k v)
	emails := map[string]string {
		"Bob": "bob@gmail.com",
		"Sharon": "bob@gmail.com",
		"Waffle": "bob@gmail.com",
	}
	for k, v := range emails {
		// %s is the uninterpretted bytes of the string or slice
		fmt.Printf("%s: %s\n", k, v)
	}
}
