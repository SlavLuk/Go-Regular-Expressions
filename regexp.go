package main

import (
	"fmt"
	"math/rand"
	"time"
)

func elizaResponse(input string) string {

	responses := []string{"I’m not sure what you’re trying to say. Could you explain it to me?", "How does that make you feel?", "Why do you say that?"}

	return responses[rand.Intn(len(responses))]

}

func main() {

	// Seed the rand package with the current time.
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//print outputs
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(elizaResponse("People say I look like both my mother and father."))
	fmt.Println("Father was a teacher.")
	fmt.Println(elizaResponse("Father was a teacher."))
	fmt.Println("I was my father’s favourite.")
	fmt.Println(elizaResponse("I was my father’s favourite."))
	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(elizaResponse("I’m looking forward to the weekend."))
	fmt.Println("My grandfather was French!")
	fmt.Println(elizaResponse("My grandfather was French!"))
}
