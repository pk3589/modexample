package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{

			"message": "pong",
		})

	})
	fmt.Printf("come on")
	r.Run() // listen and serve on 0.0.0.0:8080


	//pintLogToFile()

	//mylog.Info.Println("这是一行info日志")
	//mylog.Error.Println("这是一行error日志")
	//mylog.Debug.Println("这是一行debug日志")

}


func logTest(){

	log.Print("1.Go log Print test")
	log.Println("2.Go log Println test")
	log.Fatalln("Go log Fatalln test")
	log.Print("3.Go log Print test")
	log.Println("4.Go log Println test")
}

func logFormatTest(){
	//自定义个log格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.Print("1.Go log Print test")
	log.Println("2.Go log Println test")
	log.Fatalln("Go log Fatalln test")
	log.Print("3.Go log Print test")
	log.Println("4.Go log Println test")
}

func pintLogToFile(){
	//open file
	file,err := os.OpenFile("log.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		log.Fatalln("open log file error：",err)
	}
	defer file.Close()

	//配置日志 指定输出位置  前缀  和 日志格式
	logger := log.New(file,"[LogFile] ",log.LstdFlags|log.Lshortfile)

	logger.Println("------")
	logger.Print("1.Go log Print test")
	logger.Println("2.Go log Println test")
	logger.Fatalln("Go log Fatalln test")
	logger.Print("3.Go log Print test")
	logger.Println("4.Go log Println test")
}
