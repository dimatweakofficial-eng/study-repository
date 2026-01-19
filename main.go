package main

import (
	"demo2/feature1"
	"demo2/feature2"
	newanimal "demo2/newAnimal"
	newanimal2 "demo2/newAnimal2"
	"fmt"
)

func main() {
	fmt.Println("Передаю руль коту")
	feature1.Cat()
	feature2.Dog()
	newanimal.Enot()
	newanimal2.Vidra("Ахаахахха ну меня то тут не ожидали, я выдра бля")
}
