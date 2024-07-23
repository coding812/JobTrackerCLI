package main

import (
	. "fmt"
	"log"
	"os"
	"strings"
)

const (
    dateWidth = 12 
    companyNameWidth = 15
    interviewTypeWidth = 15
)

func main() {
	file, err := openFile()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // closes the file when program finishes
	printExisting() // prints all existing entries
	
	if err := getInput(file); err != nil {
		log.Fatalf("Failed to record input: %v", err)
	} 
}

// Opens the JobsApplied.txt file, in read/write mode, or creates it if it does not exist
func openFile() (*os.File, error) {
	var file, err = os.OpenFile("AppliedJobs.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	// Check if the file is newly created
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if fileInfo.Size() == 0 {
		// Write initial information to the file
		initialInfo := "Date Applied ---- Company Name ---- Interview Type\n"
		if _, err := file.WriteString(initialInfo); err != nil {
			return nil, err
		}
	}
	return file, nil
}

// Gets the input to be saved to the AppliedJobs.txt
func getInput(file *os.File) error {
    Println("What date did you apply?")
    var date string
    if _, err := Scanln(&date); err != nil {
        return Errorf("date not valid: %v", err)
    }

    Println("What was the company name?")
    var companyName string
    if _, err := Scanln(&companyName); err != nil {
        return Errorf("company name not valid: %v", err)
    }

    Println("What type of interview was it? Phone? Video call?")
    var interviewType string
    if _, err := Scanln(&interviewType); err != nil {
        return Errorf("interview type not valid: %v", err)
    }

    // Format the entry with padded fields
    formattedEntry := padRight(date, dateWidth) + "    " + padRight(companyName, companyNameWidth) + "    " + padRight(interviewType, interviewTypeWidth) + "\n"

    if _, err := file.WriteString(formattedEntry); err != nil {
        return Errorf("failed to write to file: %v", err)
    }
    return nil
}

// Function to pad a string to a fixed width
func padRight(str string, length int) string {
    if len(str) >= length {
        return str
    }
    return str + strings.Repeat(" ", length-len(str))
}

// Fetches all the current entries in the AppliedJobs.txt file and prints them out
func printExisting() {
	existingBytes, err := os.ReadFile("AppliedJobs.txt")
	if err != nil {
		log.Fatalf("Failed to read existing entries: %v", err)
	}
	Println("Current entries in AppliedJobs.txt:")
	Println(string(existingBytes))
}
