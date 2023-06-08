package greetings

import "fmt"

func Hello (name string) string {
    message := fmt.Sprintf("Go fly, %v!", name)
    return message
}
