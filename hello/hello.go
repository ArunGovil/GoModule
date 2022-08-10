package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("First Name: ")
	fname, _ := reader.ReadString('\n')
	fmt.Print("Last Name: ")
	lname, _ := reader.ReadString('\n')
	name := fname + lname

	// names := []string{"Rock", "Paper", "Scissor"}

	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
