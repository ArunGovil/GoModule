package server

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Welcome(name string) (string, error) {

	if name == "" {
		return "", errors.New("Name Empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func WelcomeMore(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Welcome(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v, Welcome to GoLang!",
		"Great to see you, %v, Enjoy Coding!",
		"Hey there %v, Get ready to Go!",
	}
	return formats[rand.Intn(len(formats))]
}
