package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	genword "github.com/dustinkirkland/golang-petname"
)

func jumbled(word string) string {
	var jumble string
	for len(word) > 0 {
		i := rand.Intn(len(word))
		// fmt.Println("i", i)
		jumble += string(word[i])
		// fmt.Println("jumble", jumble)
		word = word[:i] + word[i+1:]
		// fmt.Println("WORD", word)
	}
	return jumble
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	var name string
	fmt.Print("Please Insert Your Name: ")
	fmt.Scanln(&name)
	// fmt.Println("LEN", len(name))
	genNum := len(name)
	// fmt.Println("Num", genNum)
	get := genword.Generate(genNum, ",")
	// fmt.Println("GET", get)
	tes := strings.Split(get, ",")

	Words := tes
	fmt.Printf(
		`Welcome to Scramble Words!
		Find The Right Word.
		`)

	points := 0
	input := bufio.NewScanner(os.Stdin)

	for len(Words) > 0 {
		randomNum := rand.Intn(len(Words))
		// fmt.Println("randomNum", randomNum)
		randomWord := Words[randomNum]
		// fmt.Println("randomword", randomWord)
		jumbledWord := jumbled(randomWord)

		fmt.Println("\nThe jumbled word is: ", jumbledWord)
		fmt.Println("Your guess is:", string(randomWord[0]), "....")
		fmt.Println("Now your score is :", points)
		var guess string
		if input.Scan() {
			guess = input.Text()
		}
		if guess == randomWord {
			fmt.Println("Thats Correct. Excellent.!!")
			points += 100
			fmt.Println("Your score:", points)
		} else {
			fmt.Println("Thats not correct")
			points -= 100
			fmt.Println("Correct word is:", randomWord)
		}
		if len(Words) > 1 {
			firstWords := Words[:randomNum]
			fmt.Printf("%d Words remaining \n\n\n", len(Words)-1)
			lastWords := Words[randomNum+1:] //Remove this word from the Word list
			Words = append(firstWords, lastWords...)
		} else {
			break
		}
		fmt.Println("Do you want to play again? (y/n)")
		if input.Scan() {
			play := input.Text()
			if play != "y" {
				break
			}
		}

	}
	fmt.Println("\nThanks for playing.")
	fmt.Println("\nYour SCORE:", points)
	fmt.Println("Congratulations..!!!..See you again.")
}
