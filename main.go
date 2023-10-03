package main

import (
	"fmt"
	"log"
)

func main() {
	var selectedTemplate int
	fmt.Println("Please Select your template:")
	fmt.Printf("1- %s\n2- %s\n", "HTML and standard CSS", "HTML Boilerplate")
	_, err := fmt.Scanln(&selectedTemplate)
	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}
}
