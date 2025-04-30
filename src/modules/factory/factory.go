package factory

import (
	"main/src/modules/constants"
	miraeconnect "main/src/modules/factory/miraeConnect"
	"main/src/modules/factory/scrapper"
	"os"
)

type Dataprocesser interface {
	GetFileDataConnectAndProcess(file *os.File, needErrors bool, Error ...string) bool
}

func Factory(typ string) Dataprocesser {
	switch typ {
	case constants.Mirae:
		return &miraeconnect.MiraeDefault{}
	case constants.Scrapper:
		return &scrapper.ScrapperDefault{} /// Need to Fix lot
	default:
		return nil
	}
}
