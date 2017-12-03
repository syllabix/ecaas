package ecaas

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

//CalculateTotalCost return the
func CalculateTotalCost(details *JobDetails, moveType MoveType) *EstimateRange {

	totalHours := decimal.NewFromFloat(details.Hours)
	hourlyRate, _ := decimal.NewFromString(details.HourlyRate)

	taxRate, _ := decimal.NewFromString(moveType.GetTaxRate())
	jobMultiplier, _ := decimal.NewFromString(moveType.GetMultiplier())

	complexity, _ := jobComplexityForDate(details.PreferredMoveDate)
	complexityFactor := decimal.NewFromFloat(complexity)

	lowTotal := addPercentageToTotal(totalHours.Mul(hourlyRate), jobMultiplier)
	lowTotalWithTax := addTax(lowTotal, taxRate)
	highTotal := addPercentageToTotal(lowTotal, complexityFactor)
	highTotalWithTax := addTax(highTotal, taxRate)

	return &EstimateRange{
		Low:  fmt.Sprintf("$%s", lowTotalWithTax.StringFixedBank(2)),
		High: fmt.Sprintf("$%s", highTotalWithTax.StringFixedBank(2)),
	}
}

func addPercentageToTotal(total decimal.Decimal, percentage decimal.Decimal) decimal.Decimal {
	return total.Mul(percentage.Add(decimal.NewFromFloat(1)))
}

func addTax(total decimal.Decimal, taxRate decimal.Decimal) decimal.Decimal {
	return addPercentageToTotal(total, taxRate)
}

func jobComplexityForDate(dateString string) (float64, error) {
	date, err := time.Parse(time.UnixDate, dateString)
	if err != nil {
		return 0, ErrorInvalidDateFormat
	}
	complexity := 0.15
	if date.Weekday() == time.Friday ||
		date.Weekday() == time.Saturday {
		complexity = 0.3
	}
	return complexity, nil
}
