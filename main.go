package main
import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)

	fmt.Println(message)
	fmt.Println(alternate)
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " "+ name
	alternate = "Hey " + name
	return
}


func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
