package bb8

import "fmt"

// Collision contains collision data reported by BB-8
type Collision struct {
	x int16
	y int16
}

// OnCollision handles collision events.
func OnCollision(bb8 *bb8, data Collision) {
	fmt.Printf("Collision detected: %+v\n", data)
	bb8.Flash()
	bb8.TurnAround()
}
