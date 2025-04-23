package main

import (
	"fmt"
	"main/src/utils/constants"
	"main/src/utils/job"
)

func main() {
	err := job.Init(constants.GRAFANA_FILE_PATH)
	errors := []string{`Client.Timeout`, 
	"invalid session. Kindly logout and login again",
}

	if err != nil {
		fmt.Errorf("Error opening file:", err)
		return
	}
	// job.Start(errors, false)

	job.Start(errors, true)

}
