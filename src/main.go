package main

import (
	"fmt"
	"io/ioutil"
	"main/src/modules/job"
	"net/http"
	"os"
)

func main() {
	errors := []string{`Client.Timeout exceeded while awaiting header`,
		"invalid session. Kindly logout and login again",
	}

	job.Start(errors, false)

}
