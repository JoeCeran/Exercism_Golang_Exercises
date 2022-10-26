package elon

import "fmt"

func (c *Car) Drive() string {
	fmt.Println(c.battery - c.batteryDrain)
	if c.battery - c.batteryDrain > 0 {
		c.battery = c.battery - c.batteryDrain
		c.distance = c.distance + c.speed
	}
	return fmt.Sprintf("car is now Car(speed: %d, batteryDrain: %d, battery: %d, distance: %d)", c.speed, c.batteryDrain,
	c.battery, c.distance)
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d", c.battery) + "%"
}

func (c *Car) CanFinish(tracDistance int) bool {
	var returnFinish bool
	for c.distance < tracDistance && c.battery - c.batteryDrain >= 0{
		c.battery = c.battery - c.batteryDrain
		c.distance = c.distance + c.speed
		if c.distance >= tracDistance {
			 returnFinish = true
		 } else {
			 returnFinish = false
		 }

	 }
	 return returnFinish
}