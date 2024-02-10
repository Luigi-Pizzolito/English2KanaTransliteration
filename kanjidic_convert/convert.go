package main

import (
	"bufio"
	"os"
	"fmt"
	// "foosoft.net/projects/jmdict"
	"git.foosoft.net/alex/jmdict"
	"encoding/json"
)

func main() {
	// Open the file
    file, err := os.Open("kanjidic2.xml")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Wrap the file descriptor with bufio.NewReader
    reader := bufio.NewReader(file)

	// Load the character dictionary
	fmt.Println("Loading Kanji Dictionary")
	kanjidic, err := jmdict.LoadKanjidic(reader)

	// Iterate and copy to map
	pronounceMap := make(map[string]string)
	for _, kanji := range kanjidic.Characters {
		if kanji.ReadingMeaning != nil {
			reading := ""
			readings := kanji.ReadingMeaning.Readings
			if readings != nil {
				for _, option := range readings {
					if option.Type == "ja_on" {
						reading = option.Value
						fmt.Println(kanji.Literal, "->", reading)
						pronounceMap[kanji.Literal] = reading
						break
					}
				}
			}
		}
	}
	fmt.Println("Extraction done")

	// Convert the data to JSON
	jsonData, err := json.Marshal(pronounceMap)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Write the JSON data to a file
	file, err = os.Create("kanjidic2_pronounce.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data written to kanjidic2_pronounce.json")
}