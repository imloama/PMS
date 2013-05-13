package main

import (
	. "PMS/controllers"
	. "PMS/util"
	"fmt"
	"net/http"
	"os"
)

var (
	AppConfig *Config
)

func init() { //初始化基本功能

	//fmt.Println(AppConfig.String("appname"))
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {

	AppPath, _ := os.Getwd()
	fmt.Println(AppPath)
	/*
		var err error
		AppConfig, err = LoadConfig(AppPath + "/conf/app.conf")

		if err != nil {
			fmt.Println("server init error!" + err.Error())
			return
		}*/

	//http.HandleFunc("/", handler)

	http.HandleFunc("/", Index)
	http.ListenAndServe(":8088", nil)

	//data, count, err := ReadHtml("views/index.html")
	//fmt.Printf("file content:%q\n", data[:count])
}
