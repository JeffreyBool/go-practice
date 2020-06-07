/**
* Author: JeffreyBool
* Date: 2020/6/6
* Time: 15:52
* Software: GoLand
 */

package wire

import (
	"testing"
)

func TestWire(t *testing.T) {
	mission := InitMission("dj")
	mission.Start()
}
