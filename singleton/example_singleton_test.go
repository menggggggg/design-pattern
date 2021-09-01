package singleton

import "fmt"

func ExampleGetSingleInstance() {
	single := GetSingleInstance()
	fmt.Println(single.message())
	// Output:
	// singleton object!
}
