package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func getFileName() string {
	var filename string

	if len(os.Args) == 2 {
		// User passed filename as param
		filename = os.Args[1]
	} else {
		// Prompting user for a filename
		fmt.Print("Enter a filename:")
		_, err := fmt.Scanln(&filename)

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(0)
		}
	}
	return filename
}

func writeFile(filename string) string {
	filepath := "/Users/jonashansen/SecondBrain/00. Inbox/" + filename + ".md"

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	title := fmt.Sprintf("# %s", filename)
	id := fmt.Sprintf("id: %s", time.Now().Format("20060102150405"))
	datetime := fmt.Sprintf("created_at: %s", time.Now().Format("2006-01-02 15:04:05"))
	lines := []string{
		"---",
		datetime,
		id,
		"tags:",
		"---",
		title,
		"",
		"",
		"## Links",
	}

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")

		if err != nil {
			fmt.Println("Error writing to file: ", err)
			os.Exit(0)
		}
	}

	err = writer.Flush()

	if err != nil {
		fmt.Println("Error flushing writer: ", err)
		os.Exit(0)
	}
	return filepath
}

func openFile(filepath string) {
	cmd := exec.Command("nvim", "+7", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	error := cmd.Run()
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(0)
	}
}

func main() {
	var filename string = getFileName()

	if filename == "" {
		fmt.Println("No filename found!")
		os.Exit(0)
	}

	filepath := writeFile(filename)
	openFile(filepath)

	fmt.Println("Note successfully saved!")
}
