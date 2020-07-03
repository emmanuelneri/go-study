package main

import "fmt"

func main() {
	languages := [4]string{"Go", "Java", "Python", "NodeJS"}
	fmt.Println(languages)

	for index, language := range languages {
		fmt.Println(index, "-", language)
	}

	fmt.Println("----------")

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	fmt.Println("----------")

	var emptyArray []int
	fmt.Println(emptyArray)
	fmt.Println(emptyArray == nil)

}
