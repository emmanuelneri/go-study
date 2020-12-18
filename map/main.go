package main

import "fmt"

func main() {

	productMap := map[string]string{"001": "Product A", "002": "Product B"}
	fmt.Println(productMap)

	colorByCode := map[string]string{}
	colorByCode["001"] = "RED"
	colorByCode["002"] = "BLUE"
	delete(colorByCode, "002")
	fmt.Println(colorByCode)

}
