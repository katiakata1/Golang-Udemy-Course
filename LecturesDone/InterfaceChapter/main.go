package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
	"example.com/go-project/todo"
)

type saver interface {
	Save() error
}

// embedded interface
type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(5)
	printSomething(6.78)
	printSomething("Hello")
	title, content := getNoteData()
	todoText := getUserData("To do text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

// any type of value using another approach
func printSomething(value any) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
	}
}

// embedded interfaces
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note successed")
	return nil
}

func getNoteData() (string, string) {
	title := getUserData("Please type yout title: ")
	content := getUserData("Please type your content: ")
	return title, content
}

func getUserData(promt string) string {
	fmt.Printf("%v ", promt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
