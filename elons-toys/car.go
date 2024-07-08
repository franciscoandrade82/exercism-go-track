package elon

import "fmt"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

func (c *Car) String() string {
	return fmt.Sprintf("Speed: %d\nBatteryDrain: %d\nBattery: %d\nDistance: %d", c.speed, c.batteryDrain, c.battery, c.distance)
}
