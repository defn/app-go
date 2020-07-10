package main

import "fmt"

func GreetSomeone(name string) {
	fmt.Println("Hello,", name)
}

func plus(a int, b int) int {
  return a
}

func main() {
	names := [4]string{"lamda", "Hana", "Tolan", "defn"}

	for i := 0; i < 4; i++ {
		GreetSomeone(names[i])
	}

  sum := plus(plus(1,2), plus(3,4))
  fmt.Println("Sum is", sum)
}

