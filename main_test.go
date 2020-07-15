package main

import "testing"

func TestGreeter_GetGreetingForHahahaha(t *testing.T) {
	g := &Greeter{WelcomeText: "hey, "}
	if ret := g.GetGreetingForColleague("Matti"); ret != "hey, Matti" {
		t.Fatal("unexpected greeting text: " + ret)
	}
}
