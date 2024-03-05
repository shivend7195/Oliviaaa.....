package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomElement(slice []string) string {
	length := len(slice)
	index := rand.Intn(length)
	return slice[index]
}

func main() {

	fmt.Println("Wait! where is my cat?")

	fmt.Println("Olivia....Olivia....Olivia")

	guests := [5]string{"gendeya", "mostlySane", "taylorSwift", "ryanGoslin", "johnWick"}
	objects := [6]string{"box", "basket", "cardign", "handbag", "foodTrolly", "cupboard"}

	fmt.Println("EVERYBODY will stay where they are.")

	fmt.Println("These are guests here", guests)
	fmt.Println("Here possibly we can find Olivia", objects)

	// Setting up the random seed.
	rand.Seed(time.Now().UnixNano())

	// Randomly selecting a guilty guest and a hiding place for the cat for surprise checking.
	guilty := getRandomElement(guests[:])
	place := getRandomElement(objects[:])

	if guilty == "taylorSwift" && place == "cardign" {
		fmt.Println("We knew it", guilty, "hid the cat in her old", place)
	} else {
		fmt.Println(guilty, "hid the cat in", place)
	}
}
