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

func createFolder() {
  date := time.Now().Format("2006-01-02")
  folderpath := fmt.Sprintf("/Users/jonashansen/SecondBrain/04. Meetings/%s", date)
  if _, err := os.Stat(folderpath); os.IsNotExist(err) {
    fmt.Println("Creating folder: ", folderpath)
    err := os.Mkdir(folderpath, 0755)
    if err != nil {
      fmt.Println("Error creating folder: ", err)
      os.Exit(0)
    }
  }
}

func writeFile(filename string) string {
  date := time.Now().Format("2006-01-02")
  full_filename := date + " - " + filename
  filepath := fmt.Sprintf("/Users/jonashansen/SecondBrain/04. Meetings/%s/%s.md", date, full_filename)
  fmt.Println(filepath)

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	id := fmt.Sprintf("id: %s", time.Now().Format("20060102150405"))
	datetime := fmt.Sprintf("created_at: %s", time.Now().Format("2006-01-02 15:04:05"))
	title := fmt.Sprintf("# %s - %s", time.Now().Format("2006-01-02"), filename)
	lines := []string{
		"---",
		datetime,
		id,
		"tags:",
		"- meeting",
		"---",
		title,
		"",
		"## Agenda",
		"",
		"",
		"## Attendees",
		"",
		"",
		"## Notes",
		"",
		"",
		"## Action Items",
		"",
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
	cmd := exec.Command("nvim", "+11", filepath)
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
  createFolder()
	filepath := writeFile(filename)
  openFile(filepath)

	fmt.Println("Meeting successfully saved!")
}
