package bb8

import (
	"math/rand"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero"
	gobotBB8 "gobot.io/x/gobot/platforms/sphero/bb8"
)

// BB8 is an interface to do things with a BB-8 robot.
type BB8 interface {
	Start(callback func()) error
	Stop() error
	AddCollisionEventHandler(handler func(bb8 *bb8, data Collision))
	SetSpeed(speed uint8)
	TurnAround()
	Flash()
}

type bb8 struct {
	bleAdapter *ble.ClientAdaptor
	driver     *gobotBB8.BB8Driver
	robot      *gobot.Robot
	speed      uint8
	heading    uint16
}

// NewBB8 initializes a bluetooth connection to a BB-8 robot.
func NewBB8(bluetoothID string) BB8 {
	bleAdapter := ble.NewClientAdaptor(bluetoothID)
	driver := gobotBB8.NewDriver(bleAdapter)

	return bb8{bleAdapter, driver, nil, 0, 0}
}

// Start turns on the robot.
func (b bb8) Start(callback func()) error {
	b.robot = gobot.NewRobot("bb8",
		[]gobot.Connection{b.bleAdapter},
		[]gobot.Device{b.driver},
		callback,
	)

	return b.robot.Start()
}

// Stop shuts down the robot and ends the connection.
func (b bb8) Stop() error {
	return b.robot.Stop()
}

// AddCollisionEventHandler registers a callback function for collision events.
func (b bb8) AddCollisionEventHandler(handler func(bb8 *bb8, data Collision)) {
	b.driver.On("collision", func(s interface{}) {
		collisionPacket, ok := s.(sphero.CollisionPacket)
		if ok {
			handler(&b, Collision{collisionPacket.X, collisionPacket.Y})
		}
	})
}

// SetSpeed makes BB-8 start rolling.
func (b bb8) SetSpeed(speed uint8) {
	b.speed = speed
	b.driver.Roll(b.speed, b.heading)
}

// TurnAround makes BB-8 do a roughly 180 degree turn and keep on rolling.
func (b bb8) TurnAround() {
	// Pick a new heading that's roughly in the opposite direction.
	b.heading = uint16((int(b.heading) + 180 + rand.Intn(21) - 10) % 360)
	b.driver.Roll(b.speed, b.heading)
}

// Flash makes the LED in BB-8 turn red for one second.
func (b bb8) Flash() {
	b.driver.SetRGB(255, 0, 0)
	time.AfterFunc(time.Second, func() {
		b.driver.SetRGB(0, 255, 0)
	})
}
