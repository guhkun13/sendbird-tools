// source: https://gosamples.dev/read-csv/
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/guhkun13/sendbird-tools/service"
)

// main
func main() {
	// open file
	csvFile := os.Args[1]

	f, err := os.Open(csvFile)
	if err != nil {
		fmt.Printf("error: %v \n", err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// call service
	userList := service.NewService().CreateUserList(data)
	logFile := service.NewService().CreateLogFile(csvFile)
	req := service.WorkerRequest{
		Users:   userList,
		LogFile: logFile,
	}

	service.NewService().OnboardingUser(req)

	// close logFile
	defer logFile.Close()
}
