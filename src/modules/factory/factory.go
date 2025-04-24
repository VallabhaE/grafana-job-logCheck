package factory

import (
	"main/src/modules/constants"
	miraeconnect "main/src/modules/factory/miraeConnect"
	"os"
)

type Dataprocesser interface {
	GetFileDataMiraeConnectAndProcess(file *os.File, needErrors bool, Error ...string) bool
}

func Factory(typ string) Dataprocesser {
	switch typ {
	case constants.Mirae:
		return &miraeconnect.MiraeDefault{}
	default:
		return nil
	}
}
