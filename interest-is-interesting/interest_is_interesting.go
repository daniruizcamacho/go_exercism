package interest

const (
	InterestRateNegativeBalance        float32 = 3.213
	InterestRatePositiveFirstInterval  float32 = 0.5
	InterestRatePositiveSecondInterval float32 = 1.621
	InterestRatePositiveThirdInterval  float32 = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return InterestRateNegativeBalance
	}

	if balance >= 0 && balance < 1000 {
		return InterestRatePositiveFirstInterval
	}

	if balance >= 1000 && balance < 5000 {
		return InterestRatePositiveSecondInterval
	}

	return InterestRatePositiveThirdInterval
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interestRate := InterestRate(balance)

	return balance * float64(interestRate) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	count := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		count += 1
	}

	return count
}
