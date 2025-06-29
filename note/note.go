package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// defining
type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_at"`
}

// constructor function
func CreateNote(title, desc string) (Note, error) {
	if title == "" || desc == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:       title,
		Description: desc,
		CreatedOn:   time.Now(),
	}, nil
}

// struct method - receiver param
func (n Note) PrintNote() {
	fmt.Printf("Note title: %v\n", n.Title)
	fmt.Printf("Note Description: %v\n", n.Description)
	fmt.Printf("Note Created On: %v\n", n.CreatedOn)
}

func (n Note) SaveNote() {
	fileName := strings.ToLower(strings.ReplaceAll(n.Title, " ", "_")) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		fmt.Print(err)
	}

	os.WriteFile(fileName, json, 0644)
}
