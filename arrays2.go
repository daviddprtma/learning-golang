package main

import "fmt"

func main() {
   var letters = [...]string("a","b","c","d","e")

   letters[0] = "f"
   fmt.Println(letters)

   fmt.Println(letters[1:4])
}