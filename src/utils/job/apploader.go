package job

import (
	"bufio"
	"fmt"
	"main/src/utils/constants"
	"main/src/utils/utils"
	"os"
	"strings"
)

var file *os.File

func Init(fileName string) error {
	var err error
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
func __getFileDataMiraeConnectAndProcess(needErrors bool, Error ...string) {

	// file related vars
	defer file.Close()
	var cache = make(map[string]string)
	scanner := bufio.NewScanner(file)
	var res = ""
	fMatched, err := os.Create(constants.OUPUT_DIR + "matchedErrors.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fMatched.Close()

	// line number vars
	lineNum := 1
	fileLineNum := 1
	skipedLines := 1
	matchedErrCound := 1
	for scanner.Scan() {
		line := scanner.Text()
		time := strings.Split(line, ",")[0]
		Erridx := utils.GetErrorIdxCheck(line, Error, needErrors) // Error Existence check
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
		if _, ok := cache[brokerUserIdStr]; ok {
			lineNum++
			matchedErrCound++
			fMatched.Write([]byte(line + "\n"))
			continue
		}
		cache[brokerUserIdStr] = "Exist"

		res += fmt.Sprintf("%d: %s %s\n", fileLineNum, time, brokerUserIdStr[:idx2])
		fileLineNum++
		lineNum++
	}

	f, err := os.Create(constants.OUT_FILE_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte(res))
	f.Write([]byte("\n\n\n\n"))
	f.Write([]byte("====================================\n"))
	f.Write([]byte(fmt.Sprintf("Skipped Line Count : %d\n", skipedLines)))
	f.Write([]byte(fmt.Sprintf("Total Errors on backup.csv : %d\n", lineNum-skipedLines)))
	f.Write([]byte(fmt.Sprintf("Total Errors Filtered : %d\n", matchedErrCound-1)))
	f.Write([]byte("====================================\n"))

}
