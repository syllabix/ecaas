package ecaas

import (
	"errors"

	"github.com/shopspring/decimal"
)

//CalculateTotalCost is a function used to total and provide an estimate range
//for a move based upon the following business rules:
//	1. Multiply total estimated hours by the provided hourly rate for an initial subtotal
//  2. Apply the cost multiplier - effectively a service fee - to the subtotal
//  3. Apply tax to the subtotal plus service fee to generate the low end of the etimate
//  4. To estimate the high end of the estimate, if the preferred move date is on Friday or Saturday
//		- add a complexity factor of 30% of the low estimate before tax - otherwise add a standard 15% weekday
//		complexity factor to get a estimate high end subtotal
//  5. Apply tax to the high end of the estimate, end return the range formatted in USD
func CalculateTotalCost(details *JobDetails, moveType MoveType) *EstimateRange {
	//TODO - Implement
	return nil
}

func addTax(total decimal.Decimal, taxRate decimal.Decimal) decimal.Decimal {
	return addPercentageToTotal(total, taxRate)
}

func addPercentageToTotal(total decimal.Decimal, percentage decimal.Decimal) decimal.Decimal {
	return total.Mul(percentage.Add(decimal.NewFromFloat(1)))
}

func jobComplexityForDate(dateString string) (float64, error) {
	//TODO: Implement function
	return 0, errors.New("Not Implemented")
}
