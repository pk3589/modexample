package mylog

import (
	"log"
	"os"
)

var(
	Info *log.Logger
	Error *log.Logger
	Debug *log.Logger
)

func init(){

	//init info log
	infoFile,err := os.OpenFile("info.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		log.Fatalln("Open info log file error:",err)
	}

	Info = log.New(infoFile,"[InfoLog] ",log.LstdFlags|log.Lshortfile)

	//init error log
	errorFile,err := os.OpenFile("error.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		log.Fatalln("Open error log file error:",err)
	}

	Error = log.New(errorFile,"[ErrorLog] ",log.LstdFlags|log.Lshortfile)

	//init debug log
	debugFile,err := os.OpenFile("debug.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		log.Fatalln("Open debug log file error:",err)
	}

	Debug = log.New(debugFile,"[InfoLog] ",log.LstdFlags|log.Lshortfile)



}


