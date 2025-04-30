package scrapper

import (
	"bufio"
	"fmt"
	"main/src/modules/constants"
	"os"
	"strings"
	"time"
)

type ScrapperDefault struct{}

var scrapperData []string

func initAndProcessData(filename string) {
	if len(scrapperData) > 0 {
		return
	}
	var err error
	if _, err := os.Stat(constants.OUPUT_DIR); err != nil {
		os.MkdirAll(constants.OUPUT_DIR, 0744)
	}
	file, err := os.Open(filename)
	scanner := bufio.NewScanner(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	for scanner.Scan() {
		line := scanner.Text()
		scrapperData = append(scrapperData, line)
	}
	fmt.Println("Data Loaded for Scrapper")
	return
}

// func purely expects brokerUserId which means works only for authorized user api calls only,
// change it according to needs further
func (s *ScrapperDefault) GetFileDataConnectAndProcess(file *os.File, needErrors bool, Error ...string) bool {
	// file related vars
	defer file.Close()
	initAndProcessData(constants.GRAFANA_FILE_PATH)
	errorFile, err := os.Create(constants.OUPUT_DIR + "AllErrorsFile.txt")
	if err != nil {
		fmt.Println(errorFile)
		return false
	}
	errorFile.Write([]byte("Task Triggered at : " + string(time.Now().Format("2006-01-02 15:04:05")+"\n")))
	for _, line := range scrapperData {
		//regardless of what happens we want error details and brokeruserId

		time := strings.Split(line, ",")[0]
		errorKeyidx := strings.LastIndex(line, `""ex"":`)
		if errorKeyidx != -1 {
			if len(line[errorKeyidx:]) > 100 {
				errorFile.Write([]byte(fmt.Sprintf("%s %s", time, line[errorKeyidx+10:errorKeyidx+90])))
				errorFile.Write([]byte(fmt.Sprintf("\n")))

				continue
			}
			errorFile.Write([]byte(fmt.Sprintf("%s %s", time, line[errorKeyidx:])))
			errorFile.Write([]byte(fmt.Sprintf("\n")))

		}
	}
	return true
}
