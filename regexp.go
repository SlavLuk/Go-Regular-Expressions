package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var reflections = [][]string{
	{"am", "are"},
	{"was", "were"},
	{"i", "you"},
	{"i'd", "you would"},
	{"i've", "you have"},
	{"i'll", "you will"},
	{"my", "your"},
	{"are", "am"},
	{"you've", "I have"},
	{"you'll", "I will"},
	{"your", "my"},
	{"yours", "mine"},
	{"you", "i"},
	{"me", "you"},
	{"myself", "yourself"},
	{"you're", "i'am"}}

func elizaResponse(input string) string {

	re := regexp.MustCompile(`(?i).*\bfather.*`)
	reg := regexp.MustCompile(`(?i)(i\s?[am']?m)\s?([^.?!]*)[.?!]?`)

	if matched := re.MatchString(input); matched {

		return "Why don’t you tell me more about your father?"

	}

	if matched := reg.MatchString(input); matched {

		converted := reg.ReplaceAllString(input, "$2")

		// Split the input on word boundaries.
		boundaries := regexp.MustCompile(`\b `)
		tokens := boundaries.Split(converted, -1)

		// Loop through each token, reflecting it if there's a match.
		for i, token := range tokens {

			for _, reflection := range reflections {

				matched := strings.Compare(strings.Trim(token, `$(.!)?`), reflection[0])

				if matched == 0 {

					tokens[i] = reflection[1]

				}
			}
		}

		// Put the tokens back together.
		converted = strings.Join(tokens, ` `)

		return "How do you know you are " + converted
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
	fmt.Println("I am supposed to just take what you're saying at face value?")
	fmt.Println(elizaResponse("I am supposed to just take what you're saying at face value?"))

	fmt.Println("------------------------------------------------")
	fmt.Println("I'm happy.")
	fmt.Println(elizaResponse("I'm happy."))
	fmt.Println("Im not happy with your responses.")
	fmt.Println(elizaResponse("Im not happy with your responses."))
	fmt.Println("i'm not sure that you understand the effect that your questions are having on me.")
	fmt.Println(elizaResponse("i'm not sure that you understand the effect that your questions are having on me."))
	fmt.Println("I AM supposed to just take what you're saying at face value?")
	fmt.Println(elizaResponse("I AM supposed to just take what you're saying at face value?"))
	fmt.Println("------------------------------------------------")
	fmt.Println("I am not sure that you understand the effect your questions are having on me.")
	fmt.Println(elizaResponse("I am not sure that you understand the effect your questions are having on me."))

}
