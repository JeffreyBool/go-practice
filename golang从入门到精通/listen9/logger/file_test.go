/**
* Author: JeffreyBool
* Date: 2020/7/9
* Time: 23:51
* Software: GoLand
 */

package logger_test

import (
	"testing"

	"github.com/JeffreyBool/go-practice/golang从入门到精通/listen9/logger"
)

func TestNewFile(t *testing.T) {
	log := logger.NewFile(logger.DebugLevel,"/Users/zhanggaoyuan/go/src/github.com/JeffreyBool/go-practice/golang从入门到精通/listen9/logger/logs","test")
	defer log.Close()
	log.Debug("user id %d is come from china",1000)
	log.Warn("test warn log")
	log.Fatal("test fatal log")
}
