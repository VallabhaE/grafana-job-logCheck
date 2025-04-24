package miraeconnect

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"main/src/modules/constants"
	"main/src/modules/utils"
	"os"
	"strings"
	"time"
)

type MiraeDefault struct{}

var MiraeData []string

func initAndProcessData(filename string) {
	if len(MiraeData) > 0 {
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
		MiraeData = append(MiraeData, line)
	}
	fmt.Println("Data Loaded for Mirae")
	return
}

// func purely expects brokerUserId which means works only for authorized user api calls only,
// change it according to needs further
func (s *MiraeDefault) GetFileDataMiraeConnectAndProcess(file *os.File, needErrors bool, Error ...string) bool {

	// file related vars
	defer file.Close()
	var cache = make(map[string]string)
	initAndProcessData(constants.GRAFANA_FILE_PATH)
	var res = ""
	matchedErrorsfName := "matchedErrors.txt"
	if !needErrors {
		matchedErrorsfName = "exc" + matchedErrorsfName
	} else {
		matchedErrorsfName = "inc" + matchedErrorsfName
	}
	fMatched, err1 := os.Create(constants.OUPUT_DIR + matchedErrorsfName)
	var outFileName string
	if !needErrors {
		outFileName = constants.OUPUT_DIR + "exc" + constants.OUT_FILE
	} else {
		outFileName = constants.OUPUT_DIR + "inc" + constants.OUT_FILE
	}
	f, err2 := os.Create(outFileName)
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
	for _, line := range MiraeData {
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
		matchedErrCound++

		if _, ok := cache[brokerUserIdStr]; ok {
			lineNum++
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
