// http://blog.denevell.org/golang-closures-anonymous-functions.html
package main

import "fmt"

func main() {
	anon := func(name string) string {
		return "Hiya, " + name
	}
	anotherFunction(anon)
}

func anotherFunction(f func(string) string) {
	result := f("David")
	fmt.Println(result) // Prints "Hiya, David"
}
