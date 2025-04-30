package job

import (
	"main/src/modules/constants"
	"main/src/modules/factory"
	"os"
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

func StartMiraeJob(errors []string, needErrors bool) {
	var dataProcesser factory.Dataprocesser
	// Suggested to provide only text data such as pure values or keys
	//reason: some places "key" :"value" might reach to code as "\"key\""\n
	// no regex,used purely made by utilizing Index functions available on Strings package
	dataProcesser = factory.Factory(constants.Mirae)

	// process data
	dataProcesser.GetFileDataConnectAndProcess(file, true, errors...)
	dataProcesser.GetFileDataConnectAndProcess(file, false, errors...)

}


var dataProcesser factory.Dataprocesser
func StartScrapperJob(errors []string, needErrors bool) {
	// Suggested to provide only text data such as pure values or keys
	//reason: some places "key" :"value" might reach to code as "\"key\""\n
	// no regex,used purely made by utilizing Index functions available on Strings package
	dataProcesser = factory.Factory(constants.Scrapper)

	// process data
	dataProcesser.GetFileDataConnectAndProcess(file, true, errors...)
	dataProcesser.GetFileDataConnectAndProcess(file, false, errors...)
	


}
