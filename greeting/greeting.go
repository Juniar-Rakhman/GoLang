package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation Salutation, print Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting, "Hi", "yo")
	if prefix := GetPrefix(salutation.Name); isFormal {
		print(prefix + message)
	} else {
		print(alternate)
	}
}

func GetPrefix(name string) (prefix string) {

	switch name {
		case "Bob" :
			prefix = "mr "
		case "Marry" :
			prefix = "dr "
		case "Joe","Amy" :
			prefix = "sir "
		default:
			prefix = "Dude "
	}

	return
}

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("saluation")
	default:
		fmt.Println("uknown")
	}
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

//...means we are passing unknown number of var
func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[2] + " " + name
	alternate = "Hey " + name
	return
}
