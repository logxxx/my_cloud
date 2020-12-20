package temp

import (
	"log"
	"os"
)
import "../locate"

func commitTempObject(datFile string, tempinfo *tempInfo) {
	log.Println("commitTempObject:", datFile)
	e := os.Rename(datFile, os.Getenv("STORAGE_ROOT") + "/objects/" + tempinfo.Name)
	if e != nil {
		log.Println("commitTempObject err:", e)
	}
	locate.Add(tempinfo.Name)
}
