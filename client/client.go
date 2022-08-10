package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"example.com/server"
)

func main() {

	log.SetPrefix("Server: ")
	log.SetFlags(0)

	fmt.Println("Welcome to GoLang")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your First Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("-------------------------")

	message, err := server.Welcome(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
