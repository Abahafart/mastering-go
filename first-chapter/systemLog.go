package main

import (
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sysLog)
		log.Println("Everything is fine")
	} //nothing on screen printed, with journalctl -xe command you can see the output log
}
