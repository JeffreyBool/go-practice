/**
* Author: JeffreyBool
* Date: 2020/7/9
* Time: 23:51
* Software: GoLand
 */

package logger_test

import (
	"testing"
	"time"

	"github.com/JeffreyBool/go-practice/golang从入门到精通/listen9/logger"
)

func TestNewLogger(t *testing.T) {
	log := logger.New(logger.SetLevel(logger.DebugLevel), logger.SetPath("/Users/zhanggaoyuan/go/src/github.com/JeffreyBool/go-practice/golang从入门到精通/listen9/logger/logs/all.log"))
	defer log.Close()
	for {
		log.Debug("user id %d is come from china", 1000)
		log.Warn("test warn log")
		log.Fatal("test fatal log")
		time.Sleep(time.Second)
	}
}


