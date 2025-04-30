package main

import (
	"main/src/modules/job"
)

func main() {
	errors := []string{`Client.Timeout exceeded while awaiting header`,
		"invalid session. Kindly logout and login again",
		"Symbol not found",
		"Data too long for column 'iBasketName' at row",
	}

	job.StartScrapperJob(errors, false)

}
