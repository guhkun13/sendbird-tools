package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const logDirName = "logs"

func (s *ServiceImpl) CreateLogFile(csvFile string) (res *os.File) {
	// ignore the error
	_ = os.Mkdir(logDirName, os.ModePerm)

	ts := time.Now().Format("2006-01-02-15-04")
	strTs := fmt.Sprintf("%v", ts)
	outFile := logDirName + "/input_" + csvFile + "_" + strTs + ".log"

	res, err := os.OpenFile(outFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("output file = %v \n", outFile)
	return
}

func (s *ServiceImpl) WriteLog(data HttpLog, f *os.File) {
	dataByte, _ := json.MarshalIndent(data, "", " ")

	n, err := f.Write(dataByte)
	if err != nil {
		fmt.Println(n, err)
	}
	if n, err = f.WriteString(",\n"); err != nil {
		fmt.Println(n, err)
	}
}
