package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting, "Hi", "yo")
	if prefix := "Mr "; isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {

	switch name {
		case "bob" : prefix = "mr "
		case "marry" : prefix = "dr "
		case "joe" : prefix = "sir "
		default: prefix = "Dude "
	}

	return
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

//...means we are passing unknown number of var
func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[2] + " "+ name
	alternate = "Hey " + name
	return
}
