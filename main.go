package main

import (
	"fmt"
	"log"
)

type CanSayGreeting interface {
	GetGreetingForColleague(name string) string
}

type Greeter struct {
	WelcomeText string
}

func (greeter *Greeter) GetGreetingForColleague(name string) string {
	return greeter.WelcomeText + name
}

func PrintGreeting(greeterImpl CanSayGreeting, name string) {
	fmt.Println(greeterImpl.GetGreetingForColleague(name))
}

func LogGreeting(greeterImpl CanSayGreeting, name string) {
	log.Print(greeterImpl.GetGreetingForColleague(name))
}

func main() {
	greeter := &Greeter{WelcomeText: "Hi there "}
	for i, personName := range []string{"Jarkko", "SweeMeng", "Szymon"} {
		fmt.Printf("Index number is %d\n", i)
		PrintGreeting(greeter, personName)
		LogGreeting(greeter, personName)
	}
}
