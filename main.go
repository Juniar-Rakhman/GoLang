package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Test", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction(" !!!"), true)
	greeting.TypeSwitchTest(s)
}

