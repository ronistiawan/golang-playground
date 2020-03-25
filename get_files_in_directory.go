// What this program do?
// Get All text file in current directory
// Read the content, and append some string on each line
// Write new files

package main

import (
    "fmt"
	 "log"
	 "os"
	 "bufio"
	 "path/filepath"
	 "strings"
)

func main() {
    files, err := filepath.Glob("./*.TXT")
    if err != nil {log.Fatal(err)}

    for _, filename := range files {

		fmt.Println(filename)
		f, _ := os.Open(filename)
		
		newContent := ""
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			newContent += scanner.Text()
			newContent += "                                             0\r\n"
		}

		f.Close()
		err := os.Remove(filename)

		if err != nil {
			fmt.Println(err)
		}

		newFile , err := os.Create(strings.Replace(filename, ".TXT", ".RSP", 1))
		if err != nil {log.Fatal(err)}
		newFile.Sync()
		w := bufio.NewWriter(newFile)
		w.WriteString(newContent)
		
		w.Flush()
	}
}