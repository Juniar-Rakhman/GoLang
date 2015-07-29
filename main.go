package main
import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Print func(string) ()

func Greet(salutation Salutation, do func(string)) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "Hi", "yo")
	do(message)
	do(alternate)
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[2] + " "+ name
	alternate = "Hey " + name
	return
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s, Print)
	Greet(s, PrintLine)
}
