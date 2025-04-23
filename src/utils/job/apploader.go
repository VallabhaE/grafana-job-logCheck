package job

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"main/src/utils/constants"
	"main/src/utils/utils"
	"os"
	"strings"
	"time"
)

var file *os.File

func Init(fileName string) error {
	var err error
	if _, err := os.Stat(constants.OUPUT_DIR); err != nil {
		os.MkdirAll(constants.OUPUT_DIR, 0744)
	}
	file, err = os.Open(fileName)
	if err != nil {
		return err
	}
	return nil
}

func Start(errors []string, needErrors bool) {

	// Suggested to provide only text data such as pure values or keys
	//reason: some places "key" :"value" might reach to code as "\"key\""\n
	// no regex,used purely made by utilizing Index functions available on Strings package
	__getFileDataMiraeConnectAndProcess(needErrors, errors...) 
}

// func purely expects brokerUserId which means works only for authorized user api calls only,
// change it according to needs further
func __getFileDataMiraeConnectAndProcess(needErrors bool, Error ...string) bool {

	// file related vars
	var cache = make(map[string]string)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var res = ""
	fMatched, err1 := os.Create(constants.OUPUT_DIR + "matchedErrors.txt")
	f, err2 := os.Create(constants.OUT_FILE_PATH)
	if err1 != nil || err2 != nil {
		fmt.Println(errors.Join(err1, err2))
		return false
	}

	defer fMatched.Close()
	// line number vars
	lineNum := 1
	fileLineNum := 1
	skipedLines := 1
	matchedErrCound := 1

	f.Write([]byte("Task Triggered at : " + string(time.Now().Format("2006-01-02 15:04:05")+"\n")))
	for scanner.Scan() {
		line := scanner.Text()
		time := strings.Split(line, ",")[0]
		Erridx := utils.GtErrorIdxCheck(line, Error, needErrors) // Error Existence check
		idx := strings.LastIndex(line, "brokerUserId")
		if idx == -1 || Erridx == -1 {
			lineNum++
			skipedLines++
			// res = fmt.Sprintln("Line ", lineNum-1, "Skipped") + res
			continue
		}
		brokerUserIdStr := line[idx : idx+30]
		idx2 := strings.LastIndex(brokerUserIdStr, `,"`)
		if idx2 == -1 {
			lineNum++
			skipedLines++
			// res = fmt.Sprintln("Line ", lineNum-1, "Skipped") + res
			continue
		}
		fMatched.Write([]byte(line + "\n"))
		if _, ok := cache[brokerUserIdStr]; ok {
			lineNum++
			matchedErrCound++
			continue
		}
		cache[brokerUserIdStr] = "Exist"

		res += fmt.Sprintf("%d: %s %s\n", fileLineNum, time, brokerUserIdStr[:idx2])
		fileLineNum++
		lineNum++
	}
	d, err := json.Marshal(utils.Counter)
	if err != nil {
		return false
	}

	f.Write([]byte(res))

	f.Write([]byte(d))

	f.Write([]byte("\n\n\n\n"))
	f.Write([]byte("\n\n\n\n"))
	f.Write([]byte("====================================\n"))
	f.Write([]byte(fmt.Sprintf("Skipped Line Count : %d\n", skipedLines-1)))
	f.Write([]byte(fmt.Sprintf("Total Errors on backup.csv : %d\n", lineNum-1)))
	f.Write([]byte(fmt.Sprintf("Total Errors Filtered : %d\n", matchedErrCound-1)))

	if needErrors {
		f.Write([]byte("Read matchedErrors.txt for more details about new errors found other then provided\n"))
	}
	f.Write([]byte("====================================\n"))

	return true
}
