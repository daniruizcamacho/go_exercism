package elon

import "fmt"

func (car *Car) Drive() {
	if car.batteryDrain > car.battery {
		return
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
	batteryNeedIt := trackDistance / car.speed * car.batteryDrain

	if car.battery >= batteryNeedIt {
		return true
	}
	return false
}
