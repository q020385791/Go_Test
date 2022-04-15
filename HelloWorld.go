package main

import "fmt"

func HelloWorld(name string) string {
	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message

}
