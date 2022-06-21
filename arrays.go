package main

import "fmt"

func main() {
    var myArray = [3]string("ab","b","c")

	fmt.Println(myArray)

	var myArrayV2 = [...]string("d","e","f","g")

	fmt.Println(myArrayV2)
}