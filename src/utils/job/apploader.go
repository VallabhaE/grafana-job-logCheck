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

func Start(errors []string) {

	// Suggested to provide only text data such as pure values or keys
	//reason: some places "key" :"value" might reach to code as "\"key\""
	// no regex used purely made by utilizing Index functions available on Strings package
	__getFileDataMiraeConnectAndProcess(errors...)
}

// func purely expects brokerUserId which means works only for authorized user api calls only,
// change it according to needs further
func __getFileDataMiraeConnectAndProcess(Error ...string) {

	defer file.Close()
	var cache = make(map[string]string)
	var res = ""
	lineNum := 1
	fileLineNum := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		time := strings.Split(line, ",")[0]
		Erridx := utils.GetErrorIdxCheck(line, Error) // Error Existence check
		idx := strings.LastIndex(line, "brokerUserId")
		if idx == -1 || Erridx == -1 {
			lineNum++
			file.Write([]byte(fmt.Sprintf("Line ", lineNum-1, "Skipped")))
			continue
		}
		brokerUserIdStr := line[idx : idx+30]
		idx2 := strings.LastIndex(brokerUserIdStr, `,"`)
		if idx2 == -1 {
			lineNum++
			file.Write([]byte(fmt.Sprintf("Line ", lineNum-1, "Skipped")))
			continue
		}
		if _, ok := cache[brokerUserIdStr]; ok {
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
}
