package nfs

import (
	"os"
)

const configFile = "/etc/exports"

func ReadConfig() error {
	dat, err := os.Open(configFile)
	defer dat.Close()
	if err != nil {
		//handle error and such
	}
	//config, err := DeserializeConfig(dat)
}
