package ecaas

import "testing"

//mockMoveType implements the MoveType Interface
type mockMoveType struct {
	name       string
	multiplier string
	taxRate    string
}

func (m *mockMoveType) GetName() string {
	return m.name
}

func (m *mockMoveType) GetMultiplier() string {
	return m.multiplier
}

func (m *mockMoveType) GetTaxRate() string {
	return m.taxRate
}

var (
	//Test Dates
	monday   = "Mon Dec 04 12:00:00 EST 2017"
	saturday = "Sat Dec 02 12:00:00 EST 2017"
)

func TestJobComplexityForDate(t *testing.T) {

	actual, err := jobComplexityForDate(monday)
	if err != nil {
		t.Error(err)
	}
	expected := 0.15
	if actual != expected {
		t.Errorf("Job complexity factor should be %f, but got %f", expected, actual)
	}

	actual, err = jobComplexityForDate(saturday)
	if err != nil {
		t.Error(err)
	}
	expected = 0.3
	if actual != expected {
		t.Errorf("Job complexity factor should be %f, but got %f", expected, actual)
	}
}

func TestCalculateTotalCost(t *testing.T) {
	//Test a normal weekday move
	details := NewJobDetails(10, "120.50", monday)
	moveType := &mockMoveType{
		name:       "Local",
		multiplier: "0.05",
		taxRate:    "0.06",
	}

	estimateRange := CalculateTotalCost(details, moveType)
	expectedLow := "$1341.17"
	expectedHigh := "$1542.34"
	if estimateRange.Low != expectedLow {
		t.Errorf("Expected estimate low to be %s, but instead got %s", expectedLow, estimateRange.Low)
	}
	if estimateRange.High != expectedHigh {
		t.Errorf("Expected estimate high to be %s, but instead got %s", expectedHigh, estimateRange.High)
	}

	//Testing weekend move
	details = NewJobDetails(10, "120.50", saturday)
	weekendRange := CalculateTotalCost(details, moveType)
	expectedLow = "$1341.17"
	expectedHigh = "$1743.51"
	if weekendRange.Low != expectedLow {
		t.Errorf("Expected estimate low to be %s, but instead got %s", expectedLow, weekendRange.Low)
	}
	if weekendRange.High != expectedHigh {
		t.Errorf("Expected estimate high to be %s, but instead got %s", expectedHigh, weekendRange.High)
	}

}
