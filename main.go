package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	bb := bb8.NewDriver(bleAdaptor)

	work := func() {
		heading := 0

		bb.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
			bb.SetRGB(255, 0, 0)
			time.AfterFunc(time.Second, func() {
				bb.SetRGB(0, 255, 0)
			})

			// Pick a new heading.
			heading = (heading + 180 + rand.Intn(21) - 10) % 360
			bb.Roll(100, uint16(heading))
		})

		bb.SetRGB(0, 255, 0)
		bb.Roll(100, 0)
	}

	robot := gobot.NewRobot("bb8",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{bb},
		work,
	)

	robot.Start()

}

