package main

import (
	"fmt"

	"github.com/JeffreyBool/go-practice/golang从入门到精通/listen10/config_reflect"
)

type Config struct {
	Server Server `ini:"server"`
	Mysql  Mysql  `ini:"mysql"`
}

type Server struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}

type Mysql struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Timeout  int    `ini:"timeout"`
}

func main() {
	filename := "C:/tmp/config.ini.txt"
	var conf Config
	err := config_reflect.UnMarshalFile(filename, &conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("conf:%#v\n", conf)
}
