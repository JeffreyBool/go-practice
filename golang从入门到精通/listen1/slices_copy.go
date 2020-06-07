/**
* Author: JeffreyBool
* Date: 2020/6/5
* Time: 19:40
* Software: GoLand
 */

package main

import (
	"fmt"
)

// Driver data.
type Driver struct {
	trips []Trip
}

// []Trip data source.
type Trip struct {
	Id int
}

// bad
//func (d *Driver) SetTrips(trips []Trip) {
//	d.trips = trips
//}

// good
func (d *Driver) SetTrips(trips []Trip) {
	d.trips = make([]Trip, len(trips))
	copy(d.trips,trips)
}

func (d Driver) GetTrips() []Trip {
	return d.trips
}

func main() {
	trips := []Trip{0:{1},1:{2}}
	d := new(Driver)
	d.SetTrips(trips)
	trips[0] = Trip{Id: 10}
	d.SetTrips(trips)
	fmt.Printf("%#v\n",trips)
	fmt.Printf("%#v\n",d.GetTrips())
}
