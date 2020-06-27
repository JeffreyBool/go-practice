/**
* Author: JeffreyBool
* Date: 2020/6/9
* Time: 18:01
* Software: GoLand
 */

package main

import (
	"fmt"
	"time"

	"github.com/JeffreyBool/go-practice/pkg/rabbitmq"
)

func main() {
	fmt.Println("hello mq_handler")
	var handler = rabbitmq.MQService{}
	go sendSomething(handler)
	handler.Read(tt)
}

func tt(jsonStr []byte) {
	fmt.Println("already in listenning mq: ", string(jsonStr))
}

func sendSomething(handler rabbitmq.MQ) {
	time.Sleep(time.Second * 1)

	handler.Delay("testKey1", "testValue1", "11114294967296")

	//time.Sleep(time.Second * 1)
	//handler.Delay("testKey2", "testValue2", "4000")
	//
	//time.Sleep(time.Second * 1)
	//handler.Delay("testKey3", "testValue3", "5000")
	//
	//handler.Delay("delayKey1", "this is delay key233", "6000")
}
