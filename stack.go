package englishtocron

import "fmt"

// StartEnd holds optional start and end integer values.
type StartEnd struct {
	Start *int
	End   *int
}

// StartEndString holds optional start and end string values.
type StartEndString struct {
	Start *string
	End   *string
}

// Stack holds the parsing state for one token context.
type Stack struct {
	Owner          Kind
	Frequency      *int
	FrequencyEnd   *int
	FrequencyStart *int
	Min            *StartEnd
	Hour           *StartEnd
	Day            *StartEndString
	Month          *StartEndString
	Year           *StartEnd
	DayOfWeek      *string
	IsAndConnector bool
	IsBetweenRange bool
}

func newStack(owner Kind) *Stack {
	return &Stack{Owner: owner}
}

func intPtr(v int) *int       { return &v }
func strPtr(v string) *string { return &v }

func (s *Stack) frequencyToString() string {
	if s.Frequency == nil {
		return "*"
	}
	return fmt.Sprintf("%d", *s.Frequency)
}
