package lasagna

	const OvenTime = 40
	var remain int
	var preptime int
	var elapsed int

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    remain = OvenTime - actualMinutesInOven 
    return remain
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    preptime = numberOfLayers * 2
    return preptime
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    preptime = PreparationTime(numberOfLayers)
    elapsed = preptime + actualMinutesInOven
    return elapsed
}
