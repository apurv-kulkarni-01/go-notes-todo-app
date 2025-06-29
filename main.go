package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note"
)

func main() {
	title, description := makeNote()
	fmt.Printf("%v \n%v \n", title, description)
	fmt.Println("-----------------------------")
	//Declaring myNote
	myNote, err := note.CreateNote(title, description)

	if err != nil {
		fmt.Println(err)
	}

	myNote.PrintNote()
	myNote.SaveNote()
}

func makeNote() (title, description string) {
	title = getUserInput("Enter Title: ")
	description = getUserInput("Enter Description: ")
	return title, description
}

func getUserInput(s string) string {
	fmt.Print(s)
	// var input string
	// fmt.Scanln(&input)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
