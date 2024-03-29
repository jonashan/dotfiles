package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func writeFile() string {
  today := time.Now().Format("2006-01-02")
  filepath := fmt.Sprintf("/Users/jonashansen/SecondBrain/05. Journals/Daily/%s.md", today)

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
	defer file.Close()

  tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
  yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	writer := bufio.NewWriter(file)
	id := fmt.Sprintf("id: %s", time.Now().Format("20060102150405"))
	datetime := fmt.Sprintf("created_at: %s", time.Now().Format("2006-01-02 15:04:05"))
	lines := []string{
		"---",
		datetime,
		id,
		"tags:",
		"- daily",
    "- journal",
		"---",
    "# " + today,
    "[[" + yesterday + "]] - [[" + tomorrow + "]]",
		"",
		"## Gratitude",
		"",
    "",
		"## What's todays most important task?",
		"",
    "",
		"## Free journaling / Reflection",
		"",
    "",
    "## Links",
    "[[Daily Journals]]",
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
  cmd := exec.Command("nvim", "+12", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	error := cmd.Run()
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(0)
	}
}

func main() {
	filepath := writeFile()
  openFile(filepath)

	fmt.Println("Meeting successfully saved!")
}
