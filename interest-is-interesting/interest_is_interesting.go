package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var returnValue float32
	if balance < 0 {
		returnValue = 3.213
	} else if balance >= 0 && balance < 1000 {
		returnValue = 0.5
	} else if balance >= 1000 && balance < 5000 {
		returnValue = 1.621
	} else if balance >= 5000 {
		returnValue =  2.475
	}
	return returnValue
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int = 0
	for balance < targetBalance{
		balance = AnnualBalanceUpdate(balance)
		years = years + 1
	}
	return years
}
