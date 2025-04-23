package main

import (
	"fmt"
	"main/src/utils/constants"
	"main/src/utils/job"
)

func main() {

	needErrors := true

	err := job.Init(constants.GRAFANA_FILE_PATH)
	errors := []string{`context deadline exceeded (Client.Timeout exceeded while awaiting headers)`, 
	"invalid session. Kindly logout and login again",
}

	if err != nil {
		fmt.Errorf("Error opening file:", err)
		return
	}
	job.Start(errors, needErrors)
}
