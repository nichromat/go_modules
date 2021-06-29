package greetings

import "fmt"

func main(name string) string{
	message := fmt.Sprintf("Hello %v", name)
	return message;
}