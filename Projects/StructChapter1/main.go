package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
)

func getNoteData() (string, string) {
	title := getUserData("Please type yout title: ")
	content := getUserData("Please type your content: ")
	return title, content
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note successed")

}

func getUserData(promt string) string {
	fmt.Printf("%v ", promt)
	// fmt.Scan(&value) - Scan and Scanln accepts only single word inputs
	// We need to specify from which source we want to read
	// In our case it is a command line, that is why it is os.Stdin
	reader := bufio.NewReader(os.Stdin)
	// Now it want to know at which byte should it stop reading. We say at line break
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	// now we need to remove the line breal \n because it will be included
	// TrimSuffix allows to remove text from the end of the string
	text = strings.TrimSuffix(text, "\n")

	// We also need to remove this random character coz windows creates \r\n instead of \n
	text = strings.TrimSuffix(text, "\r")
	return text
}
