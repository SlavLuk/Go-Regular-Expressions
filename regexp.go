package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func elizaResponse(input string) string {

	re := regexp.MustCompile(`(?i).*\bfather.*`)
	reg := regexp.MustCompile(`(?i)(i\s?[am']?m)\s?([^.?!]*)[.?!]?`)

	if matched := re.MatchString(input); matched {

		return "Why don’t you tell me more about your father?"

	} else if matched := reg.MatchString(input); matched {

		replace := reg.ReplaceAllString(input, "How do you know you are $2?")

		return replace
	}

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
	fmt.Println("------------------------------------------------")
	fmt.Println("I am happy.")
	fmt.Println(elizaResponse("I am happy."))
	fmt.Println("I am not happy with your responses.")
	fmt.Println(elizaResponse("I am not happy with your responses."))
	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(elizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(elizaResponse("I am supposed to just take what you’re saying at face value?"))

	fmt.Println("------------------------------------------------")
	fmt.Println("I'm so happy")
	fmt.Println(elizaResponse("I'm so happy"))
	fmt.Println("Im so happy")
	fmt.Println(elizaResponse("Im so happy"))
}
